package twelve

import "testing"

func Test_load_boat(t *testing.T) {
	b, err := loadBoat("sample.txt")
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}

	if len(b.orders) != 5 {
		t.Logf("Expected 5 orders, got %d", len(b.orders))
		t.FailNow()
	}
}

func Test_boat_navigate(t *testing.T) {
	b, _ := loadBoat("sample.txt")

	result := b.navigate()
	if result != 25 {
		t.Logf("Expected a navigation result of 25, got %d", result)
		t.FailNow()
	}
}

func Test_waypoint_navigate(t *testing.T) {
	b, _ := loadBoat("sample.txt")

	result := b.waypointNavigate()
	if result != 286 {
		t.Logf("Expected a navigation result of 286, got %d", result)
		t.FailNow()
	}
}
