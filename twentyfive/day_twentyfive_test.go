package twentyfive

import "testing"

func Test_generate_encryption(t *testing.T) {
	card := key{
		public: 5764801,
		sub:    7,
	}

	door := key{
		public: 17807724,
		sub:    7,
	}

	card.determineLoopSize()
	door.determineLoopSize()

	if card.loopsize != 8 {
		t.Logf("expected loop size of 8, got %d", card.loopsize)
		t.Fail()
	}

	if door.loopsize != 11 {
		t.Logf("expected loop size of 11, got %d", door.loopsize)
		t.Fail()
	}

	encryption1 := card.encryption(door.public)
	encryption2 := door.encryption(card.public)

	if encryption1 != 14897079 {
		t.Logf("Expected encryption key of 14897079, got %d", encryption1)
		t.FailNow()
	}

	if encryption1 != encryption2 {
		t.Logf("Expected to generate the same encryption key, got %d and %d", encryption1, encryption2)
		t.Fail()
	}
}

func Test_part_one(t *testing.T) {
	card := key{
		public: 2069194,
		sub:    7,
	}

	door := key{
		public: 16426071,
		sub:    7,
	}

	card.determineLoopSize()
	door.determineLoopSize()

	encryption1 := card.encryption(door.public)

	if encryption1 != 11576351 {
		t.Logf("Expected encryption key 11576351, got %d", encryption1)
		t.FailNow()
	}
}
