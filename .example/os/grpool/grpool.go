package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/mickbai/gf/os/grpool"
	"github.com/mickbai/gf/os/gtime"
)

func main() {
	start := gtime.TimestampMilli()
	wg := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		grpool.Add(func() {
			time.Sleep(time.Second)
			wg.Done()
		})
	}
	wg.Wait()
	fmt.Println(grpool.Size())
	fmt.Println("time spent:", gtime.TimestampMilli()-start)
}
