package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1304D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func() {
		var n int
		var s string
		Fscan(in, &n, &s)
		ans := make([]int, n)
		up := n
		for i := range s {
			if s[i] == '>' {
				ans[i] = up
				up--
			}
		}
		down := 1
		for i := n - 1; i >= 0; i-- {
			if ans[i] > 0 {
				continue
			}
			end := i
			for ; i >= 0 && ans[i] == 0; i-- {
			}
			i++
			for j := i; j <= end; j++ {
				ans[j] = down
				down++
			}
		}
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)

		ans = make([]int, n)
		up = n
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == '<' {
				ans[i+1] = up
				up--
			}
		}
		down = 1
		for i := 0; i < n; i++ {
			if ans[i] > 0 {
				continue
			}
			end := i
			for ; i < n && ans[i] == 0; i++ {
			}
			i--
			for j := i; j >= end; j-- {
				ans[j] = down
				down++
			}
		}
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}

	var t int
	Fscan(in, &t)
	for _case := 1; _case <= t; _case++ {
		solve()
	}
}

//func main() { CF1304D(os.Stdin, os.Stdout) }
