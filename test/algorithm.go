package main

import (
	"fmt"
	"tgin/algorithm"
)

func main()  {
	arr := []int{13,25,2,16,45,5,0,6,1}
	algorithm.BubbleSort(arr)
	//algorithm.CocktailSort(arr)
	fmt.Println(arr)
}
