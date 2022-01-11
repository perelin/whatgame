package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

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

	var igdbMatches = 0
	var games = []Game{}
	var missingGames = []Game{}
	var missingGamesIGDBResults [][]*igdb.Game

	// Load mapped missing games
	var mappedMissingGames []Game
	mappedMissingGamesJSON, err := os.ReadFile("missingGames.json")
	if err != nil {
		log.Println(err)
	} else {
		if err := json.Unmarshal(mappedMissingGamesJSON, &mappedMissingGames); err != nil {
			log.Println(err)
		}
	}
	spew.Dump(mappedMissingGames)

	// Load MSGP games
	msgpGames, err := microsoftgp.GetAllMSPGGames()
	if err != nil {
		return err
	}

	for i, msgpGame := range msgpGames {
		var game Game = Game{}
		var gameMissing bool = true
		game.Name = msgpGame.LocalizedProperties[0].ProductTitle
		game.GPID = msgpGame.ProductID
		igdbGameSearchResults, err := igdbapi.GetIGDBGameSearchResults(msgpGame.LocalizedProperties[0].ProductTitle)
		if err != nil {
			log.Println(err)
		}
		for _, igdbGame := range igdbGameSearchResults {
			msgpTitle := helper.StandardizeString(msgpGame.LocalizedProperties[0].ProductTitle)
			igdbTitle := helper.StandardizeString(igdbGame.Name)
			log.Println(msgpTitle, igdbTitle)
			if msgpTitle == igdbTitle {
				igdbMatches = igdbMatches + 1
				game.IGDBURL = igdbGame.URL
				game.IGDBID = igdbGame.ID
				game.Rating = igdbGame.Rating
				gameMissing = false
				break
			}
		}
		if gameMissing {
			for _, mappedMissingGame := range mappedMissingGames {
				if msgpGame.ProductID == mappedMissingGame.GPID {
					igdbGame, err := igdbapi.GetIGDBGameByID(mappedMissingGame.IGDBID)
					if err != nil {
						log.Println(err)
						break
					}
					//spew.Dump(igdbGameSearchResults)
					igdbMatches = igdbMatches + 1
					game.IGDBURL = igdbGame.URL
					game.IGDBID = igdbGame.ID
					game.Rating = igdbGame.Rating
					gameMissing = false
				}
			}
		}
		if gameMissing {
			missingGames = append(missingGames, game)
			missingGamesIGDBResults = append(missingGamesIGDBResults, igdbGameSearchResults)
		}
		log.Println(gameMissing, game)
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
	missingGamesIGDBResultsJSON, _ := json.Marshal(missingGamesIGDBResults)
	err = rdb.Set(ctx, "missingGamesIGDBResults", missingGamesIGDBResultsJSON, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
