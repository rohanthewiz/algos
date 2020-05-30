package fib

func FibMemo(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}

	memo := make(map[int]int, n) // preallocate the map's capacity
	memo[1] = 1
	memo[2] = 1

	for i := 3; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	return memo[n]
}
