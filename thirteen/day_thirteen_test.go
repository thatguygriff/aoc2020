package thirteen

import "testing"

func Test_note_load(t *testing.T) {
	n, err := notes("sample.txt")
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	if len(n.routes) != 5 {
		t.Logf("Expected 5 routes, found %d", n.routes)
		t.Fail()
	}

	if n.departure != 939 {
		t.Logf("Expected departure of 939, found %d", n.departure)
		t.Fail()
	}
}

func Test_find_first(t *testing.T) {
	n, _ := notes("sample.txt")

	result := n.firstBus()

	if result != 295 {
		t.Logf("Expected product of 295 for first bus, got %d", result)
		t.Fail()
	}
}

func Test_contest_start(t *testing.T) {
	n, _ := notes("sample.txt")

	result := n.sequenceStart()
	if result != 1068781 {
		t.Logf("Expected time of 1068781 for sequence, got %d", result)
		t.Fail()
	}

	n.sequence = []int{17, -1, 13, 19}
	result = n.sequenceStart()
	if result != 3417 {
		t.Logf("Expected time of 3417 for sequence, got %d", result)
		t.Fail()
	}

	n.sequence = []int{67, 7, 59, 61}
	result = n.sequenceStart()
	if result != 754018 {
		t.Logf("Expected time of 754018 for sequence, got %d", result)
		t.Fail()
	}

	n.sequence = []int{67, -1, 7, 59, 61}
	result = n.sequenceStart()
	if result != 779210 {
		t.Logf("Expected time of 779210 for sequence, got %d", result)
		t.Fail()
	}

	n.sequence = []int{67, 7, -1, 59, 61}
	result = n.sequenceStart()
	if result != 1261476 {
		t.Logf("Expected time of 1261476 for sequence, got %d", result)
		t.Fail()
	}

	n.sequence = []int{1789, 37, 47, 1889}
	result = n.sequenceStart()
	if result != 1202161486 {
		t.Logf("Expected time of 1202161486 for sequence, got %d", result)
		t.Fail()
	}
}
