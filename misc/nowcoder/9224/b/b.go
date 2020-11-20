package main

// github.com/EndlessCheng/codeforces-go
func tree2(k int, a []int) int64 {
	q := []int{a[0]}
	a = a[1:]
	ans := 0
	for len(a) > 0 {
		qq := q
		q = nil
		for _, v := range qq {
			for i := 0; i < k && len(a) > 0; i++ {
				q = append(q, a[0])
				ans += v ^ a[0]
				a = a[1:]
			}
		}
	}
	return int64(ans)
}
