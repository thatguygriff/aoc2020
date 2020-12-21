package twentyone

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type food struct {
	ingredients map[string]bool
	allergens   []string
}

type list struct {
	foods                       []food
	allergens                   map[string]bool
	possibleIngredientAllergens map[string][]string
}

func (l *list) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	l.allergens = map[string]bool{}
	l.possibleIngredientAllergens = map[string][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f := food{
			ingredients: map[string]bool{},
			allergens:   []string{},
		}
		foodParts := strings.Split(scanner.Text(), " (contains ")
		if len(foodParts) != 2 {
			return fmt.Errorf("Unable to parse %s", scanner.Text())
		}
		ingredients := strings.Split(foodParts[0], " ")
		for _, i := range ingredients {
			f.ingredients[i] = true
			l.possibleIngredientAllergens[i] = []string{}
		}

		allergens := strings.Split(foodParts[1], ")")
		f.allergens = strings.Split(allergens[0], ", ")
		for _, a := range f.allergens {
			l.allergens[a] = true
		}

		l.foods = append(l.foods, f)
	}

	return nil
}

func (l *list) isolateIngredientAllergens() {
	// For each allergen iterate through the foods to find the possible ingredients
	for allergen := range l.allergens {
		possibleFoods := []food{}
		possibleIngredients := map[string]bool{}

		// Check food for allergen
		for _, f := range l.foods {
			for _, a := range f.allergens {
				if allergen == a {
					possibleFoods = append(possibleFoods, f)
					for i := range f.ingredients {
						possibleIngredients[i] = true
					}
				}
			}
		}

		// Go through foods and eliminate ingredients that aren't present
		for _, p := range possibleFoods {
			for pi := range possibleIngredients {
				if !p.ingredients[pi] {
					possibleIngredients[pi] = false
				}
			}
		}

		for i, cause := range possibleIngredients {
			if cause {
				l.possibleIngredientAllergens[i] = append(l.possibleIngredientAllergens[i], allergen)
			}
		}
	}
}

func (l *list) countAllergenFreeAppearances() (int, []string) {
	count := 0

	// Find allergen free ingredients
	ingredients := []string{}
	for ingredient, allergens := range l.possibleIngredientAllergens {
		if len(allergens) == 0 {
			ingredients = append(ingredients, ingredient)
		}
	}

	// Count the appearances of the allergen free ingredients in all foods
	for _, i := range ingredients {
		for _, f := range l.foods {
			for ingredient := range f.ingredients {
				if i == ingredient {
					count++
				}
			}
		}
	}

	return count, ingredients
}

// PartOne How many times do allergen free ingredients appear
func PartOne() string {
	l := list{}
	l.load("twentyone/input.txt")

	l.isolateIngredientAllergens()
	count, ingredients := l.countAllergenFreeAppearances()
	return fmt.Sprintf("There are %d appearances of %d allergen free ingredients", count, len(ingredients))
}
