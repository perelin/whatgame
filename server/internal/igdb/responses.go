package igdb

type IGDBGameResponse struct {
	Games []IGDBGame `json:"games"`
}

type IGDBGame struct {
	ID                    string      `json:"id"`
	AgeRatings            interface{} `json:"age_ratings"`
	AggregatedRating      int         `json:"aggregated_rating"`
	AggregatedRatingCount int         `json:"aggregated_rating_count"`
	AlternativeNames      interface{} `json:"alternative_names"`
	Artworks              interface{} `json:"artworks"`
	Bundles               interface{} `json:"bundles"`
	Category              int         `json:"category"`
	Collection            int         `json:"collection"`
	Cover                 int         `json:"cover"`
	CreatedAt             int         `json:"created_at"`
	Dlcs                  interface{} `json:"dlcs"`
	Expansions            interface{} `json:"expansions"`
	ExternalGames         interface{} `json:"external_games"`
	FirstReleaseDate      int         `json:"first_release_date"`
	Follows               int         `json:"follows"`
	Franchise             int         `json:"franchise"`
	Franchises            interface{} `json:"franchises"`
	GameEngines           interface{} `json:"game_engines"`
	GameModes             interface{} `json:"game_modes"`
	Genres                interface{} `json:"genres"`
	Hypes                 int         `json:"hypes"`
	InvolvedCompanies     interface{} `json:"involved_companies"`
	Keywords              interface{} `json:"keywords"`
	MultiplayerModes      interface{} `json:"multiplayer_modes"`
	Name                  string      `json:"name"`
	ParentGame            int         `json:"parent_game"`
	Platforms             interface{} `json:"platforms"`
	PlayerPerspectives    interface{} `json:"player_perspectives"`
	Rating                int         `json:"rating"`
	RatingCount           int         `json:"rating_count"`
	ReleaseDates          []int       `json:"release_dates"`
	Screenshots           interface{} `json:"screenshots"`
	SimilarGames          interface{} `json:"similar_games"`
	Slug                  string      `json:"slug"`
	StandaloneExpansions  interface{} `json:"standalone_expansions"`
	Status                int         `json:"status"`
	Storyline             string      `json:"storyline"`
	Summary               string      `json:"summary"`
	Tags                  interface{} `json:"tags"`
	Themes                []int       `json:"themes"`
	TotalRating           float64     `json:"total_rating"`
	TotalRatingCount      int         `json:"total_rating_count"`
	UpdatedAt             int         `json:"updated_at"`
	URL                   string      `json:"url"`
	VersionParent         int         `json:"version_parent"`
	VersionTitle          string      `json:"version_title"`
	Videos                interface{} `json:"videos"`
	Websites              []int       `json:"websites"`
}
