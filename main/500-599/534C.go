package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF534C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, s, sa int
	Fscan(in, &n, &s)
	s -= n
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		a[i]--
		sa += a[i]
	}
	for _, v := range a {
		mn := max(s-sa+v, 0)
		mx := min(s, v)
		Fprint(out, v-mx+mn, " ")
	}
}

//func main() { CF534C(os.Stdin, os.Stdout) }
