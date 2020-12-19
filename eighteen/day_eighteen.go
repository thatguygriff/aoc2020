package eighteen

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type worksheet struct {
	problems []problem
}

type problem struct {
	lhs      int
	lhsP     *problem
	operator string
	rhs      *problem
}

func parse(input string) *problem {
	p := problem{}
	value := ""
	operation := false

	for i := 0; i < len(input); i++ {
		switch string(input[i]) {
		case " ":
			v, _ := strconv.Atoi(value)
			p.lhs = v
		case "*", "+":
			p.operator = string(input[i])
			operation = true
		case "(":
			p.lhsP = parse(input[i+1:])
			chars := 0
			b := 1
			for j := i + 1; j < len(input); j++ {
				chars++
				if string(input[j]) == "(" {
					b++
				} else if string(input[j]) == ")" {
					b--
				}
				if b == 0 {
					break
				}
			}
			i += chars
		case ")":
			v, _ := strconv.Atoi(value)
			p.lhs = v
			return &p
		default:
			value += string(input[i])
		}

		if operation {
			p.rhs = parse(input[i+2:])
			return &p
		}
	}
	v, _ := strconv.Atoi(value)
	p.lhs = v

	return &p
}

func (p *problem) eval() int {
	result := p.lhs
	if p.lhsP != nil {
		result = p.lhsP.eval()
	}

	op := p.operator
	next := p.rhs
	for next != nil {
		v := next.lhs
		if next.lhsP != nil {
			v = next.lhsP.eval()
		}
		switch op {
		case "+":
			result += v
		case "*":
			result *= v
		}

		op = next.operator
		next = next.rhs
	}

	return result
}

func (p *problem) advEval() int {
	sums := []int{}
	result := p.lhs
	if p.lhsP != nil {
		result = p.lhsP.advEval()
	}

	op := p.operator
	next := p.rhs
	for next != nil {
		v := next.lhs
		if next.lhsP != nil {
			v = next.lhsP.advEval()
		}
		switch op {
		case "+":
			result += v
		case "*":
			sums = append(sums, result)
			result = v
		}

		if next.rhs == nil {
			sums = append(sums, result)
		}

		op = next.operator
		next = next.rhs
	}

	product := 1
	for _, sum := range sums {
		product *= sum
	}

	return product
}

func (w *worksheet) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		w.problems = append(w.problems, *parse(scanner.Text()))
	}

	return nil
}

func (w *worksheet) sum() int {
	sum := 0
	for _, p := range w.problems {
		sum += p.eval()
	}

	return sum
}

func (w *worksheet) advSum() int {
	sum := 0
	for _, p := range w.problems {
		sum += p.advEval()
	}

	return sum
}

// PartOne What is the sum of the homework
func PartOne() string {
	w := worksheet{}
	w.load("eighteen/input.txt")

	return fmt.Sprintf("The sum of the homework is %d", w.sum())
}

// PartTwo What is the advanced sum of the homework
func PartTwo() string {
	w := worksheet{}
	w.load("eighteen/input.txt")

	return fmt.Sprintf("The advanced sum of the homework is %d", w.advSum())
}
