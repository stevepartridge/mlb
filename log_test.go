package mlb

import (
	"errors"
	"testing"
)

func Test_IfErrorSuccess(t *testing.T) {
	err := errors.New("Dummy error")
	ifError(err)
}
