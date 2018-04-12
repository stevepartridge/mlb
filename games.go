package mlb

import (
	"time"
)

func (m *Mlb) GetGamesByDate(date time.Time) ([]Game, error) {
	return m.GetGamesByDateRange(date, date)
}

func (m *Mlb) GetGamesByDateRange(start, end time.Time) ([]Game, error) {

	// &season=2018&startDate=2018-08-01&endDate=2018-08-31&teamId=119&eventTypes=primary&scheduleTypes=games
	params := map[string]string{
		"season":        start.Format("2006"),
		"eventTypes":    "primary",
		"scheduleTypes": "games",
		"startDate":     start.Format("2006-01-02"),
		"endDate":       end.Format("2006-01-02"),
	}

	resp, err := m.call("/schedule", params)
	if err != nil {
		ifError(err)
		return []Game{}, err
	}

	if len(resp.Dates) == 0 {
		return []Game{}, err
	}

	return resp.Dates[0].Games, nil
}
