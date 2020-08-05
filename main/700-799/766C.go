package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF766C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod int64 = 1e9 + 7

	var n int
	var s []byte
	Fscan(in, &n, &s)
	for i := range s {
		s[i] -= 'a'
	}
	a := [26]int{}
	for i := range a {
		Fscan(in, &a[i])
	}

	mxLen := 0
	dp := make([]int64, n)
	var f func(int) int64
	f = func(p int) (res int64) {
		if p == n {
			return 1
		}
		if dp[p] > 0 {
			return dp[p]
		}
		defer func() { dp[p] = res }()
		curMx := int(1e9)
		for i := p; i < n; i++ {
			if a[s[i]] < curMx {
				curMx = a[s[i]]
			}
			if i-p+1 > curMx {
				break
			}
			if i-p+1 > mxLen {
				mxLen = i - p + 1
			}
			res = (res + f(i+1)) % mod
		}
		return
	}
	Fprintln(out, f(0))
	Fprintln(out, mxLen)
	num := 0
	for i := 0; i < n; {
		curMx := int(1e9)
		for st := i; i < n; i++ {
			if a[s[i]] < curMx {
				curMx = a[s[i]]
			}
			if i-st+1 > curMx {
				break
			}
		}
		num++
	}
	Fprint(out, num)
}

//func main() { CF766C(os.Stdin, os.Stdout) }
