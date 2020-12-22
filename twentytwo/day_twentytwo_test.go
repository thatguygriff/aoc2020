package twentytwo

import "testing"

func Test_load_combat(t *testing.T) {
	c := combat{}
	if err := c.load("sample.txt"); err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}

	if len(c.one.deck) != 5 || len(c.two.deck) != 5 {
		t.Logf("Expected decks of 5 and 5, got %d and %d", len(c.one.deck), len(c.two.deck))
	}
}

func Test_combat_play(t *testing.T) {
	c := combat{}
	if err := c.load("sample.txt"); err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}

	c.play()

	if c.one.score() != 0 || c.two.score() != 306 {
		t.Logf("Expected score of 0-306, got %d-%d", c.one.score(), c.two.score())
		t.Fail()
	}
}
