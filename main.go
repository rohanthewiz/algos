package main

import (
	"algos/fib"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 11; i++ {
		wg.Add(1)
		go func(i int, t time.Time){
			fmt.Println(i, "->", fib.Fib(i))
			fmt.Printf("i: %d, Fib elasped %s\n", i, time.Now().Sub(t))
			wg.Done()
		}(i, time.Now())

		wg.Add(1)
		go func(i int, t time.Time){
			fmt.Println(i, "->", fib.FibRec(i))
			fmt.Printf("i: %d, FibRec elasped %s\n", i, time.Now().Sub(t))
			wg.Done()
		}(i, time.Now())

		wg.Add(1)
		go func(i int,t time.Time){
			fmt.Println(i, "->", fib.FibRecCache(i))
			fmt.Printf("i: %d, FibRecCache elasped %s\n", i, time.Now().Sub(t))
			wg.Done()
		}(i, time.Now())

		// fmt.Println(i, "->", fib.FibRec(i))
		// fmt.Println(i, "->", fib.FibRecCache(i))
		wg.Wait()
	}
}
