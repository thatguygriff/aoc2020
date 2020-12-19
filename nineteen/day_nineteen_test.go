package nineteen

import "testing"

func Test_load_sample(t *testing.T) {
	s := satellite{}
	if err := s.load("sample1.txt"); err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	if len(s.rules) != 6 {
		t.Logf("Expected 6 rules, found %d", len(s.rules))
		t.Fail()
	}

	if len(s.messages) != 5 {
		t.Logf("Expected 5 messages, found %d", len(s.messages))
		t.Fail()
	}
}

func Test_match_message(t *testing.T) {
	s := satellite{}
	if err := s.load("sample1.txt"); err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	input := "a"
	_, result := s.match(4, input)
	if result != 1 {
		t.Logf("Expected %d match, got %d", len(input), result)
		t.Fail()
	}

	input = "b"
	_, result = s.match(5, input)
	if result != 1 {
		t.Logf("Expected %d match, got %d", len(input), result)
		t.Fail()
	}

	input = "bb"
	_, result = s.match(2, input)
	if result != 2 {
		t.Logf("Expected %d match, got %d", len(input), result)
		t.Fail()
	}

	input = "bb"
	_, result = s.match(4, input)
	if result != 0 {
		t.Logf("Expected %d match, got %d", len(input), result)
		t.Fail()
	}

	input = "ababbb"
	_, result = s.match(0, input)
	if result != 6 {
		t.Logf("Expected %d match, got %d", len(input), result)
		t.Fail()
	}

	input = "abbbab"
	_, result = s.match(0, input)
	if result != 6 {
		t.Logf("Expected %d match, got %d", len(input), result)
		t.Fail()
	}
}

func Test_not_match_message(t *testing.T) {
	s := satellite{}
	if err := s.load("sample1.txt"); err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	input := "aaabbb"
	match, _ := s.match(0, input)
	if match {
		t.Logf("Expected no match, got %v", match)
		t.Fail()
	}

	input = "bababa"
	match, _ = s.match(0, input)
	if match {
		t.Logf("Expected no match for %s, got %v", input, match)
		t.Fail()
	}

	input = "aaaabbb"
	match, length := s.match(0, input)
	if match && length == len(input) {
		t.Logf("Expected no match for %s, got %v", input, match)
		t.Fail()
	}
}

func Test_message_count(t *testing.T) {
	s := satellite{}
	if err := s.load("sample1.txt"); err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	result := s.messageCount(0)
	if result != 2 {
		t.Logf("Expected 2 messages matching rule 0, got %d", result)
		t.Fail()
	}
}

func Test_message_count2(t *testing.T) {
	s := satellite{}
	if err := s.load("sample2.txt"); err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	result := s.messageCount(0)
	if result != 3 {
		t.Logf("Expected 3 messages matching rule 0, got %d", result)
		t.Fail()
	}
}

func Test_message_count3(t *testing.T) {
	s := satellite{}
	if err := s.load("sample3.txt"); err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	result := s.messageCount(0)
	if result != 12 {
		t.Logf("Expected 12 messages matching rule 0, got %d", result)
		t.Fail()
	}
}

func Test_message_matchloops(t *testing.T) {
	s := satellite{}
	if err := s.load("sample3.txt"); err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	input := "babbbbaabbbbbabbbbbbaabaaabaaa"
	match, length := s.match(0, input)
	if !match || length != len(input) {
		t.Logf("Expected true and %d, got %v and %d", len(input), match, length)
		t.Fail()
	}
}

func Test_message_matchloops2(t *testing.T) {
	s := satellite{}
	if err := s.load("sample3.txt"); err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	input := "aaaabbaaaabbaaa"
	match, length := s.match(0, input)
	if match {
		t.Logf("Expected false and %d, got %v and %d", len(input), match, length)
		t.Fail()
	}
}
