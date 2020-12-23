package twentythree

import "testing"

func Test_load_input(t *testing.T) {
	g := game{}
	g.load("32415")

	if g.cups != 5 {
		t.Logf("Expected 5 cups. Got %d", g.cups)
		t.FailNow()
	}
}

func Test_10_moves(t *testing.T) {
	g := game{}
	g.load("389125467")

	if g.cups != 9 {
		t.Logf("Expected 9 cups. Got %d", g.cups)
		t.FailNow()
	}

	tenMoves := g.play(10, true)
	if tenMoves != "92658374" {
		t.Logf("Expected \"92658374\" but found %q", tenMoves)
		t.FailNow()
	}
}

func Test_100_moves(t *testing.T) {
	g := game{}
	g.load("389125467")

	if g.cups != 9 {
		t.Logf("Expected 9 cups. Got %d", g.cups)
		t.FailNow()
	}

	hundredMoves := g.play(100, true)
	if hundredMoves != "67384529" {
		t.Logf("Expected \"67384529\" but found %q", hundredMoves)
		t.FailNow()
	}
}

func Test_10million_moves(t *testing.T) {
	g := game{}
	g.milliload("389125467")

	if g.cups != 1000000 {
		t.Logf("Expected 1000000 cups. Got %d", g.cups)
		t.FailNow()
	}

	g.play(10000000, false)
	starProduct := g.starProduct()
	if starProduct != 149245887792 {
		t.Logf("Expected 149245887792 but got produc %q", 149245887792)
		t.FailNow()
	}
}
