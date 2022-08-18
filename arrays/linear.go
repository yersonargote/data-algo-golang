package arrays

import (
	"container/list"
	"fmt"

	"github.com/yersonargote/data-algo-golang/utils"
)

/*
Learn Data Structures and Algo. with Golang
Bhagvan Kommadi
*/

// Create, fill and return a list of numbers
func FillList() list.List {
	var numbers list.List
	for i := 0; i < 10; i++ {
		numbers.PushBack(utils.Random())
	}
	return numbers
}

func PrintList(list list.List) {
	for element := list.Front(); element != nil; element = element.Next() {
		fmt.Printf("%d ", element.Value.(int))
	}
	fmt.Println()
}
