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
				tmp := append([]int{v}, a[:i]...)
				tmp = append(tmp, a[i+1:]...)
				a = tmp
				break
			}
		}
	}
	return
}
