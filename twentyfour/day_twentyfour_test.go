package twentyfour

import "testing"

func Test_load_instructions(t *testing.T) {
	l := lobby{}
	l.load("sample.txt")

	if len(l.instructions) != 20 {
		t.Logf("Expected 20 instructions, found %d", len(l.instructions))
		t.Fail()
	}
}

func Test_layout_tiles(t *testing.T) {
	l := lobby{}
	l.load("sample.txt")

	if len(l.instructions) != 20 {
		t.Logf("Expected 20 instructions, found %d", len(l.instructions))
		t.Fail()
	}

	l.layout()
	black, white := l.countTiles()
	if black+white != 15 {
		t.Logf("Expected 15 tiles, got %d", black+white)
		t.Fail()
	}

	if black != 10 {
		t.Logf("Expected 10 black tiles, got %d", black)
		t.Fail()
	}

	if white != 5 {
		t.Logf("Expected 5 white tiles, got %d", white)
		t.Fail()
	}
}

func Test_daily_flip(t *testing.T) {
	l := lobby{}
	l.load("sample.txt")

	l.layout()
	black, white := l.countTiles()
	if black+white != 15 {
		t.Logf("Expected 15 tiles, got %d", black+white)
		t.Fail()
	}

	if black != 10 {
		t.Logf("Expected 10 black tiles, got %d", black)
		t.Fail()
	}

	l.dailyFlip(1)
	black, _ = l.countTiles()
	if black != 15 {
		t.Logf("Expected 15 black tiles, got %d", black)
		t.FailNow()
	}

	l.dailyFlip(1)
	black, _ = l.countTiles()
	if black != 12 {
		t.Logf("Expected 12 black tiles, got %d", black)
		t.FailNow()
	}

	l.dailyFlip(1)
	black, _ = l.countTiles()
	if black != 25 {
		t.Logf("Expected 25 black tiles, got %d", black)
		t.FailNow()
	}

	l.dailyFlip(1)
	black, _ = l.countTiles()
	if black != 14 {
		t.Logf("Expected 14 black tiles, got %d", black)
		t.FailNow()
	}
}

func Test_flip(t *testing.T) {
	l := lobby{}
	l.load("sample.txt")

	l.layout()
	black, white := l.countTiles()
	if black+white != 15 {
		t.Logf("Expected 15 tiles, got %d", black+white)
		t.Fail()
	}

	if black != 10 {
		t.Logf("Expected 10 black tiles, got %d", black)
		t.Fail()
	}

	l.dailyFlip(100)
	black, _ = l.countTiles()
	if black != 2208 {
		t.Logf("Expected 2208 black tiles, got %d", black)
		t.FailNow()
	}
}
