package seven

import "testing"

func Test_data_load(t *testing.T) {
	rules := rules{}
	err := rules.load("sample.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if len(rules.bags) != 9 {
		t.Logf("Expected 9 rules, Got %d rules", len(rules.bags))
		t.FailNow()
	}

	ruleCapacity := []int{2, 2, 1, 2, 2, 2, 2, 0, 0}
	bagCapacity := []int{3, 7, 1, 11, 3, 7, 11, 0, 0}

	for i, bag := range rules.bags {
		if len(bag.capacity) != ruleCapacity[i] {
			t.Logf("Expected %d bag, Got %d bags", ruleCapacity[i], len(bag.capacity))
			t.FailNow()
		}

		capacity := 0
		for _, count := range bag.capacity {
			capacity += count
		}

		if capacity != bagCapacity[i] {
			t.Logf("Expected %d bags, Got %d bags", bagCapacity[i], capacity)
			t.FailNow()
		}
	}
}

func Test_bag_counting(t *testing.T) {
	rules := rules{}
	rules.load("sample.txt")

	count, bags := rules.canContain("shiny gold")
	if count != 4 {
		t.Logf("Expected 4 bags could hold shiny gold, Got %d bags, %s", count, bags)
		t.FailNow()
	}
}

func Test_bag_capacity(t *testing.T) {
	rules := rules{}
	rules.load("sample.txt")

	count := rules.canHold("shiny gold")
	if count != 32 {
		t.Logf("Expected 32 bag capacity for shiny gold, Got %d bag capacity", count)
		t.FailNow()
	}
}
