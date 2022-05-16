package homework

func Reverse(input []int64) (result []int64) {
	// slice := make([]int64, len(input))
	// copy(slice, input)
	var slice []int64

	for i := len(input) - 1; i >= 0; i-- {
		slice = append(slice, input[i])
	}

	return slice
}
