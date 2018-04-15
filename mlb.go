package mlb

const (
	baseApiURL = "https://statsapi.mlb.com/api/v1"
)

type Mlb struct {
	Debug bool
}

func New() (*Mlb, error) {
	return &Mlb{}, nil
}
