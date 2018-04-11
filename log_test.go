package mlb

import (
	"errors"
	"testing"
)

// This is to appease the ocd of 100% coverage, in theory
func Test_IfErrorSuccess(t *testing.T) {
	err := errors.New("Dummy error")
	ifError(err)
}
