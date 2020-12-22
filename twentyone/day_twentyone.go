package twentyone

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type food struct {
	ingredients map[string]bool
	allergens   []string
}

type list struct {
	foods               []food
	allergens           map[string]bool
	ingredientAllergens map[string][]string
}

func (l *list) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	l.allergens = map[string]bool{}
	l.ingredientAllergens = map[string][]string{}

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
			l.ingredientAllergens[i] = []string{}
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
				l.ingredientAllergens[i] = append(l.ingredientAllergens[i], allergen)
			}
		}
	}

	tmpIngredients := map[string][]string{}
	// Remove duplicate allergens
	for ingredient, allergens := range l.ingredientAllergens {
		if len(allergens) < 1 {
			continue
		}

		tmpIngredients[ingredient] = allergens
	}

	stable := false
	for !stable {
		causes := l.getIngredientsWithAllergens()
		// Check to see if every ingredient has 1 allergen
		stable = true
		for _, allergens := range causes {
			if len(allergens) != 1 {
				stable = false
			}
		}

		if stable {
			break
		}

		// Find all allergens with with only 1 option. Remove that option from others
		for i, a := range causes {
			if len(a) == 1 {
				l.claimAllergen(i, a[0])
			}
		}
	}
}

func (l *list) getIngredientsWithAllergens() map[string][]string {
	tmpIngredients := map[string][]string{}
	// Remove duplicate allergens
	for ingredient, allergens := range l.ingredientAllergens {
		if len(allergens) < 1 {
			continue
		}

		tmpIngredients[ingredient] = allergens
	}
	return tmpIngredients
}

func (l *list) claimAllergen(ingredient, allergen string) {
	for i, alls := range l.ingredientAllergens {
		if i == ingredient {
			continue
		}

		filtered := []string{}
		for _, a := range alls {
			if a != allergen {
				filtered = append(filtered, a)
			}
		}
		l.ingredientAllergens[i] = filtered
	}
}

func (l *list) countAllergenFreeAppearances() (int, []string) {
	count := 0

	// Find allergen free ingredients
	ingredients := []string{}
	for ingredient, allergens := range l.ingredientAllergens {
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

func (l *list) canonicalDangerList() string {
	dangerous := []string{}

	for _, allergens := range l.ingredientAllergens {
		if len(allergens) == 1 {
			dangerous = append(dangerous, allergens[0])
		}
	}

	sort.Strings(dangerous)
	ingredients := []string{}
	for _, allergen := range dangerous {
		for ingredient, allergens := range l.ingredientAllergens {
			if len(allergens) == 1 {
				if allergens[0] == allergen {
					ingredients = append(ingredients, ingredient)
					break
				}
			}
		}
	}

	return strings.Join(ingredients, ",")
}

// PartOne How many times do allergen free ingredients appear
func PartOne() string {
	l := list{}
	l.load("twentyone/input.txt")

	l.isolateIngredientAllergens()
	count, ingredients := l.countAllergenFreeAppearances()
	return fmt.Sprintf("There are %d appearances of %d allergen free ingredients", count, len(ingredients))
}

// PartTwo What is the canonical dangerous list of ingredients
func PartTwo() string {
	l := list{}
	l.load("twentyone/input.txt")

	l.isolateIngredientAllergens()

	return fmt.Sprintf("The canonical list of dangerous ingredients is %q", l.canonicalDangerList())
}
