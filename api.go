package mlb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// schedule?lang=en&sportId=1&season=2018&startDate=2018-08-01&endDate=2018-08-31&teamId=119&eventTypes=primary&scheduleTypes=games

func (m *Mlb) Call(endpoint string, query map[string]string) (Response, error) {

	result := Response{}

	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	if len(endpoint) < 2 {
		return result, errors.New("Invalid endpoint")
	}

	if string(endpoint[0]) == "/" {
		endpoint = endpoint[1:]
	}

	u := fmt.Sprintf("%s/%s",
		baseApiURL,
		endpoint,
	)

	if query["lang"] == "" {
		query["lang"] = "en"
	}

	if query["sportId"] == "" {
		query["sportId"] = "1"
	}

	for k, v := range query {
		m.log(" - ", k, v)
		if !strings.Contains(u, "?") {
			u = u + "?" + k + "=" + url.QueryEscape(v)
			continue
		}
		u = u + "&" + k + "=" + url.QueryEscape(v)
	}

	_url, err := url.Parse(u)

	if err != nil {
		ifError(err)
		return result, err
	}

	m.log("calling:", _url.String())

	resp, err := client.Get(_url.String())

	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		ifError(err)
		return result, err
	}

	if body != nil {
		err = json.Unmarshal([]byte(body), &result)
		if err != nil {
			ifError(err)
			return result, err
		}

	}

	m.log("status code:", resp.StatusCode)
	switch resp.StatusCode {
	case 200:
		return result, nil
	case 404:
		return result, errors.New("Error (Not Found): " + result.Message)
	default:
		return result, errors.New("Invalid response code from MLB.com, StatusCode: " + strconv.Itoa(resp.StatusCode))
	}

	// m.log(" - body: \n %s", string(body))

	return result, nil

}
