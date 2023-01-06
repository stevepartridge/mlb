package mlb_test

import "testing"

func Test_GetVenues_Success(t *testing.T) {

	venues, err := mlbApi.GetVenues()
	if err != nil {
		t.Error("Should not error when retrieving venues")
	}

	if len(venues) == 0 {
		t.Error("Should be more than zero")
	}

}

func Test_GetVenue_Success(t *testing.T) {
	venue, err := mlbApi.GetVenue(10) // "Oakland Coliseum"
	if err != nil {
		t.Error("Should not error when retrieving venue")
	}

	if venue.Name != "Oakland Coliseum" {
		t.Errorf("Expected `Oakland Coliseum`, got %q", venue.Name)

	}
}

func Test_GetVenue_Failure_InvalidVenueID(t *testing.T) {
	_, err := mlbApi.GetVenue(8675309)
	if err == nil {
		t.Error("Should be error when retrieving invalid venue")
	}
}
