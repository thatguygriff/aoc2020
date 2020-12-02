package one

import "testing"

func Test_expenseReport_load(t *testing.T) {
	report := expenseReport{}
	if err := report.load("sample.txt"); err != nil {
		t.Log(err)
		t.FailNow()
	}

	if len(report.expenses) != 6 {
		t.Logf("Expected 6 expenses, Got %d expesnes", len(report.expenses))
		t.FailNow()
	}
}

func Test_expenseReport_findPairAndMultiply(t *testing.T) {
	report := expenseReport{
		expenses: []int{
			1721,
			979,
			366,
			299,
			675,
			1456,
		},
	}

	result, _ := report.findPairAndMultiply(2020)
	if result != 514579 {
		t.Logf("Expected 514579, Got %d", result)
		t.FailNow()
	}
}

func Test_expenseReport_findTrioAndMultiply(t *testing.T) {
	report := expenseReport{
		expenses: []int{
			1721,
			979,
			366,
			299,
			675,
			1456,
		},
	}

	result, _ := report.findTrioAndMultiply(2020)
	if result != 241861950 {
		t.Logf("Expected 241861950, Got %d", result)
		t.FailNow()
	}
}
