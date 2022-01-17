package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"whatgameserver/internal/helper"
	"whatgameserver/internal/igdbapi"
	"whatgameserver/internal/microsoftgp"

	"github.com/Henry-Sarabia/igdb/v2"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func init() {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}
	rdb = redis.NewClient(opt)
}

func ExampleClient() {

	err := rdb.Set(ctx, "key", "doof", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func cacheAllGames() error {

	log.Println("starting to retrieve and cache all games from MSGP")

	var igdbMatches = 0
	var games = []Game{}
	var missingGames = []Game{}
	//var missingGamesIGDBResults [][]*igdb.Game

	// Load MSGP games
	// msgpGames, err := microsoftgp.GetAllMSPGGames()
	msgpGames, err := microsoftgp.GetAllMSPGGames()
	if err != nil {
		return err
	}
	log.Printf("found %d MSGP games", len(msgpGames))

	// Load mapped missing games
	var mappedMissingGames []Game
	mappedMissingGamesJSON, err := os.ReadFile("./missingGames.json")
	if err != nil {
		log.Println("no missingGames mapping file found", err)
	} else {
		if err := json.Unmarshal(mappedMissingGamesJSON, &mappedMissingGames); err != nil {
			log.Println(err)
		}
		log.Println("loaded manually maped games from file:", len(mappedMissingGames))
	}

	// Iterate over all loaded MSGP games
	for i, msgpGame := range msgpGames {

		log.Println("processing MSGP game:", i, msgpGame.ProductID, msgpGame.LocalizedProperties[0].ProductTitle)

		var game Game = Game{}
		var gameMissing bool = true

		game = mapMSGPData(msgpGame, game)

		igdbGameSearchResults, err := igdbapi.GetIGDBGameSearchResults(msgpGame.LocalizedProperties[0].ProductTitle)
		if err != nil {
			log.Println("couldnt retrieve IGDB results:", err)
		}

		// Find matching IGDB game in manually mapped games
		if gameMissing {
			for _, mappedMissingGame := range mappedMissingGames {
				if msgpGame.ProductID == mappedMissingGame.GPID {
					igdbGame, err := igdbapi.GetIGDBGameByID(mappedMissingGame.IGDBID)
					if err != nil {
						log.Println(err)
						break
					}
					log.Println("found matching game in manually mapped:", igdbGame.Name)
					game = mapIGDBData(igdbGame, game)
					igdbMatches = igdbMatches + 1
					gameMissing = false
					break
				}
			}
		}

		// Find matching IGDB game through titel comparison
		if gameMissing {
			for _, igdbGame := range igdbGameSearchResults {
				if helper.NormalizeTitelString(msgpGame.LocalizedProperties[0].ProductTitle) ==
					helper.NormalizeTitelString(igdbGame.Name) {
					log.Println("found matching name", igdbGame.Name, msgpGame.LocalizedProperties[0].ProductTitle)

					// double check if developer / publisher also fits
					if match, name := checkIfCompaniesMatch(msgpGame, igdbGame); match {
						log.Println("found matching company:", name)
						game = mapIGDBData(igdbGame, game)
						igdbMatches = igdbMatches + 1
						gameMissing = false
						break
					}
				}
			}
		}

		if gameMissing {
			missingGames = append(missingGames, game)
			//missingGamesIGDBResults = append(missingGamesIGDBResults, igdbGameSearchResults)
		}

		games = append(games, game)

		if os.Getenv("GIN_MODE") == "debug" && i > 10 {
			break
		}
	}
	log.Println("total igdbMatches:", igdbMatches, "from", len(msgpGames))
	gamesJSON, _ := json.Marshal(games)
	err = rdb.Set(ctx, "games", gamesJSON, 0).Err()
	if err != nil {
		return err
	}
	log.Println("total missing:", len(missingGames))
	missingGamesJSON, _ := json.Marshal(missingGames)
	err = rdb.Set(ctx, "missingGames", missingGamesJSON, 0).Err()
	if err != nil {
		return err
	}
	// missingGamesIGDBResultsJSON, _ := json.Marshal(missingGamesIGDBResults)
	// err = rdb.Set(ctx, "missingGamesIGDBResults", missingGamesIGDBResultsJSON, 0).Err()
	// if err != nil {
	// 	return err
	// }
	return nil
}

func checkIfCompaniesMatch(msgpGame microsoftgp.GamepassGameDetails, igdbGame *igdb.Game) (bool, string) {

	for _, involvedCompanyID := range igdbGame.InvolvedCompanies {

		idgbInvolvedCompany, err := igdbapi.GetIGDBInvolvedCompanyByID(involvedCompanyID)
		if err != nil {
			log.Println("couldnt retrieve IGDB involved company:", involvedCompanyID, err)
		}

		idgbCompany, err := igdbapi.GetIGDBCompany(idgbInvolvedCompany.Company)
		if err != nil {
			log.Println("couldnt retrieve IGDB company:", idgbCompany, err)
		}

		if helper.CompareNormalizedStrings(msgpGame.LocalizedProperties[0].DeveloperName, idgbCompany.Name) ||
			helper.CompareNormalizedStrings(msgpGame.LocalizedProperties[0].PublisherName, idgbCompany.Name) {
			return true, idgbCompany.Name
		}
	}
	return false, ""
}

func getMSGPStartDate(msgpGame microsoftgp.GamepassGameDetails) {
	//log.Println("MSGP Game:", i, msgpGame.LocalizedProperties[0].ProductTitle)
	//log.Println("MSGP Game:", i, msgpGame.MarketProperties[0].OriginalReleaseDate)
	//game.GPStartDate = msgpGame.DisplaySkuAvailabilities[0].Sku.Properties.Packages[0].PlatformDependencies[0].PlatformName

	for _, msgpgamePackages := range msgpGame.DisplaySkuAvailabilities[0].Sku.Properties.Packages {

		//log.Println("PlatformName", msgpgamePackages.PlatformDependencies[0].PlatformName)
		if msgpgamePackages.PlatformDependencies[0].PlatformName == "Windows.Xbox" {
			spew.Dump(msgpgamePackages)
		}
	}
}

func mapMSGPData(msgpGame microsoftgp.GamepassGameDetails, game Game) Game {
	game.Name = msgpGame.LocalizedProperties[0].ProductTitle
	game.GPID = msgpGame.ProductID
	game.GPReleaseDate = msgpGame.MarketProperties[0].OriginalReleaseDate
	game.GPReleaseDateTimestamp = msgpGame.MarketProperties[0].OriginalReleaseDate.Unix()
	game.GPIDAverageRatingAllTime = microsoftgp.GetAverageRatingAllTime(msgpGame)
	return game
}

func mapIGDBData(igdbGame *igdb.Game, game Game) Game {
	game.IGDBID = igdbGame.ID
	game.IGDBURL = igdbGame.URL
	game.Name = igdbGame.Name
	game.Rating = igdbGame.TotalRating
	game.IGDBFirstReleaseDate = time.Unix(int64(igdbGame.FirstReleaseDate), 0)
	game.IGDBFirstReleaseDateTimestamp = igdbGame.FirstReleaseDate
	return game
}
