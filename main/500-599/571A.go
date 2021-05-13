package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF571A(in io.Reader, out io.Writer) {
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}
	a, l := [3]int64{}, int64(0)
	Fscan(in, &a[0], &a[1], &a[2], &l)
	sum := a[0] + a[1] + a[2]
	ans := (l + 1) * (l + 2) * (l + 3) / 6
	for i := int64(0); i <= l; i++ {
		for _, v := range a {
			if x := min(v*2-sum+i, l-i); x >= 0 {
				ans -= (x + 1) * (x + 2) / 2
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF571A(os.Stdin, os.Stdout) }
