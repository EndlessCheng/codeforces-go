package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

func CF1092C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	var s string
	Fscan(in, &n)
	type pair struct { s string; i int }
	a := make([][]pair, n-1)
	for i := 0; i < n*2-2; i++ {
		Fscan(in, &s)
		a[len(s)-1] = append(a[len(s)-1], pair{s, i})
	}

	ans := make([]byte, n*2-2)
	p, s := a[n-2][0].s, a[n-2][1].s
next:
	for _, comb := range []string{p + s[len(s)-1:], s + p[len(p)-1:]} {
		for _, p := range a {
			if strings.HasPrefix(comb, p[0].s) && strings.HasSuffix(comb, p[1].s) {
				ans[p[0].i], ans[p[1].i] = 'P', 'S'
			} else if strings.HasSuffix(comb, p[0].s) && strings.HasPrefix(comb, p[1].s) {
				ans[p[0].i], ans[p[1].i] = 'S', 'P'
			} else {
				continue next
			}
		}
		Fprint(out, string(ans))
		return
	}
}

//func main() { CF1092C(os.Stdin, os.Stdout) }
