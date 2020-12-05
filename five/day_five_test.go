package five

import "testing"

func Test_data_load(t *testing.T) {
	plane := plane{}
	err := plane.load("sample.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if len(plane.passes) != 4 {
		t.Logf("Expected 4 data, Got %d data", len(plane.passes))
		t.FailNow()
	}
}

func Test_pass_decode(t *testing.T) {
	plane := plane{}
	plane.load("sample.txt")

	var row, seat, id int
	row, seat, id = decode(plane.passes[0].raw)
	if row != 44 || seat != 5 || id != 357 {
		t.Logf("Expected row 44, seat 5, id 357: Got row %d, seat %d, id %d", row, seat, id)
		t.FailNow()
	}

	row, seat, id = decode(plane.passes[1].raw)
	if row != 70 || seat != 7 || id != 567 {
		t.Logf("Expected row 70, seat 7, id 567: Got row %d, seat %d, id %d", row, seat, id)
		t.FailNow()
	}

	row, seat, id = decode(plane.passes[2].raw)
	if row != 14 || seat != 7 || id != 119 {
		t.Logf("Expected row 14, seat 7, id 119: Got row %d, seat %d, id %d", row, seat, id)
		t.FailNow()
	}

	row, seat, id = decode(plane.passes[3].raw)
	if row != 102 || seat != 4 || id != 820 {
		t.Logf("Expected row 102, seat 4, id 820: Got row %d, seat %d, id %d", row, seat, id)
		t.FailNow()
	}
}
