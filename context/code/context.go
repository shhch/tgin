package code

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func RunTest() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	num := CountTest(ctx)
	for n := range num {
		if n >= 10 {
			cancel()
			break
		}
	}

	fmt.Println("正在统计结果...")
	wg.Wait()

}

func CountTest(ctx context.Context) <-chan int {
	c := make(chan int)
	n := 0
	t := 0
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("耗时 %d 秒，完成了 %d 次 \n", t, n)
				wg.Done()
				return
			case c <- n:
				n += rand.Intn(5)

				if n > 10 {
					n = 10
					continue
				} else {
					fmt.Printf("完成 %d 次\n", n)
					t++
				}

			}
		}
	}()

	return c
}
