package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF986C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v int
	Fscan(in, &n, &m)
	mask := 1<<n - 1
	has := make([]bool, 1<<n)
	for ; m > 0; m-- {
		Fscan(in, &v)
		has[v] = true
	}
	vis := make([]bool, 1<<n)
	var f func(int)
	f = func(v int) {
		if vis[v] {
			return
		}
		vis[v] = true
		if has[v] {
			f(mask ^ v)
		}
		for s := v; s > 0; {
			lb := s & -s
			f(v ^ lb)
			s ^= lb
		}
	}
	ans := 0
	for i, b := range has {
		if b && !vis[i] {
			ans++
			f(mask ^ i)
		}
	}
	Fprint(out, ans)
}

//func main() { CF986C(os.Stdin, os.Stdout) }
