package fib

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// fibs holds the fib for each index
var fibs = []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
var fib3s = []int{0, 1, 1, 2, 4, 7, 13, 24, 44, 81}

func TestFib(t *testing.T) {
	wg := sync.WaitGroup{}

	for i := 0; i < len(fibs); i++ {
		wg.Add(1)
		go func(i int, t0 time.Time) {
			f := Fib(i)
			fmt.Printf("i: %d, Fib elasped %s\n", i, time.Now().Sub(t0))
			if f != fibs[i] {
				t.Errorf("Fib(%d): expected %d, got %d\n", i, fibs[i], f)
			}
			wg.Done()
		}(i, time.Now())

		wg.Add(1)
		go func(i int, t0 time.Time) {
			f := FibRec(i)
			fmt.Printf("i: %d, FibRec elasped %s\n", i, time.Now().Sub(t0))
			if f != fibs[i] {
				t.Errorf("FibRec(%d): expected %d, got %d\n", i, fibs[i], f)
			}
			wg.Done()
		}(i, time.Now())

		wg.Add(1)
		go func(i int, t0 time.Time) {
			f := FibRecCache(i)
			fmt.Printf("i: %d, FibRecCache elasped %s\n", i, time.Now().Sub(t0))
			if f != fibs[i] {
				t.Errorf("FibRecCache(%d): expected %d, got %d\n", i, fibs[i], f)
			}
			wg.Done()
		}(i, time.Now())

		wg.Add(1)
		go func(i int, t0 time.Time) {
			f := FibMemo(i)
			fmt.Printf("i: %d, FibMemo elasped %s\n", i, time.Now().Sub(t0))
			if f != fibs[i] {
				t.Errorf("FibMemo(%d): expected %d, got %d\n", i, fibs[i], f)
			}
			wg.Done()
		}(i, time.Now())

		wg.Add(1)
		go func(i int, t0 time.Time) {
			f := Fib3Memo(i)
			fmt.Printf("i: %d, Fib3Memo elasped %s\n", i, time.Now().Sub(t0))
			if f != fib3s[i] {
				t.Errorf("FibM3emo(%d): expected %d, got %d\n", i, fib3s[i], f)
			}
			wg.Done()
		}(i, time.Now())

		wg.Wait()
	}
}
