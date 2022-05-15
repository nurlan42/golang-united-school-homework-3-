package homework

func average(input [15]float32) (result float32) {
	var ok, d float32

	for _, el := range input {
		if el != 0 {
			ok += el
			d++
		}
	}

	result = ok / d
	return
}
