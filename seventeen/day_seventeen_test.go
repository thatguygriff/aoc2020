package seventeen

import "testing"

func Test_dimension_init(t *testing.T) {
	p := pocket{}
	if err := p.initialize("sample.txt"); err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}

	active := p.active()
	if active != 5 {
		t.Logf("Expected 5 active after 0 rounds, found %d", active)
		t.FailNow()
	}
}

func Test_dimension_simulate1(t *testing.T) {
	p := pocket{}
	if err := p.initialize("sample.txt"); err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	p.print()

	p = simulate(p, 1)
	p.print()

	active := p.active()
	if active != 12 {
		t.Logf("Expected 11 active after 1 rounds, found %d", active)
		t.FailNow()
	}
}

func Test_dimension_simulate(t *testing.T) {
	p := pocket{}
	if err := p.initialize("sample.txt"); err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}

	p = simulate(p, 6)
	p.print()
	active := p.active()
	if active != 112 {
		t.Logf("Expected 112 active after 6 rounds, found %d", active)
		t.FailNow()
	}
}
