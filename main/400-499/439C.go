package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF439C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, even, v int
	Fscan(in, &n, &k, &even)
	g := [2][]int{}
	for ; n > 0; n-- {
		Fscan(in, &v)
		g[v&1] = append(g[v&1], v)
	}
	e, o := g[0], g[1]

	odd := k - even
	if leftOdd := len(o) - odd; leftOdd < 0 || leftOdd&1 > 0 || len(e)+leftOdd/2 < even {
		Fprint(out, "NO")
		return
	}

	Fprintln(out, "YES")
	if even == 0 {
		odd--
	}
	for _, v := range o[:odd] {
		Fprintln(out, 1, v)
	}
	o = o[odd:]
	for i := 1; i < even; i++ {
		if len(e) > 0 {
			Fprintln(out, 1, e[0])
			e = e[1:]
		} else {
			Fprintln(out, 2, o[0], o[1])
			o = o[2:]
		}
	}
	e = append(e, o...)
	Fprint(out, len(e))
	for _, v := range e {
		Fprint(out, " ", v)
	}
}

//func main() { CF439C(os.Stdin, os.Stdout) }
