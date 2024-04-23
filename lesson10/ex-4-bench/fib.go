package fib

func FibRecursive(n int) int {
	if n <= 1 {
		return n
	}

	return FibRecursive(n-1) + FibRecursive(n-2)
}

func FibIterative(n int) int {
	a, b := 0, 1

	for range n {
		a, b = b, a+b
	}

	return a
}
