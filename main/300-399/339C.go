package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF339C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	var m int
	Fscan(in, &s, &m)
	ans := make([]int, m)
	var f func(p, d int) bool
	f = func(p, d int) bool {
		if p == m {
			return true
		}
		for ans[p] = d + 1; ans[p] <= 10; ans[p]++ {
			if s[ans[p]-1] == '1' && (p == 0 || ans[p] != ans[p-1]) && f(p+1, ans[p]-d) {
				return true
			}
		}
		return false
	}
	if f(0, 0) {
		Fprintln(out, "YES")
		for _, v := range ans {
			Fprint(out, v, " ")
		}
	} else {
		Fprint(out, "NO")
	}
}

//func main() { CF339C(os.Stdin, os.Stdout) }
