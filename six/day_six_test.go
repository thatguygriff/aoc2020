package six

import "testing"

func Test_data_load(t *testing.T) {
	f := flight{}
	err := f.load("sample.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if len(f.groups) != 6 {
		t.Logf("Expected 6 groups, Got %d groups", len(f.groups))
		t.FailNow()
	}

	uQ := f.groups[0].uniqueQuestion()
	if uQ != 6 {
		t.Logf("Expected 6 questions in Group 1, Got %d questions", uQ)
		t.FailNow()
	}

	uQ = f.groups[1].uniqueQuestion()
	if uQ != 3 {
		t.Logf("Expected 3 questions in Group 2, Got %d questions", uQ)
		t.FailNow()
	}

	uQ = f.groups[2].uniqueQuestion()
	if uQ != 3 {
		t.Logf("Expected 3 questions in Group 3, Got %d questions", uQ)
		t.FailNow()
	}

	uQ = f.groups[3].uniqueQuestion()
	if uQ != 3 {
		t.Logf("Expected 3 questions in Group 4, Got %d questions", uQ)
		t.FailNow()
	}

	uQ = f.groups[4].uniqueQuestion()
	if uQ != 1 {
		t.Logf("Expected 1 questions in Group 5, Got %d questions", uQ)
		t.FailNow()
	}

	uQ = f.groups[5].uniqueQuestion()
	if uQ != 1 {
		t.Logf("Expected 1 questions in Group 6, Got %d groups", uQ)
		t.FailNow()
	}
}

func Test_flight_consensus_questions(t *testing.T) {
	f := flight{}
	err := f.load("sample.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if len(f.groups) != 6 {
		t.Logf("Expected 6 groups, Got %d groups", len(f.groups))
		t.FailNow()
	}

	cQ := f.groups[1].consensusQuestions()
	if cQ != 3 {
		t.Logf("Expected 3 questions in Group 2, Got %d questions", cQ)
		t.FailNow()
	}

	cQ = f.groups[2].consensusQuestions()
	if cQ != 0 {
		t.Logf("Expected 0 questions in Group 3, Got %d questions", cQ)
		t.FailNow()
	}

	cQ = f.groups[3].consensusQuestions()
	if cQ != 1 {
		t.Logf("Expected 3 questions in Group 4, Got %d questions", cQ)
		t.FailNow()
	}

	cQ = f.groups[4].consensusQuestions()
	if cQ != 1 {
		t.Logf("Expected 1 questions in Group 5, Got %d questions", cQ)
		t.FailNow()
	}

	cQ = f.groups[5].consensusQuestions()
	if cQ != 1 {
		t.Logf("Expected 1 questions in Group 6, Got %d groups", cQ)
		t.FailNow()
	}
}
