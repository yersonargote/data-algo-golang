package main

import (
	"fmt"

	"github.com/yersonargote/data-algo-golang/arrays"
)

func main() {
	var scores []int
	var max int

	scores = arrays.Fill(scores)
	fmt.Println(scores)

	max = arrays.Max(scores)
	fmt.Println(max)

	numbers := arrays.FillList()
	arrays.PrintList(numbers)
}
