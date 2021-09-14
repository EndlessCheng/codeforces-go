package main

// github.com/EndlessCheng/codeforces-go
func canFormArray(a []int, pieces [][]int) (ans bool) {
o:
	for len(a) > 0 {
		for _, p := range pieces {
			if p[0] == a[0] {
				for i, v := range p {
					if v != a[i] {
						return
					}
				}
				a = a[len(p):]
				continue o
			}
		}
		return
	}
	return true
}
