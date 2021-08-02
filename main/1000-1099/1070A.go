package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1070A(in io.Reader, out io.Writer) {
	var d, s int
	Fscan(in, &d, &s)
	type pair struct{ d, s int }
	type pd struct {
		pair
		digit byte
	}

	vis := make([][]bool, d)
	from := make([][]pd, d)
	for i := range vis {
		vis[i] = make([]bool, s+1)
		from[i] = make([]pd, s+1)
	}
	vis[0][0] = true
	q := []pair{{}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for i := 0; i < 10; i++ {
			nd, ns := (p.d*10+i)%d, p.s+i
			if ns <= s && !vis[nd][ns] {
				vis[nd][ns] = true
				from[nd][ns] = pd{p, byte(i)}
				q = append(q, pair{nd, ns})
			}
		}
	}
	if !vis[0][s] {
		Fprint(out, -1)
		return
	}

	ans := []byte{}
	for d, s := 0, s; s > 0; {
		f := from[d][s]
		ans = append(ans, '0'+f.digit)
		d, s = f.d, f.s
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	Fprintf(out, "%s", ans)
}

//func main() { CF1070A(os.Stdin, os.Stdout) }
