package homework

func reverse(input []int64) (result []int64) {
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-i-1; j++ {
			if input[j] < input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	result = input
	return
}
