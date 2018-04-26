package mlb

import (
	"strconv"
)

const (
	AmericanLeagueEast                 = 201
	AmericanLeagueCentral              = 202
	AmericanLeagueWest                 = 200
	NationalLeagueEast                 = 204
	NationalLeagueCentral              = 205
	NationalLeagueWest                 = 203
	InternationalLeagueNorth           = 219
	InternationalLeagueSouth           = 220
	InternationalLeagueWest            = 221
	PacificCoastLeagueAmericanNorthern = 232
	PacificCoastLeagueAmericanSouthern = 234
	PacificCoastLeaguePacificNorthern  = 231
	PacificCoastLeaguePacificSouthern  = 233
	MexicanLeagueNorte                 = 222
	EasternLeagueWestern               = 213
	SouthernLeagueNorth                = 239
	SouthernLeagueSouth                = 240
	TexasLeagueNorth                   = 241
	TexasLeagueSouth                   = 242
	CaliforniaLeagueNorth              = 208
	CaliforniaLeagueSouth              = 209
	CarolinaLeagueNorthern             = 210
	CarolinaLeagueSouthern             = 211
	FloridaStateLeagueNorth            = 214
	FloridaStateLeagueSouth            = 215
	MidwestLeagueEastern               = 224
	MidwestLeagueWestern               = 225
	SouthAtlanticLeagueNorthern        = 237
	SouthAtlanticLeagueSouthern        = 238
	NewYorkPennLeagueMcNamara          = 228
	NewYorkPennLeaguePinckney          = 229
	NewYorkPennLeagueStedler           = 230
	NorthwestLeagueNorth               = 226
	NorthwestLeagueSouth               = 227
	PioneerLeagueNorth                 = 235
	PioneerLeagueSouth                 = 236
	AppalachianLeagueEast              = 206
	AppalachianLeagueWest              = 207
	ArizonaLeagueEast                  = 560
	ArizonaLeagueCentral               = 570
	ArizonaLeagueWest                  = 561
	GulfCoastLeagueEast                = 216
	GulfCoastLeagueNortheast           = 217
	GulfCoastLeagueNorthwest           = 615
	GulfCoastLeagueSouth               = 218
	DominicanSummerLeagueNorth         = 247
	DominicanSummerLeagueSouth         = 248
	DominicanSummerLeagueNorthwest     = 249
	DominicanSummerLeagueBaseballCity  = 250
	DominicanSummerLeagueSanPedro      = 246
	DominicanSummerLeagueNortheast     = 401
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
