package two

import "testing"

func Test_db_load(t *testing.T) {
	database := db{}
	if err := database.load("sample.txt"); err != nil {
		t.Log(err)
		t.FailNow()
	}

	if len(database.passwords) != 3 {
		t.Logf("Expected 3 passwords, Got %d passwords", len(database.passwords))
		t.FailNow()
	}
}

func Test_db_valiate(t *testing.T) {
	database := db{}
	database.load("sample.txt")

	validCount := database.validate()

	if validCount != 2 {
		t.Logf("Expected 2 valid passwords, Got %d valid passwords", validCount)
		t.FailNow()
	}
}
