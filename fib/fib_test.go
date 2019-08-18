package fib

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var fibs = []int{0, 1, 1, 2, 3, 5, 8, 13, 21}

func TestFib(t *testing.T) {
	wg := sync.WaitGroup{}

	// TODO - don't pass time.Now() into goroutines, call time.Now at the start of goroutines
	for i := 0; i < len(fibs); i++ {
		wg.Add(1)
		go func(i int, t0 time.Time){
			f := Fib(i)
			//fmt.Println(i, "->", f)
			fmt.Printf("i: %d, Fib elasped %s\n", i, time.Now().Sub(t0))
			if f != fibs[i] {
				t.Errorf("Fib(%d): expected %d, got %d\n", i, fibs[i], f)
			}
			wg.Done()
		}(i, time.Now())

		wg.Add(1)
		go func(i int, t0 time.Time){
			f := FibRec(i)
			//fmt.Println(i, "->", f)
			fmt.Printf("i: %d, FibRec elasped %s\n", i, time.Now().Sub(t0))
			if f != fibs[i] {
				t.Errorf("FibRec(%d): expected %d, got %d\n", i, fibs[i], f)
			}
			wg.Done()
		}(i, time.Now())

		wg.Add(1)
		go func(i int, t0 time.Time){
			f := FibRecCache(i)
			//fmt.Println(i, "->", f)
			fmt.Printf("i: %d, FibRecCache elasped %s\n", i, time.Now().Sub(t0))
			if f != fibs[i] {
				t.Errorf("FibRecCache(%d): expected %d, got %d\n", i, fibs[i], f)
			}
			wg.Done()
		}(i, time.Now())

		// fmt.Println(i, "->", fib.FibRec(i))
		// fmt.Println(i, "->", fib.FibRecCache(i))
		wg.Wait()
	}

}