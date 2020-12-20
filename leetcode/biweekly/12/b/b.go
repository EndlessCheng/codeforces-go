package main

// github.com/EndlessCheng/codeforces-go
func transformArray(a []int) (ans []int) {
	for {
		change := false
		b := append([]int(nil), a...)
		for i := 1; i < len(a)-1; i++ {
			if a[i-1] > a[i] && a[i] < a[i+1] {
				b[i]++
				change = true
			} else if a[i-1] < a[i] && a[i] > a[i+1] {
				b[i]--
				change = true
			}
		}
		if !change {
			return b
		}
		a = b
	}
}
