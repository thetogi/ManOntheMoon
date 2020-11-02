package seed

import "testing"

func TestMockPlayerData(t *testing.T) {

	mockPlayer := MockPlayerData()

	if mockPlayer.PlayerId == "" || mockPlayer.Name == "" {
		t.Errorf("No Player was generated")
	}
}

func TestMockSessionData(t *testing.T) {

	mockSession := MockSessionData()

	if mockSession.PlayerId == "" || mockSession.SessionId == "" {
		t.Errorf("No Session was generated")
	}
}

func TestMockRatingData(t *testing.T) {

	mockRating := MockRatingData()

	if mockRating.PlayerId == "" || mockRating.SessionId == "" {
		t.Errorf("No Rating was generated")
	}
}
