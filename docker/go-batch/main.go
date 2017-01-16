package main

import (
	"fmt"
	"time"
	"runtime"
)

func wait(waitSec time.Duration) <-chan float64 {
	ch := make(chan float64)
	go func() {
		start := time.Now()
		time.Sleep(waitSec * time.Second)
		end := time.Now()
		ch <- end.Sub(start).Seconds()
	}()
	return ch
}

func main() {

	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	start := time.Now()

	ch1 := wait(5)
	ch2 := wait(10)
	ch3 := wait(3)

	ret1 := <-ch1
	ret2 := <-ch2
	ret3 := <-ch3

	fmt.Printf("各処理時間合計 %f sec\n", ret1+ret2+ret3)
	end := time.Now()

	fmt.Printf("実時間 %f sec \n", end.Sub(start).Seconds())
}