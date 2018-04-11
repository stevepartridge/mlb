package mlb_test

import (
	"testing"
	"time"

	"github.com/stevepartridge/mlb"
)

func Test_GetGamesForDate_Success(t *testing.T) {
	mlbApi, err := mlb.New()
	if err != nil {
		t.Error("Should not see error:" + err.Error())
	}
	if mlbApi == nil {
		t.Error("mlbApi should not be nil.")
	}

	mlbApi.Debug = true

	date, _ := time.Parse("2006/01/02", "2017/05/17")

	games, err := mlbApi.GetGamesByDate(date)
	if err != nil {
		t.Error("Should not error: " + err.Error())
	}

	if len(games) == 0 {
		t.Error("Games should be more than zero.")
	}
}
