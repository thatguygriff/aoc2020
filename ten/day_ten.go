package ten

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func newBag(f string) *bag {
	b := &bag{}
	b.load(f)
	return b
}

type adapter struct {
	output int
}

type bag struct {
	adapters []adapter
}

func (b *bag) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		b.adapters = append(b.adapters, adapter{
			output: v,
		})
	}

	return nil
}

func (b *bag) distribution() (one, two, three int, err error) {
	if len(b.adapters) == 0 {
		return one, two, three, nil
	}
	sort.SliceStable(b.adapters, func(i, j int) bool { return b.adapters[i].output < b.adapters[j].output })

	current := 0
	for i := 0; i < len(b.adapters); i++ {
		next := b.adapters[i].output
		switch next - current {
		case 1:
			one++
		case 2:
			two++
		case 3:
			three++
		default:
			return one, two, three, fmt.Errorf("Unsupported joltage difference %d", b.adapters[i+1].output-b.adapters[i].output)
		}
		current = next
	}
	// The device is always 3 higher
	three++

	return one, two, three, nil
}

func (b *bag) countArrangements() int {
	if len(b.adapters) == 0 {
		return 0
	}
	sort.SliceStable(b.adapters, func(i, j int) bool { return b.adapters[i].output < b.adapters[j].output })

	// Pilfered from reddit after beating my head against a wall not getting the looping correct
	outputs := []int{0}
	for _, a := range b.adapters {
		outputs = append(outputs, a.output)
	}
	outputs = append(outputs, outputs[len(outputs)-1]+3)

	dp := make([]int, len(outputs))

	dp[0] = 1
	dp[1] = 1

	for i := 2; i < len(dp); i++ {
		for j := i - 1; j >= 0; j-- {
			if outputs[i]-outputs[j] <= 3 {
				dp[i] += dp[j]
			} else {
				break
			}
		}
	}

	return dp[len(dp)-1]
}

// PartOne Find the jolt difference product
func PartOne() string {
	b := newBag("ten/input.txt")
	one, two, three, err := b.distribution()
	if err != nil {
		return fmt.Sprintf("Unable to find a distribution: %v", err)
	}

	return fmt.Sprintf("Got %d, %d, %d jolts differences with a %d product", one, two, three, one*three)
}

// PartTwo Find the number of arrangements of adapters
func PartTwo() string {
	b := newBag("ten/input.txt")
	return fmt.Sprintf("Found %d possible arrangements", b.countArrangements())
}
