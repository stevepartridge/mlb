package mlb_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stevepartridge/mlb"
)

var (
	mlbApi *mlb.Mlb
)

func TestMain(m *testing.M) {

	var err error
	mlbApi, err = mlb.New()
	if err != nil {

		fmt.Println("Error - Should not see error:" + err.Error())
		return
	}

	mlbApi.Debug = true

	os.Exit(m.Run())

}

func Test_NewMlbInstance_Success(t *testing.T) {
	mlbApi, err := mlb.New()
	if err != nil {
		t.Error("Should not see error:" + err.Error())
	}
	if mlbApi == nil {
		t.Error("mlbApi should not be nil.")
	}
}

func Test_NewMlbInstanceWithDebug_Success(t *testing.T) {
	mlbApi, err := mlb.New()
	if err != nil {
		t.Error("Should not see error:" + err.Error())
	}
	if mlbApi == nil {
		t.Error("mlbApi should not be nil.")
	}

	mlbApi.Debug = true

	if !mlbApi.Debug {
		t.Error("Debug should be true, because, seriously just set it.")
	}
}
