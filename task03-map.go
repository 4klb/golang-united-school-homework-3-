package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	var arr []string
	keys := make([]int, 0, len(input))

	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		arr = append(arr, input[k])
	}

	return arr
}
