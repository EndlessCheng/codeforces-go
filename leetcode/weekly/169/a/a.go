package main

func sumZero1(n int) []int {
	a := make([]int, n)
	m := n / 2
	for i := range m {
		a[i] = i + 1
		a[i+m] = -i - 1
	}
	return a
}

func sumZero(n int) []int {
	a := make([]int, n)
	a[0] = -n * (n - 1) / 2
	for i := 1; i < n; i++ {
		a[i] = i
	}
	return a
}
