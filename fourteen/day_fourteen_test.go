package fourteen

import "testing"

func Test_computer_load(t *testing.T) {
	c := newComputer("sample.txt")
	if c == nil {
		t.Logf("Unable to load input")
		t.FailNow()
	}

	if len(c.input) != 4 {
		t.Logf("Expected 4 instructions, got %d", len(c.input))
		t.Fail()
	}
}

func Test_computer_exec(t *testing.T) {
	c := newComputer("sample.txt")

	sum := c.sum(1)
	if sum != 165 {
		t.Logf("Expected memory sum of 165, got %d", sum)
		t.Fail()
	}
}

func Test_computer_mask(t *testing.T) {
	c := newComputer("sample2.txt")

	c.execute()
	if c.ones != 20691600039 || c.zeros != 8833416264 {
		t.Logf("Expected mask parse of 20691600039 ons and 8833416264 zeros, got %d and %d", c.ones, c.zeros)
		t.Fail()
	}
}

func Test_computer_version2(t *testing.T) {
	c := newComputer("sample3.txt")

	sum := c.sum(2)
	if sum != 208 {
		t.Logf("Expected memory sum of 208, got %d", sum)
		t.Fail()
	}
}
