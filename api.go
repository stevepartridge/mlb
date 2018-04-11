package mlb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	baseApiURL = "https://statsapi.mlb.com/api/v1"
)

// schedule?lang=en&sportId=1&season=2018&startDate=2018-08-01&endDate=2018-08-31&teamId=119&eventTypes=primary&scheduleTypes=games

type Mlb struct {
	Debug bool
}

type Response struct {
	Copyright            string `json:"copyright"`
	TotalItems           int    `json:"totalItems"`
	TotalEvents          int    `json:"totalEvents"`
	TotalGames           int    `json:"totalGames"`
	TotalGamesInProgress int    `json:"totalGamesInProgress"`
	Dates                []Date `json:"dates,omitempty"`
}

func (m *Mlb) call(endpoint string, query map[string]string) (Response, error) {

	result := Response{}

	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	_url := fmt.Sprintf("%s/%s?sportId=1", baseApiURL, endpoint)

	if query["lang"] == "" {
		query["lang"] = "en"
	}

	for k, v := range query {
		m.log(" - ", k, v)
		_url = _url + "&" + k + "=" + url.QueryEscape(v)
	}

	m.log("calling:", _url)

	resp, err := client.Get(_url)

	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		ifError(err)
		return result, err
	}

	m.log("status code:", resp.StatusCode)
	if resp.StatusCode != 200 {
		return result, errors.New("Invalid response code from MLB.com, StatusCode: " + strconv.Itoa(resp.StatusCode))
	}

	m.log(" - body: \n%s", string(body))

	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		ifError(err)
		return result, err
	}

	return result, nil

}
