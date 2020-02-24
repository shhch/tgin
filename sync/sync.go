package sync

import (
	"fmt"
	"sync"
)

func Test()  {
	sw := sync.WaitGroup{}
	sw.Add(20)
	for i :=0; i < 20; i++ {
		go f(i, &sw)
	}
	sw.Wait()
}

// wg对象不是引用类型，需要传地址
func f(i int, wg *sync.WaitGroup) {
	fmt.Println(i)
	wg.Done()
}
