package main

import "fmt"

var ch1 chan int
var ch2 chan int
var chs = []chan int{ch1, ch2}
var numbers = []int{1, 2, 3, 4, 5}

func main() {

	// 1、select会执行一个分支
	// 2、如果有default，且其他均不满足时，会执行default，然后结束，如果没有default，则会阻塞，直到可以执行
	// 3、如果同时满足多个分支，则随机选取一个执行
	// 4、所有channel表达式都会被求值、所有被发送的表达式都会被求值。求值顺序：自上而下、从左到右
	// 待验证：当既存在均不满足，且无default的多channel表达式时（select阻塞），channel表达式会被执行几次
	select {
	case getChan(1) <- getNumber(1):
		fmt.Println("1th case is selected.")
	case getChan(1) <- getNumber(3):
		fmt.Println("2th case is selected.")
	default:
		fmt.Println("default!.")
	}
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("chs[%d]\n", i)
	fmt.Println("--", chs[i], "--")
	return chs[i]
}
