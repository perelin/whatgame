package main

// todo:
// - title comparison: lowercase and remove special chars

import (
	"fmt"
	"log"
	"os"

	"whatgameserver/internal/igdbapi"

	"github.com/Henry-Sarabia/igdb/v2"
	"github.com/gin-gonic/gin"
)

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

		igdbClient := igdb.NewClient(igdbapi.ClientID, igdbapi.AccessToken, nil)
		games, err := igdbClient.Games.Search("Viva Piñata",
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

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must bse set2")
	}
	return port
}
