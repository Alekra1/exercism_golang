package lasagna

import "strings"

// TODO: define the 'PreparationTime()' function

func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		return len(layers) * 2
	}
	return len(layers) * avgTime
}

// TODO: define the 'Quantities()' function

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		if strings.Contains(layer, "noodles") {
			noodles += 50
		} else if strings.Contains(layer, "sauce") {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(friendList, ownList []string) {
	ownList[len(ownList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe(amountsNeeded []float64, desiredPortions int) (scaledAmounts []float64) {
	for _, v := range amountsNeeded {
		scaledAmounts = append(scaledAmounts, (v/2)*float64(desiredPortions))
	}
	return
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
