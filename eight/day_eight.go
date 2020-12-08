package eight

import (
	"bufio"
	"fmt"
	"os"
)

type program struct {
	instructions []instruction
	accumulator  int
	preempted    bool
}

type instruction struct {
	operation string
	argument  int
	executed  bool
}

func (p *program) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var op string
		var value int

		count, err := fmt.Sscanf(scanner.Text(), "%s %d", &op, &value)
		if count != 2 || err != nil {
			return fmt.Errorf("Unable to parse %s: %w", scanner.Text(), err)
		}

		p.instructions = append(p.instructions, instruction{
			operation: op,
			argument:  value,
			executed:  false,
		})
	}

	return nil
}

func (p *program) exec() int {
	i := 0
	for {
		if i >= len(p.instructions) {
			p.preempted = false
			return p.accumulator
		} else if i < 0 {
			return p.accumulator
		}

		if p.instructions[i].executed {
			break
		}

		p.instructions[i].executed = true
		switch p.instructions[i].operation {
		case "acc":
			p.accumulator += p.instructions[i].argument
			i++
		case "jmp":
			i += p.instructions[i].argument
		default:
			i++
		}
	}

	return p.accumulator
}

func (p *program) reset() {
	p.accumulator = 0
	p.preempted = true
	for i := 0; i < len(p.instructions); i++ {
		p.instructions[i].executed = false
	}
}

func (p *program) flipJmpNop(index int) {
	if index < 0 || index > len(p.instructions) {
		return
	}

	switch p.instructions[index].operation {
	case "nop":
		p.instructions[index].operation = "jmp"
	case "jmp":
		p.instructions[index].operation = "nop"
	}
}

func (p *program) heal() int {
	for i := 0; i < len(p.instructions); i++ {
		p.reset()
		p.flipJmpNop(i)
		p.exec()

		if !p.preempted {
			return i
		}
		p.flipJmpNop(i)
	}

	return -1
}

// PartOne Find the value of the accumulator before any instruction is executed twice
func PartOne() string {
	p := program{}
	p.load("eight/handheld.code")

	return fmt.Sprintf("The accumulator is at %d before the loop begins", p.exec())
}

// PartTwo Find the corrupt instruction, flip it, and execute to termination
func PartTwo() string {
	p := program{}
	p.load("eight/handheld.code")
	p.exec()
	index := p.heal()

	return fmt.Sprintf("The accumulator is at %d after termination by correcting instruction %d", p.accumulator, index+1)
}
