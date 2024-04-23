package fib

import "testing"

func BenchmarkFibRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibRecursive(10)
	}
}

func BenchmarkFibIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibIterative(10)
	}
}

func TestFib(t *testing.T) {
	t.Run("FibRecursive", func(t *testing.T) {
		for _, test := range get_cases() {
			if got := FibRecursive(test.n); got != test.expected {
				t.Errorf("Fib(%d) = %d; want %d", test.n, got, test.expected)
			}
		}
	})

	t.Run("FibIterative", func(t *testing.T) {
		for _, test := range get_cases() {
			if got := FibIterative(test.n); got != test.expected {
				t.Errorf("Fib(%d) = %d; want %d", test.n, got, test.expected)
			}
		}
	})
}

func get_cases() []struct {
	n        int
	expected int
} {
	return []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{15, 610},
	}
}
