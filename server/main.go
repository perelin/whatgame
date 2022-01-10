package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"whatgameserver/internal/microsoftgp"

	"github.com/Henry-Sarabia/igdb/v2"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/go-resty/resty/v2"
)

var idgbClientID string
var idgbAccessToken string

// var redisURI string
var rdb *redis.Client
var ctx = context.Background()

type Game struct {
	Name    string  `json:"name"`
	Rating  float64 `json:"rating"`
	Image   string  `json:"image"`
	IGDBURL string  `json:"igdburl"`
	IGDBID  int     `json:"igdbid"`
	GPID    string  `json:"gpid"`
}

func CORSMiddleware() gin.HandlerFunc {
	//  https://stackoverflow.com/questions/29418478/go-gin-framework-cors
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func init() {
	log.Println(os.Getenv("GIN_MODE"))
	idgbClientID = os.Getenv("IGDB_CLIENT_ID")
	idgbAccessToken = os.Getenv("IGDB_ACCESS_TOKEN")

	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}
	rdb = redis.NewClient(opt)

	// redisURI = os.Getenv("REDIS_URL")
	// rdb = redis.NewClient(&redis.Options{
	// 	Addr: redisURI,
	// })
}

func main() {

	log.Println("Port!", getPort())

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "po2ng",
		})
	})
	r.GET("/testredis", func(c *gin.Context) {

		ExampleClient()

		c.JSON(200, gin.H{
			"message": "po2ng",
		})
	})
	r.GET("/test", func(c *gin.Context) {

		// gamepassResponseAllGamesTruncated, err := getMSGPGameIDs()
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		// for i, game := range gamepassResponseAllGamesTruncated {
		// 	fmt.Println(i, game)
		// 	gameDetails, err := getMSGPGameDetails(game)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		break
		// 	}
		// 	fmt.Println(gameDetails.LocalizedProperties[0].ProductTitle)

		// 	if i > 10 {
		// 		break
		// 	}
		// }

		igdbClient := igdb.NewClient(idgbClientID, idgbAccessToken, nil)
		// fmt.Println(igdbClient)
		games, err := igdbClient.Games.Search("Viva Piñata",
			//igdb.SetOrder("popularity", igdb.OrderDescending),
			igdb.SetFields("name",
				"rating",
				"rating_count",
				"aggregated_rating",
				"total_rating",
				"total_rating_count",
				"url",
				"category",
				"status",
				"artworks",
				"release_dates",
				"storyline",
				"summary",
				"themes",
				"version_title",
				"videos",
				"websites"))
		if err != nil {
			fmt.Println(err)
			return
		}

		c.JSON(200, gin.H{
			"games": games,
		})
	})
	r.GET("/games/all", func(c *gin.Context) {
		gamesJSON, err := rdb.Get(ctx, "games").Result()
		if err != nil {
			log.Println("GET(/games/all)", err)
			c.JSON(500, gin.H{
				"msg": "couldnt find any games, maybe cache needs to be warmed up",
				"err": err,
			})
			return
		}
		c.String(200, "%s", gamesJSON)
	})

	r.GET("/igdb/game", func(c *gin.Context) {
		gamesJSON, err := rdb.Get(ctx, "games").Result()
		if err != nil {
			log.Println("GET(/games/all)", err)
			c.JSON(500, gin.H{
				"msg": "couldnt find any games, maybe cache needs to be warmed up",
				"err": err,
			})
			return
		}
		c.String(200, "%s", gamesJSON)
	})

	r.PUT("/games/cache", func(c *gin.Context) {
		err := cacheAllGames()
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{
				"msg": err,
			})
			return
		}

		c.JSON(200, gin.H{
			"msg": "ok",
		})
	})
	r.Run()
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
	msgpGames, err := getAllMSPGGames()
	if err != nil {
		return err
	}

	for i, msgpGame := range msgpGames {
		var game Game = Game{}
		var gameMissing bool = true
		game.Name = msgpGame.LocalizedProperties[0].ProductTitle
		game.GPID = msgpGame.ProductID
		igdbGameSearchResults, err := getIGDBGameSearchResults(msgpGame.LocalizedProperties[0].ProductTitle)
		if err != nil {
			log.Println(err)

		}
		for _, igdbGame := range igdbGameSearchResults {
			if msgpGame.LocalizedProperties[0].ProductTitle == igdbGame.Name {
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
					igdbGame, err := getIGDBGameByID(mappedMissingGame.IGDBID)
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

func getIGDBGameSearchResults(gameTitel string) ([]*igdb.Game, error) {
	igdbClient := igdb.NewClient(idgbClientID, idgbAccessToken, nil)
	games, err := igdbClient.Games.Search(gameTitel,
		igdb.SetFields("name",
			"rating",
			"rating_count",
			"aggregated_rating",
			"total_rating",
			"total_rating_count",
			"url",
			"category",
			"status",
			"artworks",
			"release_dates",
			"storyline",
			"summary",
			"themes",
			"version_title",
			"videos",
			"websites"))
	if err != nil {
		return []*igdb.Game{}, err
	}
	return games, nil
}

func getIGDBGameByID(gameID int) (*igdb.Game, error) {
	igdbClient := igdb.NewClient(idgbClientID, idgbAccessToken, nil)
	game, err := igdbClient.Games.Get(gameID,
		igdb.SetFields("name",
			"rating",
			"rating_count",
			"aggregated_rating",
			"total_rating",
			"total_rating_count",
			"url",
			"category",
			"status",
			"artworks",
			"release_dates",
			"storyline",
			"summary",
			"themes",
			"version_title",
			"videos",
			"websites"))
	if err != nil {
		return &igdb.Game{}, err
	}
	return game, nil
}

func getAllMSPGGames() ([]microsoftgp.GamepassGameDetails, error) {
	gamepassResponseAllGamesTruncated, err := getMSGPGameIDs()
	if err != nil {
		fmt.Println(err)
		return []microsoftgp.GamepassGameDetails{}, err
	}
	games, err := getMSGPGamesDetails(gamepassResponseAllGamesTruncated)
	if err != nil {
		fmt.Println(err)
		return []microsoftgp.GamepassGameDetails{}, err
	}
	return games, nil
}

func getMSGPGameIDs() ([]string, error) {
	client := resty.New()
	requestURL := "https://catalog.gamepass.com/sigls/v2?id=29a81209-df6f-41fd-a528-2ae6b91f719c&language=en-us&market=US"
	var gamepassResponseAllGames microsoftgp.GamepassResponseAllGames
	_, err := client.R().
		SetResult(&gamepassResponseAllGames).
		ForceContentType("application/json").
		Get(requestURL)
	if err != nil {
		return []string{}, err
	}
	var msgpIDs []string
	gamepassResponseAllGamesTruncated := gamepassResponseAllGames[1:]
	for _, game := range gamepassResponseAllGamesTruncated {
		msgpIDs = append(msgpIDs, game.ID)
		// fmt.Println(i, game.ID)
	}
	return msgpIDs, nil
}

func getMSGPGamesDetails(gameIDs []string) ([]microsoftgp.GamepassGameDetails, error) {
	client := resty.New()
	gameIDsJoined := strings.Join(gameIDs, ",")
	requestURL := "https://displaycatalog.mp.microsoft.com/v7.0/products?bigIds=" + gameIDsJoined + "&market=US&languages=en-us&MS-CV=DGU1mcuYo0WMMp"
	var gamepassResponseGamesDetails microsoftgp.GamepassResponseGamesDetails
	_, err := client.R().
		SetResult(&gamepassResponseGamesDetails).
		ForceContentType("application/json").
		Get(requestURL)
	if err != nil {
		return []microsoftgp.GamepassGameDetails{}, err
	}
	return gamepassResponseGamesDetails.Products, nil
}

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must bse set2")
	}
	return port
}
