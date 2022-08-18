package arrays

import "github.com/yersonargote/data-algo-golang/utils"

func Fill(array []int) []int {
	to := 10
	array = make([]int, to)
	for from := 0; from < to; from++ {
		random := utils.Random()
		array[from] = random
	}
	return array
}

func Max(array []int) int {
	for idx, value := range array {
		if idx < len(array)-1 {
			if value > array[idx+1] {
				temp := array[idx]
				array[idx] = array[idx+1]
				array[idx+1] = temp
			}
		}
	}
	max := array[len(array)-1]
	return max
}
