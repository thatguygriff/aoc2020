package one

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type expenseReport struct {
	expenses []int
}

func (e *expenseReport) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		e.expenses = append(e.expenses, value)
	}

	return nil
}

func (e *expenseReport) searchAndMultiply(target int) (int, error) {
	for i := 0; i < len(e.expenses); i++ {
		if e.expenses[i] > target {
			continue
		}

		for j := i + 1; j < len(e.expenses); j++ {
			if e.expenses[j] > target {
				continue
			}

			if (e.expenses[i] + e.expenses[j]) == target {
				return e.expenses[i] * e.expenses[j], nil
			}
		}
	}

	return -1, fmt.Errorf("Unable to find target pair that sums to %d", target)
}

// PartOne Find the multiplied value of the two expenses totalling 2020
func PartOne() string {
	report := expenseReport{}
	err := report.load("one/expenses.txt")
	if err != nil {
		return fmt.Sprintf("Unable to load expenses: %s", err.Error())
	}

	result, err := report.searchAndMultiply(2020)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("The output of the expenses summing to 2020 is %d", result)
}
