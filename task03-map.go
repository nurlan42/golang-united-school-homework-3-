package homework

func sortMapValues(input map[int]string) (result []string) {
	var mKeys []int
	for k := range input {
		mKeys = append(mKeys, k)
	}

	sorted := sortKeys(mKeys)
	for _, el := range sorted {
		result = append(result, input[el])
	}

	return
}

func sortKeys(input []int) []int {
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-i-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	return input
}
