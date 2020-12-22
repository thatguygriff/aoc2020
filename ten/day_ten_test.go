package ten

import "testing"

func Test_program_load(t *testing.T) {
	b := bag{}
	err := b.load("sample1.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if len(b.adapters) != 11 {
		t.Logf("Expected 11 adapters, Got %d adapters", len(b.adapters))
		t.FailNow()
	}
}

func Test_distribution_sample1(t *testing.T) {
	b := bag{}
	b.load("sample1.txt")

	one, two, three, err := b.distribution()
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}

	if one != 7 || two != 0 || three != 5 {
		t.Logf("Expected 7, 0, 5 jolt diffs, Got %d, %d, %d jolts", one, two, three)
		t.FailNow()
	}
}

func Test_distribution_sample2(t *testing.T) {
	b := bag{}
	b.load("sample2.txt")

	one, two, three, err := b.distribution()
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}

	if one != 22 || two != 0 || three != 10 {
		t.Logf("Expected 22, 0, 10 jolt diffs, Got %d, %d, %d jolts", one, two, three)
		t.FailNow()
	}
}

func Test_arrangement_sample1(t *testing.T) {
	b := newBag("sample1.txt")

	count := b.countArrangements()

	if count != 8 {
		t.Logf("Expected 8 arrangments, got %d", count)
		t.FailNow()
	}
}

func Test_arrangement_sample2(t *testing.T) {
	b := newBag("sample2.txt")

	count := b.countArrangements()

	if count != 19208 {
		t.Logf("Expected 19208 arrangments, got %d", count)
		t.FailNow()
	}
}
