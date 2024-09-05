package arrayGenerate

import (
	"math/rand/v2"
)

func arrayGenerate(n int) []int {
	var array = make([]int, n)

	for i := 0; i < n; i++ {
		array[i] = rand.IntN(10000)
	}

	return array

}
