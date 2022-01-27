package lasagna

func PreparationTime(layers []string, averagePrepTime int) int {
	// default to 2 for prep time, 0 being assumed to be "default"
	if averagePrepTime == 0 {
		averagePrepTime = 2
	}
	return len(layers) * averagePrepTime
}

func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64
	for layer := 0; layer < len(layers); layer++ {
		if layers[layer] == "sauce" {
			sauce += 0.2
		}
		if layers[layer] == "noodles" {
			noodles += 50
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendList []string, myList []string) []string {
	newList := make([]string, len(myList))
	copy(newList, myList)
	return append(newList, friendList[len(friendList)-1])
}

// Given a slice of amounts for 2 portions returns a new slice of scaled amounts based on new portions
func ScaleRecipe(amounts []float64, portions int) []float64 {
	var scaledQuanities []float64
	for i := 0; i < len(amounts); i++ {
		scaledQuanities = append(scaledQuanities, amounts[i]*float64(portions)/2.0)
	}
	return scaledQuanities
}
