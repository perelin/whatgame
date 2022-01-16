package microsoftgp

import (
	"fmt"
	"strings"
	"whatgameserver/internal/helper"

	"github.com/go-resty/resty/v2"
)

func GetAllMSPGGames() ([]GamepassGameDetails, error) {
	gamepassResponseAllGamesTruncated, err := GetAllMSGPGameIDs()
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

func GetAllMSGPGameIDs() ([]string, error) {

	var msgpAllGames []string

	params := []string{
		"id=29a81209-df6f-41fd-a528-2ae6b91f719c&language=en-us&market=US", // All MS
		"id=b8900d09-a491-44cc-916e-32b5acae621b&language=en-us&market=US", // EA1
		"id=1d33fbb9-b895-4732-a8ca-a55c8b99fa2c&language=en-us&market=US", // EA2
		"id=f6505a9f-ec7d-4eb8-a496-be83f8f35829&language=en-us&market=US", // Bethesda1
		"id=79fe89cf-f6a3-48d4-af6c-de4482cf4a51&language=en-us&market=US", // Bethesda2
	}

	for _, param := range params {
		gameIDs, err := GetMSGPGameIDsForParams(param)
		if err != nil {
			return msgpAllGames, err
		}
		msgpAllGames = append(msgpAllGames, gameIDs...)
	}

	msgpAllGamesRediced := helper.RemoveDuplicateStr(msgpAllGames)
	return msgpAllGamesRediced, nil

}

func GetMSGPGameIDsForParams(params string) ([]string, error) {
	client := resty.New()
	requestURL := "https://catalog.gamepass.com/sigls/v2?" + params
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
