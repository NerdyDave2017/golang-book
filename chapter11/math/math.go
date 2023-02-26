package math

func Average(input []float64) float64 {
	total := float64(0)
	for _, value := range input {
		total += value
	}
	return total / float64(len(input))
}

func Max(input []float64) float64 {
	var temp float64 = input[0]

	for _, v := range input {
		if v >= temp {
			temp = v
		}
	}

	return temp
}

func Min(input []float64) float64 {
	var temp float64 = input[0]

	for _, v := range input {
		if v <= temp {
			temp = v
		}
	}

	return temp
}
