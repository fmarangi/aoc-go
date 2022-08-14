package day21

import (
	"sort"
	"strings"
)

type Recipe struct {
	Ingredients, Allergens []string
}

func flip(data map[string]string) map[string]string {
	result := make(map[string]string, len(data))
	for k, v := range data {
		result[v] = k
	}
	return result
}

func Allergens(recipes []Recipe) map[string]string {
	var (
		data       = make(map[string][]Set)
		allergens  = make(map[string]string)
		candidates = make(map[string]Set)
	)

	for _, recipe := range recipes {
		set := make(Set)
		set.AddAll(recipe.Ingredients)
		for _, a := range recipe.Allergens {
			data[a] = append(data[a], set)
		}
	}

	for a, values := range data {
		options := values[0]
		for _, i := range values[1:] {
			options = Intersect(options, i)
		}

		if len(options) == 1 {
			for i := range options {
				allergens[a] = i
			}
		} else {
			candidates[a] = options
		}
	}

	for n := len(data); len(allergens) != n; {
		for a, c := range candidates {
			for _, found := range allergens {
				delete(c, found)
			}

			if len(c) == 1 {
				for i := range c {
					allergens[a] = i
				}
				delete(candidates, a)
			}
		}
	}

	return allergens
}

func Part1(input string) (safe int) {
	var (
		recipes   = parseInput(input)
		allergens = flip(Allergens(recipes))
	)

	for _, recipe := range recipes {
		for _, i := range recipe.Ingredients {
			if _, allergen := allergens[i]; !allergen {
				safe++
			}
		}
	}

	return
}

func Part2(input string) string {
	var (
		recipes   = parseInput(input)
		allergens = Allergens(recipes)
		keys      []string
		canonical []string
	)

	for k := range allergens {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		canonical = append(canonical, allergens[k])
	}

	return strings.Join(canonical, ",")
}

func parseInput(input string) (recipes []Recipe) {
	for _, row := range strings.Split(input, "\n") {
		var (
			recipe Recipe
			data   = strings.Trim(row, ")")
			parts  = strings.Split(data, " (contains ")
		)

		recipe.Ingredients = strings.Fields(parts[0])
		recipe.Allergens = strings.Split(parts[1], ", ")
		recipes = append(recipes, recipe)
	}
	return
}
