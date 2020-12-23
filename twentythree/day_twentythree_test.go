package twentythree

import "testing"

func Test_load_input(t *testing.T) {
	g := game{}
	g.load("32415")

	if len(g.cups) != 5 {
		t.Logf("Expected 5 cups. Got %d", len(g.cups))
		t.FailNow()
	}

	if g.highestCup != 5 {
		t.Logf("Expected highest cup to be 5, found %d", g.highestCup)
		t.FailNow()
	}
}

func Test_10_moves(t *testing.T) {
	g := game{}
	g.load("389125467")

	if len(g.cups) != 9 {
		t.Logf("Expected 9 cups. Got %d", len(g.cups))
		t.FailNow()
	}

	if g.highestCup != 9 {
		t.Logf("Expected highest cup to be 9, found %d", g.highestCup)
		t.FailNow()
	}

	tenMoves := g.play(10)
	if tenMoves != "92658374" {
		t.Logf("Expected \"92658374\" but found %q", tenMoves)
		t.FailNow()
	}
}

func Test_100_moves(t *testing.T) {
	g := game{}
	g.load("389125467")

	if len(g.cups) != 9 {
		t.Logf("Expected 9 cups. Got %d", len(g.cups))
		t.FailNow()
	}

	if g.highestCup != 9 {
		t.Logf("Expected highest cup to be 9, found %d", g.highestCup)
		t.FailNow()
	}

	hundredMoves := g.play(100)
	if hundredMoves != "67384529" {
		t.Logf("Expected \"67384529\" but found %q", hundredMoves)
		t.FailNow()
	}
}