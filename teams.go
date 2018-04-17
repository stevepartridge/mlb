package mlb

import (
	"strconv"
)

const (

	// Los Angeles Angels
	Angels = 108

	// Arizona Diamondbacks
	Diamondbacks = 109

	// Baltimore Orioles
	Orioles = 110

	// Boston Red Sox
	RedSox = 111

	// Chicago Cubs
	Cubs = 112

	// Cincinnati Reds
	Reds = 113

	// Cleveland Indians
	Indians = 114

	// Colorado Rockies
	Rockies = 115

	// Detroit Tigers
	Tigers = 116

	// Houston Astros
	Astros = 117

	// Kansas City Royals
	Royals = 118

	// Los Angeles Dodgers
	Dodgers = 119

	// Washington Nationals
	Nationals = 120

	// New York Mets
	Mets = 121

	// Oakland Athletics
	Athletics = 133

	// Pittsburgh Pirates
	Pirates = 134

	// San Diego Padres
	Padres = 135

	// Seattle Mariners
	Mariners = 136

	// San Francisco Giants
	Giants = 137

	// St. Louis Cardinals
	Cardinals = 138

	// Tampa Bay Rays
	Rays = 139

	// Texas Rangers
	Rangers = 140

	// Toronto Blue Jays
	BlueJays = 141

	// Minnesota Twins
	Twins = 142

	// Philadelphia Phillies
	Phillies = 143

	// Atlanta Braves
	Braves = 144

	// Chicago White Sox
	WhiteSox = 145

	// Miami Marlins
	Marlins = 146

	// New York Yankees
	Yankees = 147

	// Milwaukee Brewers
	Brewers = 158
)

func (m *Mlb) GetTeams(teamIds ...int) ([]Team, error) {

	params := map[string]string{}

	if len(teamIds) > 0 {
		for i := range teamIds {

			params["teamId"] = strconv.Itoa(teamIds[i])
			if i != len(teamIds)-1 {
				params["teamId"] = params["teamId"] + ","
			}

		}
	}

	resp, err := m.Call("/teams", params)

	if err != nil {
		return []Team{}, err
	}

	return resp.Teams, nil
}
