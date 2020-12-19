package eighteen

import "testing"

func Test_math_parse(t *testing.T) {
	w := worksheet{}
	w.load("sample.txt")

	if len(w.problems) != 6 {
		t.Logf("Expected 6 probems, found %d", len(w.problems))
		t.FailNow()
	}
}

func Test_problem_eval1(t *testing.T) {
	w := worksheet{}
	w.load("sample1.txt")

	result := w.problems[0].eval()
	if result != 71 {
		t.Logf("Expected result of 71, got %d", result)
		t.Fail()
	}
}

func Test_problem_eval2(t *testing.T) {
	w := worksheet{}
	w.load("sample2.txt")

	result := w.problems[0].eval()
	if result != 51 {
		t.Logf("Expected result of 51, got %d", result)
		t.Fail()
	}
}

func Test_problem_eval4(t *testing.T) {
	w := worksheet{}
	w.load("sample4.txt")

	result := w.problems[0].eval()
	if result != 437 {
		t.Logf("Expected result of 437, got %d", result)
		t.Fail()
	}
}

func Test_problem_eval6(t *testing.T) {
	w := worksheet{}
	w.load("sample6.txt")

	result := w.problems[0].eval()
	if result != 13632 {
		t.Logf("Expected result of 13632, got %d", result)
		t.Fail()
	}
}

func Test_problem_eval(t *testing.T) {
	w := worksheet{}
	w.load("sample.txt")

	result := w.problems[0].eval()
	if result != 71 {
		t.Logf("Expected result of 71, got %d", result)
		t.Fail()
	}

	result = w.problems[1].eval()
	if result != 51 {
		t.Logf("Expected result of 51, got %d", result)
		t.Fail()
	}

	result = w.problems[2].eval()
	if result != 26 {
		t.Logf("Expected result of 26, got %d", result)
		t.Fail()
	}

	result = w.problems[3].eval()
	if result != 437 {
		t.Logf("Expected result of 437, got %d", result)
		t.Fail()
	}

	result = w.problems[4].eval()
	if result != 12240 {
		t.Logf("Expected result of 12240, got %d", result)
		t.Fail()
	}

	result = w.problems[5].eval()
	if result != 13632 {
		t.Logf("Expected result of 13632, got %d", result)
		t.Fail()
	}
}

func Test_sheet_eval(t *testing.T) {
	w := worksheet{}
	w.load("sample.txt")

	result := w.sum()
	if result != 26457 {
		t.Logf("Expected result of 26457, got %d", result)
		t.FailNow()
	}
}

func Test_problem_advEval1(t *testing.T) {
	w := worksheet{}
	w.load("sample1.txt")

	result := w.problems[0].advEval()
	if result != 231 {
		t.Logf("Expected result of 231, got %d", result)
		t.Fail()
	}
}

func Test_problem_advEval2(t *testing.T) {
	w := worksheet{}
	w.load("sample2.txt")

	result := w.problems[0].advEval()
	if result != 51 {
		t.Logf("Expected result of 51, got %d", result)
		t.Fail()
	}
}

func Test_problem_advEval6(t *testing.T) {
	w := worksheet{}
	w.load("sample6.txt")

	result := w.problems[0].advEval()
	if result != 23340 {
		t.Logf("Expected result of 23340, got %d", result)
		t.Fail()
	}
}

func Test_sheet_advEval(t *testing.T) {
	w := worksheet{}
	w.load("sample.txt")

	result := w.advSum()
	if result != 694173 {
		t.Logf("Expected result of 694173, got %d", result)
		t.FailNow()
	}
}
