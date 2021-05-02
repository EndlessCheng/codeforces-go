package main

// github.com/EndlessCheng/codeforces-go
func getMinSwaps(s string, k int) (ans int) {
	t := []byte(s)
	for i := 0; i < k; i++ {
		nextPermutation(t)
	}
	for ; s != ""; s = s[1:] {
		if t[0] == s[0] {
			t = t[1:]
			continue
		}
		j := 1
		for ; t[j] != s[0]; j++ {
		}
		ans += j
		t = append(t[:j], t[j+1:]...)
	}
	return
}

func nextPermutation(a []byte) {
	n := len(a)
	i := n - 2
	for i >= 0 && a[i] >= a[i+1] {
		i--
	}
	j := n - 1
	for j >= 0 && a[i] >= a[j] {
		j--
	}
	a[i], a[j] = a[j], a[i]
	reverse(a[i+1:])
}

func reverse(a []byte) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
