package mlb

type Response struct {
	Copyright     string `json:"copyright"`
	Message       string `json:"message,omitempty"`
	MessageNumber int    `json:"messageNumber,omitempty"`

	TotalItems           int    `json:"totalItems,omitempty"`
	TotalEvents          int    `json:"totalEvents,omitempty"`
	TotalGames           int    `json:"totalGames,omitempty"`
	TotalGamesInProgress int    `json:"totalGamesInProgress,omitempty"`
	Dates                []Date `json:"dates,omitempty"`
	Teams                []Team `json:"teams,omitempty"`
}

type Date struct {
	Date                 string `json:"date"`
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
	DoubleHeader           string  `json:"doubleHeader"`
	GamedayType            string  `json:"gamedayType"`
	Tiebreaker             string  `json:"tiebreaker"`
	GameNumber             int     `json:"gameNumber"`
	CalendarEventID        string  `json:"calendarEventID"`
	SeasonDisplay          string  `json:"seasonDisplay"`
	DayNight               string  `json:"dayNight"`
	ScheduledInnings       int     `json:"scheduledInnings"`
	GamesInSeries          int     `json:"gamesInSeries"`
	SeriesGameNumber       int     `json:"seriesGameNumber"`
	SeriesDescription      string  `json:"seriesDescription"`
	RecordSource           string  `json:"recordSource"`
	IfNecessary            string  `json:"ifNecessary"`
	IfNecessaryDescription string  `json:"ifNecessaryDescription"`
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
	LeagueRecord LeagueRecord `json:"leagueRecord"`
	Team         Team         `json:"team"`
	IsWinner     bool         `json:"isWinner"`
	SplitSquad   bool         `json:"splitSquad"`
	SeriesNumber int          `json:"seriesNumber"`
}

type Team struct {
	Id              int      `json:"id"`
	Name            string   `json:"name"`
	Link            string   `json:"link"`
	Venue           Venue    `json:"venue"`
	TeamCode        string   `json:"teamCode"`
	FileCode        string   `json:"fileCode"`
	Abbreviation    string   `json:"abbreviation"`
	TeamName        string   `json:"teamName"`
	LocationName    string   `json:"locationName"`
	FirstYearOfPlay string   `json:"firstYearOfPlay"`
	League          League   `json:"league"`
	Division        Division `json:"division"`
	Sport           Sport    `json:"sport"`
	ShortName       string   `json:"shortName"`
	SpringLeague    League   `json:"springLeague"`
	Active          bool     `json:"active"`
}

type LeagueRecord struct {
	Wins    int    `json:"wins"`
	Losses  int    `json:"losses"`
	Percent string `json:"pct"`
}

type Sport struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type League struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Link         string `json:"link"`
	Abbreviation string `json:"abbreviation"`
}

type Division struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}
