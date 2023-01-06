package mlb

import (
	"strconv"
)

const (
	AmericanLeagueEast                = 201
	AmericanLeagueCentral             = 202
	AmericanLeagueWest                = 200
	NationalLeagueEast                = 204
	NationalLeagueCentral             = 205
	NationalLeagueWest                = 203
	InternationalLeagueEast           = 219
	InternationalLeagueWest           = 221
	PacificCoastLeagueWest            = 231
	PacificCoastLeagueEast            = 233
	MexicanLeagueNorte                = 222
	EasternLeagueSouthwest            = 213
	SouthernLeagueNorth               = 239
	SouthernLeagueSouth               = 240
	TexasLeagueNorth                  = 241
	TexasLeagueSouth                  = 242
	CaliforniaLeagueNorth             = 208
	CaliforniaLeagueSouth             = 209
	CarolinaLeagueNorth               = 210
	CarolinaLeagueSouth               = 211
	FloridaStateLeagueEast            = 214
	FloridaStateLeagueWest            = 215
	MidwestLeagueEastern              = 224
	MidwestLeagueWestern              = 225
	SouthAtlanticLeagueNorth          = 237
	SouthAtlanticLeagueSouth          = 238
	ArizonaLeagueEast                 = 560
	ArizonaLeagueCentral              = 570
	ArizonaLeagueWest                 = 561
	GulfCoastLeagueEast               = 216
	GulfCoastLeagueNortheast          = 217
	GulfCoastLeagueNorthwest          = 615
	GulfCoastLeagueSouth              = 218
	DominicanSummerLeagueNorth        = 247
	DominicanSummerLeagueSouth        = 248
	DominicanSummerLeagueNorthwest    = 249
	DominicanSummerLeagueBaseballCity = 250
	DominicanSummerLeagueSanPedro     = 246
	DominicanSummerLeagueNortheast    = 401
)

// GetDivisions returns all or specific divisions based on if any divisionIds are provided or not
func (m *Mlb) GetDivisions(divisionIds ...int) ([]Division, error) {

	params := map[string]string{}

	if len(divisionIds) > 0 {
		for i := range divisionIds {

			params["divisionId"] = strconv.Itoa(divisionIds[i])
			if i != len(divisionIds)-1 {
				params["divisionId"] = params["divisionId"] + ","
			}

		}
	}

	resp, err := m.Call("/divisions", params)

	if err != nil {
		return []Division{}, err
	}

	return resp.Divisions, nil
}
