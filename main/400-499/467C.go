package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF467C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}

	var n, m, k int
	Fscan(in, &n, &m, &k)
	s := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		s[i] += s[i-1]
	}
	pre := make([]int64, n+1)
	f := make([]int64, n+1)
	for ; k > 0; k-- {
		for j := m; j <= n; j++ {
			f[j] = max(f[j-1], pre[j-m]+s[j]-s[j-m])
		}
		pre, f = f, pre
	}
	Fprint(out, pre[n])
}

//func main() { CF467C(os.Stdin, os.Stdout) }
