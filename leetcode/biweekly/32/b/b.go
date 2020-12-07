package main

// github.com/EndlessCheng/codeforces-go
func canConvertString(s, t string, k int) (ans bool) {
	if len(s) != len(t) {
		return
	}
	c := [26]int{}
	mxI := byte(0)
	for i := range s {
		d := t[i] - s[i]
		if d > 26 {
			d += 26
		}
		if c[d]++; c[d] > c[mxI] || c[d] == c[mxI] && d > mxI {
			mxI = d
		}
	}
	return mxI == 0 || 26*(c[mxI]-1)+int(mxI) <= k
}
