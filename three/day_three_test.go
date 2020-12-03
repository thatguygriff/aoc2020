package three

import "testing"

func Test_forest_load(t *testing.T) {
	forest := forest{}
	if err := forest.load("sample.txt"); err != nil {
		t.Log(err)
		t.FailNow()
	}

	if forest.rows != 11 {
		t.Logf("Expected 11 rows, Got %d rows", forest.rows)
		t.FailNow()
	}

	if forest.width != 11 {
		t.Logf("Expected 11 columns, Got %d columns", forest.rows)
		t.FailNow()
	}
}

func Test_forest_toboggan(t *testing.T) {
	forest := forest{}
	forest.load("sample.txt")

	impacts := forest.toboggan(3, 1)

	if impacts != 7 {
		t.Logf("Expected 7 tree hits, actually hit %d trees", impacts)
		t.FailNow()
	}
}
