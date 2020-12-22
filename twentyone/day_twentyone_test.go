package twentyone

import "testing"

func Test_loading_list(t *testing.T) {
	l := list{}
	err := l.load("sample.txt")
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}

	foods := len(l.foods)
	if foods != 4 {
		t.Logf("Expected 4 foods, found %d", foods)
		t.FailNow()
	}

	ingredients := len(l.ingredientAllergens)
	if ingredients != 7 {
		t.Logf("Expected 7 ingredients, found %d", ingredients)
		t.FailNow()
	}
}

func Test_find_allergenfree(t *testing.T) {
	l := list{}
	err := l.load("sample.txt")
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}

	l.isolateIngredientAllergens()
	count, ingredients := l.countAllergenFreeAppearances()
	if count != 5 {
		t.Logf("Expected 5 allergen free ingredients, got %d %v", count, ingredients)
		t.FailNow()
	}
}

func Test_canonical_list(t *testing.T) {
	l := list{}
	err := l.load("sample.txt")
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}

	l.isolateIngredientAllergens()
	list := l.canonicalDangerList()
	if list != "mxmxvkd,sqjhc,fvjkl" {
		t.Logf("Expected list of mxmxvkd,sqjhc,fvjkl, got %s", list)
		t.Fail()
	}
}

func Test_fullcanonical_list(t *testing.T) {
	l := list{}
	err := l.load("input.txt")
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}

	l.isolateIngredientAllergens()
	list := l.canonicalDangerList()
	if list != "vmhqr,qxfzc,khpdjv,gnrpml,xrmxxvn,rfmvh,rdfr,jxh" {
		t.Logf("Expected list \"vmhqr,qxfzc,khpdjv,gnrpml,xrmxxvn,rfmvh,rdfr,jxh\", got %s", list)
		t.Fail()
	}
}
