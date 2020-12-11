package eleven

import "testing"

func Test_seat_load(t *testing.T) {
	w := waitingRoom("sample.txt")

	if w.width != 10 {
		t.Logf("Expected row width of 10, got %d", w.width)
		t.FailNow()
	}

	if w.rows != 10 {
		t.Logf("Expected 10 rows, got %d", w.rows)
		t.FailNow()
	}
}

func Test_seating_stabilization(t *testing.T) {
	w := waitingRoom("sample.txt")

	occupied := w.simulate(false)

	if occupied != 37 {
		t.Logf("Expected 37 occupied seats, found %d", occupied)
		t.FailNow()
	}
}

func Test_visible_seating_stabilization(t *testing.T) {
	w := waitingRoom("sample.txt")

	occupied := w.simulate(true)

	if occupied != 26 {
		t.Logf("Expected 26 occupied seats, found %d", occupied)
		t.FailNow()
	}
}
