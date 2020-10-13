package main

import (
	"bufio"
	. "fmt"
	"io"
)

// 与值域无关的解法 https://codeforces.com/blog/entry/78016#comment-631749

// github.com/EndlessCheng/codeforces-go
func CF1359D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for mx := 1; mx <= 30; mx++ {
		i := 0
		for {
			for ; i < n && a[i] > mx; i++ {
			}
			if i == n {
				break
			}
			s, mxS := a[i], a[i]
			for i++; i < n && a[i] <= mx; i++ {
				s = max(s+a[i], a[i])
				mxS = max(mxS, s)
			}
			ans = max(ans, mxS-mx)
		}
	}
	Fprint(out, ans)
}

//func main() { CF1359D(os.Stdin, os.Stdout) }
