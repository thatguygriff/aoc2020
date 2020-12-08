package eight

import "testing"

func (p *program) visited() int {
	highest := 0
	for i, instruction := range p.instructions {
		if instruction.executed {
			if i > highest {
				highest = i
			}
		}
	}
	return highest
}

func Test_program_load(t *testing.T) {
	p := program{}
	err := p.load("sample.code")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if len(p.instructions) != 9 {
		t.Logf("Expected 9 instructions, Got %d instructions", len(p.instructions))
		t.FailNow()
	}
}

func Test_program_execute(t *testing.T) {
	p := program{}
	p.load("sample.code")

	result := p.exec()
	if result != 5 {
		t.Logf("Expected result of 5, Got %d", result)
		t.FailNow()
	}
}

func Test_program_terminate(t *testing.T) {
	p := program{}
	p.load("sample.code")

	p.exec()
	highest := p.visited()
	p.flipJmpNop(highest)
	p.reset()
	result := p.exec()
	highest = p.visited()

	if highest+1 != len(p.instructions) {
		t.Logf("Made it to %d of %d instructions", highest+1, len(p.instructions))
		t.Fail()
	}

	if result != 8 || p.preempted {
		t.Logf("Expected result of 8, Got %d and the program was preempted", result)
		t.FailNow()
	}
}

func Test_program_heal(t *testing.T) {
	p := program{}
	p.load("sample.code")

	p.exec()
	index := p.heal()
	highest := p.visited()

	if highest+1 != len(p.instructions) {
		t.Logf("Made it to %d of %d instructions by flipping instruction %d", highest+1, len(p.instructions), index+1)
		t.Fail()
	}

	if p.accumulator != 8 || p.preempted {
		t.Logf("Expected result of 8, Got %d and the program was preempted", p.accumulator)
		t.FailNow()
	}
}

func Test_handheld_terminate(t *testing.T) {
	p := program{}
	p.load("handheld.code")

	p.exec()
	highest := p.visited()
	t.Logf("Made it to %d of %d instructions on the first run", highest+1, len(p.instructions))
	p.heal()
	highest = p.visited()

	if highest+1 != len(p.instructions) || p.preempted {
		t.Logf("Made it to %d of %d instructions", highest+1, len(p.instructions))
		t.Fail()
	}
}
