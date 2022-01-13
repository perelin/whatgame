package igdbapi

import (
	"os"

	"github.com/Henry-Sarabia/igdb/v2"
)

var ClientID string
var AccessToken string

func init() {
	ClientID = os.Getenv("IGDB_CLIENT_ID")
	AccessToken = os.Getenv("IGDB_ACCESS_TOKEN")
}

func GetIGDBGameSearchResults(gameTitel string) ([]*igdb.Game, error) {
	igdbClient := igdb.NewClient(ClientID, AccessToken, nil)
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
			"first_release_date",
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

func GetIGDBGameByID(gameID int) (*igdb.Game, error) {
	igdbClient := igdb.NewClient(ClientID, AccessToken, nil)
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
