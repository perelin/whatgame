package igdbapi

import (
	"os"

	"github.com/Henry-Sarabia/igdb/v2"
)

var igdbClient *igdb.Client

func init() {
	igdbClient = igdb.NewClient(os.Getenv("IGDB_CLIENT_ID"),
		os.Getenv("IGDB_ACCESS_TOKEN"),
		nil)
}

func GetIGDBGameSearchResults(gameTitel string) ([]*igdb.Game, error) {
	//igdbClient := igdb.NewClient(ClientID, AccessToken, nil)
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
			"involved_companies",
			"websites",
			"genres"))
	return games, err
}

func GetIGDBInvolvedCompanyByID(involvedCompanyID int) (*igdb.InvolvedCompany, error) {
	//igdbClient := igdb.NewClient(ClientID, AccessToken, nil)
	company, err := igdbClient.InvolvedCompanies.Get(involvedCompanyID, igdb.SetFields(
		"company",
		"developer",
		"game",
		"porting",
		"publisher",
		"supporting",
	))
	return company, err
}

func GetIGDBCompany(companyID int) (*igdb.Company, error) {
	//igdbClient := igdb.NewClient(ClientID, AccessToken, nil)
	company, err := igdbClient.Companies.Get(companyID,
		igdb.SetFields("name",
			"url",
			"slug"))
	return company, err
}

func GetIGDBGameByID(gameID int) (*igdb.Game, error) {
	//igdbClient := igdb.NewClient(ClientID, AccessToken, nil)
	game, err := igdbClient.Games.Get(gameID,
		igdb.SetOrder("hypes", igdb.OrderDescending),
		igdb.SetFilter("category", igdb.OpEquals, "0"),
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
