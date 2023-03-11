package lasagna_master

func PreparationTime(layers []string, layerPrepTime int) int {
	if layerPrepTime == 0 {
		layerPrepTime = 2
	}
	return layerPrepTime * len(layers)
}

func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendList []string, myList []string) []string {
	for idx, layer := range myList {
		if layer == "?" {
			myList[idx] = friendList[len(friendList)-1]
		}
	}
	return myList
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaled := make([]float64, len(quantities))
	for idx, quantity := range quantities {
		scaled[idx] = quantity * float64(portions) / 2.0
	}
	return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
