package nine

import "testing"

func Test_program_load(t *testing.T) {
	x := xmas{}
	err := x.load("sample.txt", 5)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if len(x.values) != 20 {
		t.Logf("Expected 20 values, Got %d values", len(x.values))
		t.FailNow()
	}
}

func Test_weakness_detection(t *testing.T) {
	x, err := newXmas("sample.txt", 5)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	v, err := x.detectVulnerability()
	if err != nil || v != 127 {
		t.Logf("Expected vulnerability of 127, Got vulnerability %d", v)
		t.FailNow()
	}
}

func Test_weakness_compute(t *testing.T) {
	x, err := newXmas("sample.txt", 5)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	e, err := x.computeWeakness()
	if err != nil || e != 62 {
		t.Logf("Expected weaknes 62, Got weakness %d", e)
		t.FailNow()
	}
}
