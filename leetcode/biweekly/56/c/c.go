package main

// 原题 https://codeforces.com/problemset/problem/1215/D

// github.com/EndlessCheng/codeforces-go
func sumGame(s string) (ans bool) {
	var sl, cl, sr, cr int
	n := len(s)
	for _, b := range s[:n/2] {
		if b != '?' {
			sl += int(b & 15)
		} else {
			cl++
		}
	}
	for _, b := range s[n/2:] {
		if b != '?' {
			sr += int(b & 15)
		} else {
			cr++
		}
	}
	if (cl+cr)&1 > 0 {
		return true
	}
	if sl == sr {
		return cl != cr
	}
	if cl > cr {
		sl, cl, sr, cr = sr, cr, sl, cl
	}
	return sl != sr+(cr-cl)/2*9
}
