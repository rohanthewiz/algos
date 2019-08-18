package fib

// Fib series
// 0 1 2 3 4 5 6 7
// 0 1 1 2 3 5 8 13
//   - -
//     - -

func Fib(n int) (fb int) {
	if n == 0 { return 0 }
	if n == 1 { return 1 }

	// fb hold the current fib - inited to zero
	// a holds the trailing value
	a := 0 // holds the trailing value
	fb = 1
	for i := 2;	i <= n ; i++  {
		//fmt.Println("n", n, "i", i, "a", a, "fb", fb)
		b := fb // save previous
		fb = b + a // fib is the previous plus the trailing
		a = b // trailing becomes the previous
	}
	return
}

func FibRec(n int) (fb int) {
	if n == 0 { return 0 }
	if n == 1 { return 1 }

	return FibRec(n - 1) + FibRec(n - 2)
}

func FibRecCache(n int) (fib int) {
	if n == 0 { return 0 }
	if n == 1 { return 1 }

	cache := map[int]int {
		0: 0,
		1: 1,
	}

	if v, ok := cache[n]; ok {
		return v
	}

	fib = FibRec(n - 1) + FibRec(n - 2)
	cache[n] = fib

	return
}
