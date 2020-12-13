package thirteen

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bus struct {
	id int
}

type note struct {
	departure int
	routes    []bus
	sequence  []int
}

func notes(f string) (*note, error) {
	n := &note{}
	if err := n.load(f); err != nil {
		return nil, err
	}

	return n, nil
}

func (n *note) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	first := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if first {
			time, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return err
			}
			n.departure = time
			first = false
			continue
		}

		buses := strings.Split(scanner.Text(), ",")
		for _, b := range buses {
			if b == "x" {
				n.sequence = append(n.sequence, -1)
				continue
			}

			id, err := strconv.Atoi(b)
			if err != nil {
				return err
			}
			n.routes = append(n.routes, bus{id: id})
			n.sequence = append(n.sequence, id)
		}
	}

	return nil
}

func (n *note) firstBus() int {
	wait := 0
	now := n.departure

	for {
		for _, bus := range n.routes {
			if now%bus.id == 0 {
				fmt.Println("Found", bus.id, "after a wait of", wait)
				return wait * bus.id
			}
		}

		now++
		wait++
	}
}

func (n *note) sequenceStart() int {
	if len(n.sequence) == 0 {
		return -1
	}

	time := 0
	for {
		// Naive interval skipping, only check times where the first bus is scheduled
		interval := n.sequence[0]
		found := true
		for i := 1; i < len(n.sequence); i++ {
			// Skip unconstrained rules
			if n.sequence[i] == -1 {
				continue
			}

			if (time+i)%n.sequence[i] == 0 {
				// We know that the bus at i is correct, so we can skip interval*bus id ahead for the next time this is valid
				interval *= n.sequence[i]
				continue
			}

			// We reach this point if the sequence is broken. We now advance
			// to the next time the first bus is valid
			time += interval
			found = false
			break
		}
		if found {
			return time
		}

	}
}

// PartOne find the product of the wait * bus id to get to the airport
func PartOne() string {
	n, err := notes("thirteen/input.txt")
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("Found a product of %d", n.firstBus())
}

// PartTwo What is the answer to the consequtive departure minutes contest
func PartTwo() string {
	n, err := notes("thirteen/input.txt")
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("The earliest minute is %d", n.sequenceStart())
}
