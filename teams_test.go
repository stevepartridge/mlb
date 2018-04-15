package mlb_test

import (
	"testing"

	"github.com/stevepartridge/mlb"
)

const (
	invalidTeamId1 = 9999
)

func Test_GetTeams_Success(t *testing.T) {

	teams, err := mlbApi.GetTeams()
	if err != nil {
		t.Error("Should not error when retrieving teams")
	}

	if len(teams) == 0 {
		t.Error("Should be more than zero")
	}

}

func Test_GetTeam_Success(t *testing.T) {

	teams, err := mlbApi.GetTeams(mlb.Dodgers)
	if err != nil {
		t.Error("Should not error when retrieving teams")
	}

	if len(teams) != 1 {
		t.Error("Should be exactly 1 team found.")
	}

}

func Test_GetSpecificTeams_Success(t *testing.T) {

	teams, err := mlbApi.GetTeams(mlb.Dodgers, mlb.RedSox)
	if err != nil {
		t.Error("Should not error when retrieving teams")
	}

	if len(teams) != 1 {
		t.Error("Should be exactly 1 team found.")
	}

}

func Test_GetTeam_Failure_InvalidTeamId(t *testing.T) {

	teams, err := mlbApi.GetTeams(invalidTeamId1)
	if err == nil {
		t.Error("Should be error when retrieving invalid team")
	}

	if len(teams) != 0 {
		t.Error("Zero teams should be found.")
	}

}

func Test_GetEachTeam_Success(t *testing.T) {
	testTeams := map[int]string{
		mlb.Angels:       "Los Angeles Angels",
		mlb.Diamondbacks: "Arizona Diamondbacks",
		mlb.Orioles:      "Baltimore Orioles",
		mlb.RedSox:       "Boston Red Sox",
		mlb.Cubs:         "Chicago Cubs",
		mlb.Reds:         "Cincinnati Reds",
		mlb.Indians:      "Cleveland Indians",
		mlb.Rockies:      "Colorado Rockies",
		mlb.Tigers:       "Detroit Tigers",
		mlb.Astros:       "Houston Astros",
		mlb.Royals:       "Kansas City Royals",
		mlb.Dodgers:      "Los Angeles Dodgers",
		mlb.Nationals:    "Washington Nationals",
		mlb.Mets:         "New York Mets",
		mlb.Athletics:    "Oakland Athletics",
		mlb.Pirates:      "Pittsburgh Pirates",
		mlb.Padres:       "San Diego Padres",
		mlb.Mariners:     "Seattle Mariners",
		mlb.Giants:       "San Francisco Giants",
		mlb.Cardinals:    "St. Louis Cardinals",
		mlb.Rays:         "Tampa Bay Rays",
		mlb.Rangers:      "Texas Rangers",
		mlb.BlueJays:     "Toronto Blue Jays",
		mlb.Twins:        "Minnesota Twins",
		mlb.Phillies:     "Philadelphia Phillies",
		mlb.Braves:       "Atlanta Braves",
		mlb.WhiteSox:     "Chicago White Sox",
		mlb.Marlins:      "Miami Marlins",
		mlb.Yankees:      "New York Yankees",
		mlb.Brewers:      "Milwaukee Brewers",
	}

	for id, team := range testTeams {

		teams, err := mlbApi.GetTeams(id)
		if err != nil {
			t.Error("Should not error when retrieving teams")
		}

		if len(teams) != 1 {
			t.Error("Should be exactly 1 team found.")
		}

		if teams[0].Name != team {
			t.Errorf("Team found (%s) does not match expected team (%s)",
				teams[0].Name,
				team,
			)
		}
	}

}
