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
