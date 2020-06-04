package fib

// In it's purest form there is no special treatment of 2
// The pre-established values are 0 -> 0, 1 -> 1
// Two is really fib(n-1) + fib(n-2) --> fib(1) + fib(0) => 1
func FibMemo(n int) int {
	if n < 1 {
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

// Solution for staircase example with up to 3 steps at a time
// TODO - Medium article
// In it's purest form there is no special treatment of 2
// The pre-established values are 0 -> 0, 1 -> 1
// Two is really fib(n-1) + fib(n-2) --> fib(1) + fib(0) => 1
func Fib3Memo(n int) int {
	if n < 1 {
		return 0
	}
	if n <= 2 {
		return 1
	}

	memo := make(map[int]int, n) // preallocate the map's capacity
	memo[1] = 1
	memo[2] = 1
	memo[3] = 2

	for i := 4; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2] + memo[i-3]
	}

	return memo[n]
}
