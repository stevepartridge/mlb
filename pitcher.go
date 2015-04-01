package mlb

type Pitcher struct {
	Id                string `json:"id"`
	FirstName         string `json:"first"`
	LastName          string `json:"last"`
	NameDisplayRoster string `json:"name_display_roster"`
	Number            string `json:"number"`
	Wins              string `json:"wins"`
	Losses            string `json:"losses"`
	Era               string `json:"era"`
}
