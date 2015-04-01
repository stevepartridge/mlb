package mlb

import (
	"encoding/json"
	"github.com/stevepartridge/go/log"
)

func ResponseDataToInterface(data interface{}, v interface{}) {
	str, err := json.Marshal(data)
	log.IfError(err)
	err = json.Unmarshal(str, &v)
	log.IfError(err)
	return
}
