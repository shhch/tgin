package main

import (
	"fmt"
	"tgin/algorithm/code"
)

func main()  {
	arr := []int{13,25,2,16,45,5,0,6,1}

	code.BubbleSort(arr)
	fmt.Println(arr)

	code.CocktailSort(arr)
	fmt.Println(arr)

	code.QuickSort(arr,0, len(arr) -1)
	fmt.Println(arr)
}
