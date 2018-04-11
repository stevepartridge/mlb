package mlb_test

import (
	"testing"

	"github.com/stevepartridge/mlb"
)

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
