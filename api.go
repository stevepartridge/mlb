package mlb

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stevepartridge/go/log"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	BASE_MLB_API_URL = "http://gd2.mlb.com/components/game/mlb"
)

type MlbApi struct{}

type MlbApiResponse struct {
	Subject   string      `json:"subject"`
	Copyright string      `json:"copyright"`
	Data      interface{} `json:"data"`
}

func (a *MlbApi) call(endpoint string) (MlbApiResponse, error) {
	client := &http.Client{}

	var body []byte
	var err error

	log.Debug(BASE_MLB_API_URL)
	log.Debug(endpoint)
	req, _ := http.NewRequest("GET", BASE_MLB_API_URL+endpoint, bytes.NewReader(body))

	resp, err := client.Do(req)
	if err != nil {
		return MlbApiResponse{}, err
	}

	defer resp.Body.Close()

	result := MlbApiResponse{}

	if resp.StatusCode != 200 {
		return result, errors.New("Invalid response code from MLB.com, StatusCode: " + strconv.Itoa(resp.StatusCode))
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.IfError(err)
	}

	err = json.Unmarshal([]byte(contents), &result)
	if err != nil {
		log.Error(err)
		log.Debug(string(contents))
		return MlbApiResponse{}, err
	}

	return result, nil
}
