package nineteen

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	matches    string
	composedOf [][]int
}

type satellite struct {
	rules    map[int]rule
	messages []string
}

func (s *satellite) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	s.rules = map[int]rule{}
	scanner := bufio.NewScanner(file)
	messages := false
	for scanner.Scan() {
		if !messages {
			if scanner.Text() == "" {
				messages = true
				continue
			}

			ruleSplit := strings.Split(scanner.Text(), ": ")
			ruleKey, err := strconv.Atoi(ruleSplit[0])
			if err != nil {
				return err
			}

			r := rule{}

			if strings.Contains(ruleSplit[1], "\"a\"") {
				r.matches = "a"
			} else if strings.Contains(ruleSplit[1], "\"b\"") {
				r.matches = "b"
			} else {
				options := strings.Split(ruleSplit[1], " | ")
				for _, o := range options {
					m := []int{}
					rules := strings.Split(o, " ")
					for _, r := range rules {
						t, err := strconv.Atoi(r)
						if err != nil {
							return err
						}
						m = append(m, t)
					}
					r.composedOf = append(r.composedOf, m)
				}
			}

			s.rules[ruleKey] = r
			continue
		}

		s.messages = append(s.messages, scanner.Text())
	}

	return nil
}

func (s *satellite) match(r int, input string) (bool, int) {
	if len(input) == 0 {
		return false, 0
	}

	rule := s.rules[r]

	if rule.matches == string(input[0]) {
		return true, 1
	}

	matchedChars := 0
	for i := 0; i < len(rule.composedOf); i++ {
		matchedChars = 0
		valid := true
		if len(rule.composedOf[i]) > len(input) {
			continue
		}
		for _, sr := range rule.composedOf[i] {
			match, count := s.match(sr, input[matchedChars:])
			matchedChars += count
			if !match {
				valid = false
				break
			}
		}
		if valid {
			return true, matchedChars
		}
	}

	return false, matchedChars
}

func (s *satellite) messageCount(r int) (count int) {
	for _, m := range s.messages {
		match, length := s.match(r, m)
		if match && length == len(m) {
			count++
		}
	}

	return count
}

// PartOne How many messages match message 0
func PartOne() string {
	s := satellite{}
	if err := s.load("nineteen/input.txt"); err != nil {
		return err.Error()
	}

	return fmt.Sprintf("There are %d messages that match rule 0", s.messageCount(0))
}

// PartTwo How many messages match message 0 with alternate rules
func PartTwo() string {
	s := satellite{}
	if err := s.load("nineteen/input2.txt"); err != nil {
		return err.Error()
	}

	return fmt.Sprintf("There are %d messages that match rule 0", s.messageCount(0))
}
