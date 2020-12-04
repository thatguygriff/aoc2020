package four

import "testing"

func Test_passports_load(t *testing.T) {
	customs := customs{}
	err := customs.load("sample.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if len(customs.passports) != 4 {
		t.Logf("Expected 4 passports, Got %d passports", len(customs.passports))
		t.FailNow()
	}
}

func Test_passports_check(t *testing.T) {
	customs := customs{}
	customs.load("sample.txt")
	valid := customs.check(false, "cid")
	if valid != 2 {
		t.Logf("Expected 2 valid passports, Got %d valid passwords", customs.check(false, "cid"))
		t.FailNow()
	}
}

func Test_passports_check_strictly_good(t *testing.T) {
	customs := customs{}
	customs.load("sample_strict_good.txt")

	if len(customs.passports) != 4 {
		t.Logf("Expected 4 passports, Got %d passports", len(customs.passports))
		t.FailNow()
	}

	valid := customs.check(true, "cid")
	if valid != 4 {
		t.Logf("Expected 4 valid passports, Got %d valid passwords", valid)
		t.FailNow()
	}
}

func Test_passports_check_strictly_bad(t *testing.T) {
	customs := customs{}
	customs.load("sample_strict_bad.txt")

	if len(customs.passports) != 4 {
		t.Logf("Expected 4 passports, Got %d passports", len(customs.passports))
		t.FailNow()
	}

	valid := customs.check(true, "cid")
	if valid != 0 {
		t.Logf("Expected 0 valid passports, Got %d valid passwords", valid)
		t.FailNow()
	}
}
