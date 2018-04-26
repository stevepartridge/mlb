package mlb

import (
	"strconv"
)

const (
	Angels       = 108 // Los Angeles Angels
	Diamondbacks = 109 // Arizona Diamondbacks
	Orioles      = 110 // Baltimore Orioles
	RedSox       = 111 // Boston Red Sox
	Cubs         = 112 // Chicago Cubs
	Reds         = 113 // Cincinnati Reds
	Indians      = 114 // Cleveland Indians
	Rockies      = 115 // Colorado Rockies
	Tigers       = 116 // Detroit Tigers
	Astros       = 117 // Houston Astros
	Royals       = 118 // Kansas City Royals
	Dodgers      = 119 // Los Angeles Dodgers
	Nationals    = 120 // Washington Nationals
	Mets         = 121 // New York Mets
	Athletics    = 133 // Oakland Athletics
	Pirates      = 134 // Pittsburgh Pirates
	Padres       = 135 // San Diego Padres
	Mariners     = 136 // Seattle Mariners
	Giants       = 137 // San Francisco Giants
	Cardinals    = 138 // St. Louis Cardinals
	Rays         = 139 // Tampa Bay Rays
	Rangers      = 140 // Texas Rangers
	BlueJays     = 141 // Toronto Blue Jays
	Twins        = 142 // Minnesota Twins
	Phillies     = 143 // Philadelphia Phillies
	Braves       = 144 // Atlanta Braves
	WhiteSox     = 145 // Chicago White Sox
	Marlins      = 146 // Miami Marlins
	Yankees      = 147 // New York Yankees
	Brewers      = 158 // Milwaukee Brewers
)

func (m *Mlb) GetTeams(teamIds ...int) ([]Team, error) {

	params := map[string]string{}

	if params["sportId"] == "" {
		params["sportId"] = "1"
	}

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
