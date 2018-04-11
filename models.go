package mlb

type Date struct {
	Date                 string `json:""`
	TotalItems           int    `json:"totalItems"`
	TotalEvents          int    `json:"totalEvents"`
	TotalGames           int    `json:"totalGames"`
	TotalGamesInProgress int    `json:"totalGamesInProgress"`
	Games                []Game `json:"games"`
}

type Game struct {
	Id      int                 `json:"gamePk"`
	Link    string              `json:"link"`
	Type    string              `json:"gameType"`
	Season  string              `json:"season"`
	Date    string              `json:"gameDate"`
	Status  GameStatus          `json:"status"`
	Teams   map[string]GameTeam `json:"teams"`
	Venue   Venue               `json:"venue"`
	Content struct {
		Link string `json:"link"`
	} `json:"content"`
	DoubleHeader           string  `json:""`
	GamedayType            string  `json:"gamedayType"`
	Tiebreaker             string  `json:""`
	GameNumber             int     `json:""`
	CalendarEventID        string  `json:""`
	SeasonDisplay          string  `json:""`
	DayNight               string  `json:""`
	ScheduledInnings       int     `json:""`
	GamesInSeries          int     `json:""`
	SeriesGameNumber       int     `json:""`
	SeriesDescription      string  `json:""`
	RecordSource           string  `json:""`
	IfNecessary            string  `json:""`
	IfNecessaryDescription string  `json:""`
	Events                 []Event `json:""`
}

type GameStatus struct {
	AbstractGameState string `json:"abstractGameState"`
	AbstractGameCode  string `json:"abstractGameCode"`
	CodedGameState    string `json:"codedGameState"`
	DetailedState     string `json:"detailedState"`
	StatusCode        string `json:"statusCode"`
}

type Venue struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type Event struct{}

type GameTeam struct {
	LeagueRecord LeagueRecord `json:""`
	Team         Team         `json:""`
	SplitSquad   bool         `json:""`
	SeriesNumber int          `json:""`
}

type Team struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type LeagueRecord struct {
	Wins    int    `json:""`
	Losses  int    `json:""`
	Percent string `json:""`
}
