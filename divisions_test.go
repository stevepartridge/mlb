package mlb_test

import (
	"testing"

	"github.com/stevepartridge/mlb"
)

const (
	invalidDivisionId1 = "aaa"
)

func Test_GetDivisions_Success(t *testing.T) {

	divisions, err := mlbApi.GetDivisions()
	if err != nil {
		t.Error("Should not error when retrieving divisions")
	}

	if len(divisions) == 0 {
		t.Error("Should be more than zero")
	}

}

func Test_GetDivision_Success(t *testing.T) {

	divisions, err := mlbApi.GetDivisions(mlb.NationalLeagueWest)
	if err != nil {
		t.Error("Should not error when retrieving divisions")
	}

	if len(divisions) != 1 {
		t.Error("Should be exactly 1 division found.")
	}

}

func Test_GetSpecificDivisions_Success(t *testing.T) {

	divisions, err := mlbApi.GetDivisions(mlb.NationalLeagueWest, mlb.AmericanLeagueEast)
	if err != nil {
		t.Error("Should not error when retrieving divisions")
	}

	if len(divisions) != 1 {
		t.Error("Should be exactly 1 division found.")
	}

}

// func Test_GetDivision_Failure_InvalidDivisionId(t *testing.T) {

// 	divisions, err := mlbApi.GetDivisions(invalidDivisionId1)
// 	if err == nil {
// 		t.Error("Should be error when retrieving invalid division")
// 	}

// 	if len(divisions) != 0 {
// 		t.Error("Zero divisions should be found.")
// 	}

// }

func Test_GetEachDivision_Success(t *testing.T) {
	testDivisions := map[int]string{
		mlb.AmericanLeagueEast:                "American League East",
		mlb.AmericanLeagueCentral:             "American League Central",
		mlb.AmericanLeagueWest:                "American League West",
		mlb.NationalLeagueEast:                "National League East",
		mlb.NationalLeagueCentral:             "National League Central",
		mlb.NationalLeagueWest:                "National League West",
		mlb.InternationalLeagueEast:           "International League East",
		mlb.InternationalLeagueWest:           "International League West",
		mlb.PacificCoastLeagueWest:            "Pacific Coast League West",
		mlb.PacificCoastLeagueEast:            "Pacific Coast League East",
		mlb.MexicanLeagueNorte:                "Mexican League Norte",
		mlb.EasternLeagueSouthwest:            "Eastern League Southwest",
		mlb.SouthernLeagueNorth:               "Southern League North",
		mlb.SouthernLeagueSouth:               "Southern League South",
		mlb.TexasLeagueNorth:                  "Texas League North",
		mlb.TexasLeagueSouth:                  "Texas League South",
		mlb.CaliforniaLeagueNorth:             "California League North",
		mlb.CaliforniaLeagueSouth:             "California League South",
		mlb.CarolinaLeagueNorth:               "Carolina League North",
		mlb.CarolinaLeagueSouth:               "Carolina League South",
		mlb.FloridaStateLeagueEast:            "Florida State League East",
		mlb.FloridaStateLeagueWest:            "Florida State League West",
		mlb.MidwestLeagueEastern:              "Midwest League East",
		mlb.MidwestLeagueWestern:              "Midwest League West",
		mlb.SouthAtlanticLeagueNorth:          "South Atlantic League North",
		mlb.SouthAtlanticLeagueSouth:          "South Atlantic League South",
		mlb.ArizonaLeagueEast:                 "Arizona Complex League East",
		mlb.ArizonaLeagueCentral:              "Arizona Complex League Central",
		mlb.ArizonaLeagueWest:                 "Arizona Complex League West",
		mlb.DominicanSummerLeagueNorth:        "Dominican Summer League North",
		mlb.DominicanSummerLeagueSouth:        "Dominican Summer League South",
		mlb.DominicanSummerLeagueNorthwest:    "Dominican Summer League Northwest",
		mlb.DominicanSummerLeagueBaseballCity: "Dominican Summer League Baseball City",
		mlb.DominicanSummerLeagueSanPedro:     "Dominican Summer League San Pedro",
		mlb.DominicanSummerLeagueNortheast:    "Dominican Summer League Northeast",
	}

	for id, division := range testDivisions {
		t.Run(division, func(t *testing.T) {
			divisions, err := mlbApi.GetDivisions(id)
			if err != nil {
				t.Fatal("Should not error when retrieving divisions")
			}

			if len(divisions) != 1 {
				t.Fatal("Should be exactly 1 division found.")
			}

			if divisions[0].Name != division {
				t.Errorf("Division found (%s) does not match expected division (%s)",
					divisions[0].Name,
					division,
				)
			}
		})
	}
}
