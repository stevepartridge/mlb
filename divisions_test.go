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
		mlb.AmericanLeagueEast:                 "American League East",
		mlb.AmericanLeagueCentral:              "American League Central",
		mlb.AmericanLeagueWest:                 "American League West",
		mlb.NationalLeagueEast:                 "National League East",
		mlb.NationalLeagueCentral:              "National League Central",
		mlb.NationalLeagueWest:                 "National League West",
		mlb.InternationalLeagueNorth:           "International League North",
		mlb.InternationalLeagueSouth:           "International League South",
		mlb.InternationalLeagueWest:            "International League West",
		mlb.PacificCoastLeagueAmericanNorthern: "Pacific Coast League American Northern",
		mlb.PacificCoastLeagueAmericanSouthern: "Pacific Coast League American Southern",
		mlb.PacificCoastLeaguePacificNorthern:  "Pacific Coast League Pacific Northern",
		mlb.PacificCoastLeaguePacificSouthern:  "Pacific Coast League Pacific Southern",
		mlb.MexicanLeagueNorte:                 "Mexican League Norte",
		mlb.EasternLeagueWestern:               "Eastern League Western",
		mlb.SouthernLeagueNorth:                "Southern League North",
		mlb.SouthernLeagueSouth:                "Southern League South",
		mlb.TexasLeagueNorth:                   "Texas League North",
		mlb.TexasLeagueSouth:                   "Texas League South",
		mlb.CaliforniaLeagueNorth:              "California League North",
		mlb.CaliforniaLeagueSouth:              "California League South",
		mlb.CarolinaLeagueNorthern:             "Carolina League Northern",
		mlb.CarolinaLeagueSouthern:             "Carolina League Southern",
		mlb.FloridaStateLeagueNorth:            "Florida State League North",
		mlb.FloridaStateLeagueSouth:            "Florida State League South",
		mlb.MidwestLeagueEastern:               "Midwest League Eastern",
		mlb.MidwestLeagueWestern:               "Midwest League Western",
		mlb.SouthAtlanticLeagueNorthern:        "South Atlantic League Northern",
		mlb.SouthAtlanticLeagueSouthern:        "South Atlantic League Southern",
		mlb.NewYorkPennLeagueMcNamara:          "New York-Penn League McNamara",
		mlb.NewYorkPennLeaguePinckney:          "New York-Penn League Pinckney",
		mlb.NewYorkPennLeagueStedler:           "New York-Penn League Stedler",
		mlb.NorthwestLeagueNorth:               "Northwest League North",
		mlb.NorthwestLeagueSouth:               "Northwest League South",
		mlb.PioneerLeagueNorth:                 "Pioneer League North",
		mlb.PioneerLeagueSouth:                 "Pioneer League South",
		mlb.AppalachianLeagueEast:              "Appalachian League East",
		mlb.AppalachianLeagueWest:              "Appalachian League West",
		mlb.ArizonaLeagueEast:                  "Arizona League East",
		mlb.ArizonaLeagueCentral:               "Arizona League Central",
		mlb.ArizonaLeagueWest:                  "Arizona League West",
		mlb.GulfCoastLeagueEast:                "Gulf Coast League East",
		mlb.GulfCoastLeagueNortheast:           "Gulf Coast League Northeast",
		mlb.GulfCoastLeagueNorthwest:           "Gulf Coast League Northwest",
		mlb.GulfCoastLeagueSouth:               "Gulf Coast League South",
		mlb.DominicanSummerLeagueNorth:         "Dominican Summer League North",
		mlb.DominicanSummerLeagueSouth:         "Dominican Summer League South",
		mlb.DominicanSummerLeagueNorthwest:     "Dominican Summer League Northwest",
		mlb.DominicanSummerLeagueBaseballCity:  "Dominican Summer League Baseball City",
		mlb.DominicanSummerLeagueSanPedro:      "Dominican Summer League San Pedro",
		mlb.DominicanSummerLeagueNortheast:     "Dominican Summer League Northeast",
	}

	for id, division := range testDivisions {

		divisions, err := mlbApi.GetDivisions(id)
		if err != nil {
			t.Error("Should not error when retrieving divisions")
		}

		if len(divisions) != 1 {
			t.Error("Should be exactly 1 division found.")
		}

		if divisions[0].Name != division {
			t.Errorf("Division found (%s) does not match expected division (%s)",
				divisions[0].Name,
				division,
			)
		}
	}

}
