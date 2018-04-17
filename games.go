package mlb

import (
	"errors"
	"strconv"
	"time"
)

// Get games by specific date.  Date is assumed to be relevant local to game timezone.
func (m *Mlb) GetGamesByDate(date time.Time) ([]Game, error) {
	return m.GetGamesByDateRange(date, date)
}

// Get games by specific date range.  Date is assumed to be relevant local to game timezone.
func (m *Mlb) GetGamesByDateRange(start, end time.Time) ([]Game, error) {
	return m.GetGames(start, end, 0)
}

// Get games by date for a specific team.  Date is assumed to be relevant local to game timezone.
func (m *Mlb) GetGamesByDateForTeam(date time.Time, teamId int) ([]Game, error) {
	return m.GetGames(date, date, teamId)
}

// Get games in date range for specific team.  Date is assumed to be relevant local to game timezone.
func (m *Mlb) GetGamesByDateRangeForTeam(start, end time.Time, teamId int) ([]Game, error) {
	return m.GetGames(start, end, teamId)
}

func (m *Mlb) GetGames(start, end time.Time, teamId int) ([]Game, error) {

	// &season=2018&startDate=2018-08-01&endDate=2018-08-31&teamId=119&eventTypes=primary&scheduleTypes=games

	if start.IsZero() {
		return nil, errors.New("Invalid start date")
	}

	if end.IsZero() {
		return nil, errors.New("Invalid end date")
	}

	params := map[string]string{
		"season":        start.Format("2006"),
		"eventTypes":    "primary",
		"scheduleTypes": "games",
		"startDate":     start.Format("2006-01-02"),
		"endDate":       end.Format("2006-01-02"),
	}

	if teamId > 0 {
		params["teamId"] = strconv.Itoa(teamId)
	}

	resp, err := m.Call("/schedule", params)
	if err != nil {
		ifError(err)
		return []Game{}, err
	}

	if len(resp.Dates) == 0 {
		return []Game{}, err
	}

	return resp.Dates[0].Games, nil
}
