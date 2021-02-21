package main

// github.com/EndlessCheng/codeforces-go
func canChoose(groups [][]int, a []int) (ans bool) {
o:
	for _, b := range groups {
		for len(a) >= len(b) {
			if equal(a[:len(b)], b) {
				a = a[len(b):]
				continue o
			}
			a = a[1:]
		}
		return
	}
	return true
}

func equal(a, b []int) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
