package sixteen

import "testing"

func Test_ticket_loading(t *testing.T) {
	ts := ticketScanner{}
	if err := ts.load("sample.txt"); err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	if len(ts.rules) != 3 {
		t.Logf("Expected 3 rules, go %d: %v", len(ts.rules), ts.rules)
		t.Fail()
	}

	if ts.me == nil {
		t.Log("Expected my ticket to be loaded")
		t.Fail()
	}

	if len(ts.nearby) != 4 {
		t.Logf("Expected 4 nearby tickets, got %d", len(ts.nearby))
		t.Fail()
	}
}

func Test_ticket_scanerror(t *testing.T) {
	ts := ticketScanner{}
	if err := ts.load("sample.txt"); err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	rate := ts.nearbyErrorRate()
	if rate != 71 {
		t.Logf("Expected error rate of 71, got %d", rate)
		t.Fail()
	}
}

func Test_ticket_fieldDetection(t *testing.T) {
	ts := ticketScanner{}
	if err := ts.load("sample2.txt"); err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	ts.scanValid()
	if len(ts.valid) != 3 {
		t.Logf("Expected 3 valid tickets, got %d", len(ts.valid))
		t.FailNow()
	}

	ts.determineFields()

	if len(ts.fields) != 3 || ts.fields[0] != "row" || ts.fields[1] != "class" || ts.fields[2] != "seat" {
		t.Logf("Expected row, class, seat but got %v", ts.fields)
		t.Fail()
	}
}

func Test_ticket_product(t *testing.T) {
	ts := ticketScanner{}
	if err := ts.load("sample3.txt"); err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	ts.scanValid()
	if len(ts.valid) != 3 {
		t.Logf("Expected 3 valid tickets, got %d", len(ts.valid))
		t.FailNow()
	}

	ts.determineFields()
	p := ts.fieldProduct("departure")
	if p != 132 {
		t.Logf("Expected 132 got %d", p)
		t.Fail()
	}
}
