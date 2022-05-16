package homework

func Average(input [15]float32) (result float32) {
	var sum float32
	var counter int

	for i := 0; i < len(input); i++ {
		if input[i] != 0 {
			sum += input[i]
			counter++
		}

	}

	return sum / float32(counter)
}
