package mlb

type Linescore struct {
	Runs   LinescoreItem `json:"r"`
	Hits   LinescoreItem `json:"h"`
	Errors LinescoreItem `json:"e"`

	Strikeouts  LinescoreItem `json:"so"`
	Stolenbases LinescoreItem `json:"sb"`
	Homeruns    LinescoreItem `json:"hr"`

	Innings interface{} `json:"inning"`
}

type LinescoreItem struct {
	Home string `json:"home"`
	Away string `json:"away"`
}
