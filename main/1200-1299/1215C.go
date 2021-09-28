package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1215C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var s, t []byte
	Fscan(in, &n, &s, &t)

	pos := [2][]int{}
	for i, b := range s {
		if b != t[i] {
			pos[b-'a'] = append(pos[b-'a'], i+1)
		}
	}

	a, b := pos[0], pos[1]
	m := len(a) + len(b)
	if m&1 > 0 {
		Fprint(out, -1)
		return
	}
	Fprintln(out, m/2+len(a)&1)
	for _, ps := range pos {
		for i := 1; i < len(ps); i += 2 {
			Fprintln(out, ps[i-1], ps[i])
		}
	}
	if len(a)&1 > 0 {
		p := a[len(a)-1]
		Fprintln(out, p, p)
		Fprintln(out, p, b[len(b)-1])
	}
}

//func main() { CF1215C(os.Stdin, os.Stdout) }
