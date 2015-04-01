package mlb

type Status struct {
	Inning      string `json"inning"`
	InningState string `json"inning_state"`
	Strikes     string `json:"s"`
	Balls       string `json:"b"`
	Outs        string `json:"o"`
	TopInning   string `json:"top_inning"`
	Status      string `json:"status"`
	Ind         string `json:"ind"`
	Reason      string `json:"reason"`
}

type Game struct {
	GameId           string `json:"id"`
	Time             string `json:"time"`
	AmPm             string `json:"ampm"`
	TimeDate         string `json:"time_date"`
	OriginalDate     string `json:"original_date"`
	Day              string `json:"day"`
	VenueId          string `json:"venue_id"`
	TimeZone         string `json:"time_zone"`
	ScheduledInnings string `json:"scheduled_innings"`
	GameDayId        string `json:"gameday"`

	Status Status `json:"status"`

	Linescore Linescore `json:"linescore"`

	WinningPitcher Pitcher `json:"winning_pitcher"`
	LosingPitcher  Pitcher `json:"losing_pitcher"`

	HomeTeamId            string `json:"home_team_id"`
	HomeCode              string `json:"home_code"`
	HomeFileCode          string `json:"home_file_code"`
	HomeTeamName          string `json:"home_team_name"`
	HomeNameAbbreviated   string `json:"home_name_abbrev"`
	HomeTeamCity          string `json:"home_team_city"`
	HomeWin               string `json:"home_win"`
	HomeLoss              string `json:"home_loss"`
	HomeGamesBack         string `json:"home_games_back"`
	HomeGamesBackWildcard string `json:"home_games_back_wildcard"`
	HomeTime              string `json:"home_time"`
	HomeAmPm              string `json:"home_ampm"`
	HomeTimeZone          string `json:"home_time_zone"`
	HomeTimeZoneOffset    string `json:"time_zone_hm_lg"`
	HomeDivision          string `json:"home_division"`
	HomeLeagueId          string `json:"home_league_id"`
	HomeSportCode         string `json:"home_sport_code"`

	AwayTeamId            string `json:"away_team_id"`
	AwayCode              string `json:"away_code"`
	AwayFileCode          string `json:"away_file_code"`
	AwayTeamName          string `json:"away_team_name"`
	AwayNameAbbreviated   string `json:"away_name_abbrev"`
	AwayTeamCity          string `json:"away_team_city"`
	AwayWin               string `json:"away_win"`
	AwayLoss              string `json:"away_loss"`
	AwayGamesBack         string `json:"away_games_back"`
	AwayGamesBackWildcard string `json:"away_games_back_wildcard"`
	AwayTime              string `json:"away_time"`
	AwayAmPm              string `json:"away_ampm"`
	AwayTimeZone          string `json:"away_time_zone"`
	AwayTimeZoneOffset    string `json:"time_zone_aw_lg"`
	AwayDivision          string `json:"away_division"`
	AwayLeagueId          string `json:"away_league_id"`
	AwaySportCode         string `json:"away_sport_code"`
}

type GamesResult struct {
	GamesData struct {
		Games []Game `json:"game"`
	} `json:"games"`
}

type LinescoreResult struct {
	Game Game `json:game`
}
