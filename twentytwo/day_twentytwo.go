package twentytwo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type player struct {
	deck []int
}

func (p *player) score() int {
	score := 0
	for i := 0; i < len(p.deck); i++ {
		score += p.deck[i] * (len(p.deck) - i)
	}

	return score
}

type combat struct {
	one, two player
}

func (c *combat) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	one := false
	two := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "Player 1:" {
			one = true
			continue
		}

		if scanner.Text() == "" {
			one = false
			two = false
			continue
		}

		if scanner.Text() == "Player 2:" {
			two = true
			continue
		}

		card, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}

		if one {
			c.one.deck = append(c.one.deck, card)
		} else if two {
			c.two.deck = append(c.two.deck, card)
		}
	}

	return nil
}

func (c *combat) play() (round int) {
	for {
		round++

		// play cards
		played := []int{c.one.deck[0], c.two.deck[0]}
		c.one.deck = c.one.deck[1:]
		c.two.deck = c.two.deck[1:]

		if played[0] > played[1] {
			c.one.deck = append(c.one.deck, played[0], played[1])
		} else if played[0] < played[1] {
			c.two.deck = append(c.two.deck, played[1], played[0])
		} else {
			panic("TIES ARE LIKELY IN PART 2")
		}

		if len(c.one.deck) == 0 || len(c.two.deck) == 0 {
			break
		}
	}

	return round
}

func (c *combat) recursivePlay() (winner int) {
	previousRounds := map[string]bool{}
	for {
		currentRound := ""
		for _, c := range c.one.deck {
			currentRound += fmt.Sprintf("%d,", c)
		}
		currentRound += "|"
		for _, c := range c.two.deck {
			currentRound += fmt.Sprintf("%d,", c)
		}

		if previousRounds[currentRound] {
			return 1
		} else {
			previousRounds[currentRound] = true
		}

		// Draw cards as normal
		played := []int{c.one.deck[0], c.two.deck[0]}
		c.one.deck = c.one.deck[1:]
		c.two.deck = c.two.deck[1:]

		w := 0
		if len(c.one.deck) >= played[0] && len(c.two.deck) >= played[1] {
			rPOne := player{}
			for i, c := range c.one.deck {
				if i >= played[0] {
					break
				}
				rPOne.deck = append(rPOne.deck, c)
			}

			rPTwo := player{}
			for i, c := range c.two.deck {
				if i >= played[1] {
					break
				}
				rPTwo.deck = append(rPTwo.deck, c)
			}
			rC := combat{
				one: rPOne,
				two: rPTwo,
			}
			w = rC.recursivePlay()
		} else if played[0] > played[1] {
			w = 1
		} else if played[1] > played[0] {
			w = 2
		}

		if w == 1 {
			c.one.deck = append(c.one.deck, played[0], played[1])
		} else if w == 2 {
			c.two.deck = append(c.two.deck, played[1], played[0])
		}

		if len(c.one.deck) == 0 {
			return 2
		} else if len(c.two.deck) == 0 {
			return 1
		}
	}
}

// PartOne What is the winning player's score
func PartOne() string {
	c := combat{}
	if err := c.load("twentytwo/input.txt"); err != nil {
		return err.Error()
	}

	rounds := c.play()
	return fmt.Sprintf("The game lasted %d rounds with a score of %d - %d", rounds, c.one.score(), c.two.score())
}

// PartTwo What is the winning player's score after recursive combat
func PartTwo() string {
	c := combat{}
	if err := c.load("twentytwo/input.txt"); err != nil {
		return err.Error()
	}

	winner := c.recursivePlay()
	return fmt.Sprintf("The winner was player %d with a score of %d - %d", winner, c.one.score(), c.two.score())
}
