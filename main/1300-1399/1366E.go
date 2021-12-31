package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1366E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 998244353

	var n, m, v int
	Fscan(in, &n, &m)
	type pair struct{ v, i int }
	s := []pair{{}}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		for s[len(s)-1].v >= v {
			s = s[:len(s)-1]
		}
		s = append(s, pair{v, i})
	}
	Fscan(in, &v)
	if s[1].v != v {
		Fprint(out, 0)
		return
	}
	s = append(s, pair{2e9, 0})
	ans := int64(1)
	for p := 2; m > 1; m-- {
		Fscan(in, &v)
		for s[p].v < v {
			p++
		}
		if s[p].v != v {
			Fprint(out, 0)
			return
		}
		ans = ans * int64(s[p].i-s[p-1].i) % mod
	}
	Fprint(out, ans)
}

//func main() { CF1366E(os.Stdin, os.Stdout) }
