package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF552E(in io.Reader, out io.Writer) {
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}

	var s []byte
	Fscan(bufio.NewReader(in), &s)
	a := bytes.Split(s, []byte{'+'})
	n := len(a)
	sum := make([]int64, n+1)
	muls := make([]int64, n)
	for i, s := range a {
		muls[i] = 1
		for j := 0; j < len(s); j += 2 {
			muls[i] *= int64(s[j] & 15)
		}
		sum[i+1] = sum[i] + muls[i]
	}
	ans := sum[n]
	for l, p := range a {
		for r := l + 1; r < n; r++ {
			q := a[r]
			if len(q) == 1 {
				continue
			}
			res := sum[l] + sum[n] - sum[r+1]
			for j, mulQ := 0, int64(1); j < len(q)-2; j += 2 {
				mulQ *= int64(q[j] & 15)
				mid := (sum[r] - sum[l] + mulQ) * (muls[r] / mulQ)
				ans = max(ans, res+mid)
			}
		}
		if len(p) == 1 {
			continue
		}
		for i, mulP := 0, int64(1); i < len(p)-2; i += 2 {
			mulP *= int64(p[i] & 15)
			mulP2 := muls[l] / mulP
			for r := l + 1; r < n; r++ {
				res := sum[l] + sum[n] - sum[r+1]
				q := a[r]
				if len(q) == 1 {
					mid := mulP * (mulP2 + sum[r+1] - sum[l+1])
					ans = max(ans, res+mid)
					continue
				}
				for j, mulQ := 0, int64(1); j < len(q)-2; j += 2 {
					mulQ *= int64(q[j] & 15)
					mid := mulP * (mulP2 + sum[r] - sum[l+1] + mulQ) * (muls[r] / mulQ)
					ans = max(ans, res+mid)
				}
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF552E(os.Stdin, os.Stdout) }

