package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1547E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k, t int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		p := make([]int, k)
		for i := range p {
			Fscan(in, &p[i])
		}
		f := make([]int, n)
		for i := range f {
			f[i] = 1e18
		}
		for _, i := range p {
			Fscan(in, &t)
			f[i-1] = t
		}
		for i := n - 2; i >= 0; i-- {
			f[i] = min(f[i], f[i+1]+1)
		}
		mn := int(1e18)
		for _, f := range f {
			mn = min(f, mn+1)
			Fprint(out, mn, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1547E(os.Stdin, os.Stdout) }
