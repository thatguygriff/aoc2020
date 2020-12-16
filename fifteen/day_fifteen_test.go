package fifteen

import "testing"

func Test_load_input(t *testing.T) {
	g := game{}
	if err := g.load("sample.txt"); err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}

	if len(g.numbers) != 3 {
		t.Logf("Expected 3 starting numbers, got %d", len(g.numbers))
		t.FailNow()
	}
}

func Test_get_guess1(t *testing.T) {
	g := game{}
	if err := g.load("sample.txt"); err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}

	r := g.valueAt(9)
	if r != 4 {
		t.Logf("Expected 4, but got %d", r)
		t.FailNow()
	}
}

func Test_get_guess2(t *testing.T) {
	g := game{}
	if err := g.load("sample.txt"); err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}

	r := g.valueAt(2020)
	if r != 436 {
		t.Logf("Expected 436, but got %d", r)
		t.FailNow()
	}
}

func Test_get_guess3(t *testing.T) {
	g := game{}
	if err := g.load("sample.txt"); err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}

	r := g.valueAt(30000000)
	if r != 175594 {
		t.Logf("Expected 175594, but got %d", r)
		t.FailNow()
	}
}


