package microsoftgp

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

func GetAllMSPGGames() ([]GamepassGameDetails, error) {
	gamepassResponseAllGamesTruncated, err := GetMSGPGameIDs()
	if err != nil {
		fmt.Println(err)
		return []GamepassGameDetails{}, err
	}
	games, err := GetMSGPGamesDetails(gamepassResponseAllGamesTruncated)
	if err != nil {
		fmt.Println(err)
		return []GamepassGameDetails{}, err
	}
	return games, nil
}

func GetMSGPGameIDs() ([]string, error) {
	client := resty.New()
	requestURL := "https://catalog.gamepass.com/sigls/v2?id=29a81209-df6f-41fd-a528-2ae6b91f719c&language=en-us&market=US"
	var gamepassResponseAllGames GamepassResponseAllGames
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

func GetMSGPGamesDetails(gameIDs []string) ([]GamepassGameDetails, error) {
	client := resty.New()
	gameIDsJoined := strings.Join(gameIDs, ",")
	requestURL := "https://displaycatalog.mp.microsoft.com/v7.0/products?bigIds=" + gameIDsJoined + "&market=US&languages=en-us&MS-CV=DGU1mcuYo0WMMp"
	var gamepassResponseGamesDetails GamepassResponseGamesDetails
	_, err := client.R().
		SetResult(&gamepassResponseGamesDetails).
		ForceContentType("application/json").
		Get(requestURL)
	if err != nil {
		return []GamepassGameDetails{}, err
	}
	return gamepassResponseGamesDetails.Products, nil
}
