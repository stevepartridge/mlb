package mlb

import (
	"errors"
	"fmt"
)

func (m *Mlb) GetVenues() ([]Venue, error) {
	resp, err := m.Call("/venues", nil)
	if err != nil {
		return []Venue{}, err
	}

	return resp.Venues, nil
}

func (m *Mlb) GetVenue(venueID int) (Venue, error) {
	resp, err := m.Call(fmt.Sprintf("/venues/%d", venueID), nil)
	if err != nil {
		return Venue{}, err
	}

	for _, v := range resp.Venues {
		if v.Id == venueID {
			return v, nil
		}
	}

	return Venue{}, errors.New("not found")
}
