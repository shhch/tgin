package code

import "fmt"

func Set(index int, c chan bool)  {
	a := 0
	for i := 0; i < 100; i ++ {
		a += i
	}
	fmt.Println(index, a)
	c <- true
}

func Run() {
	c := make(chan bool, 10)
	for i := 0; i < 10 ; i++ {
		go Set(i, c)
	}
	for i := 0; i < 10 ; i++ {
		fmt.Println(<-c)
	}
}