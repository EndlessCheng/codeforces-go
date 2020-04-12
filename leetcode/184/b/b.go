package main

func processQueries(qs []int, m int) (ans []int) {
	a := make([]int, m)
	for i := range a {
		a[i] = i + 1
	}
	for _, q := range qs {
		for i, v := range a {
			if v == q {
				ans = append(ans, i)
				a = append(append([]int{v}, a[:i]...), a[i+1:]...)
				break
			}
		}
	}
	return
}
