package mlb_test

import (
	"testing"
)

func Test_ApiCall_Failure_InvalidPathNotFound(t *testing.T) {
	// mlbApi, _ := mlbNew()

	_, err := mlbApi.Call("/asdf", map[string]string{})
	if err == nil {
		t.Error("Should be error.")
	}
}

func Test_ApiCall_Failure_InvalidPathTooShort(t *testing.T) {
	// mlbApi, _ := New()

	_, err := mlbApi.Call("a", map[string]string{})
	if err == nil {
		t.Error("Should be error.")
	}
}

func Test_ApiCall_Failure_InvalidPathParsing(t *testing.T) {
	// mlbApi, _ := New()

	_, err := mlbApi.Call("https://...a", map[string]string{})
	if err == nil {
		t.Error("Should be error.")
	}
}
