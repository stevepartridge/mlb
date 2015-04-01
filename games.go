package mlb

import (
	"github.com/stevepartridge/go/log"
	// "regexp"
	"strconv"
	"strings"
	"time"
)

func GetGamesByDate(date time.Time) ([]Game, error) {
	result := make([]Game, 0)

	year := strconv.Itoa(date.Year())

	m := int(date.Month())
	month := strconv.Itoa(m)
	if m < 10 {
		month = "0" + month
	}

	day := strconv.Itoa(date.Day())
	if date.Day() < 10 {
		day = "0" + day
	}

	resp, err := api.call("/year_" + year + "/month_" + month + "/day_" + day + "/master_scoreboard.json")
	if err != nil {
		log.Error(err)
		return result, err
	}

	gamesResult := GamesResult{}
	ResponseDataToInterface(resp.Data, &gamesResult)

	result = gamesResult.GamesData.Games

	return result, nil
}

func GetGameByFileCode(fileCode string) (Game, error) {
	result := Game{}
	log.Debug(fileCode)
	// 2015_03_24_balmlb_pitmlb_1
	parts := strings.Split(fileCode, "_")
	log.Debug(parts)
	year := parts[0]
	month := parts[1]
	day := parts[2]

	resp, err := api.call("/year_" + year + "/month_" + month + "/day_" + day + "/master_scoreboard.json")
	if err != nil {
		log.Error(err)
		return result, err
	}
	gamesResult := GamesResult{}
	ResponseDataToInterface(resp.Data, &gamesResult)

	for _, v := range gamesResult.GamesData.Games {
		if v.GameDayId == fileCode {
			return v, nil
		}
	}

	return result, nil
}
