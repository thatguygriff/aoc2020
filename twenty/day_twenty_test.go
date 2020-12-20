package twenty

import "testing"

func Test_tile_load(t *testing.T) {
	m := message{}
	m.load("sample.txt")

	if len(m.tiles) != 9 {
		t.Logf("Expected 9 tiles to be loaded, got %d", len(m.tiles))
		t.FailNow()
	}

	if len(m.grid) != 3 && len(m.grid[0]) != 3 {
		t.Logf("Expected a 3x3 grid, got %dx%d", len(m.grid), len(m.grid[0]))
		t.FailNow()
	}
}

func Test_arrange_tile(t *testing.T) {
	m := message{}
	m.load("sample.txt")

	result := m.cornerProduct()
	if result != 20899048083289 {
		t.Logf("Expected a result of 20899048083289, got %d", result)
		t.FailNow()
	}
}

func Test_arrange_input(t *testing.T) {
	m := message{}
	m.load("input.txt")

	result := m.cornerProduct()
	if result == -1 {
		t.Logf("Expected a result other than -1, got %d", result)
		t.FailNow()
	}
}

func Test_count_monsters(t *testing.T) {
	m := message{}
	m.load("sample.txt")

	grid := m.layout()
	if !grid {
		t.Log("Expected a grid")
		t.FailNow()
	}

	m.combineTiles()
	m.alignMonsters()

	result := m.seaMonsters()
	if result != 2 {
		t.Logf("Expected 2 monsters, found %d", result)
		t.FailNow()
	}
}

func Test_count_roughness(t *testing.T) {
	m := message{}
	m.load("sample.txt")

	m.layout()
	m.combineTiles()
	m.alignMonsters()

	result := m.countRoughness()
	if result != 273 {
		t.Logf("Expected 273 # not part of a monster, found %d", result)
		t.FailNow()
	}
}
