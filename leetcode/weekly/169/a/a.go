package main

func sumZero(n int) []int {
	a := make([]int, n)
	m := n / 2
	for i := range m {
		a[i] = i + 1
		a[i+m] = -i - 1
	}
	return a
}
