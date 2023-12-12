package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1730B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, t int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		mn, mx := int(1e8), 0
		for _, x := range a {
			Fscan(in, &t)
			mn = min(mn, x-t)
			mx = max(mx, x+t)
		}
		Fprintf(out, "%.1f\n", float64(mn+mx)/2)
	}
}

//func main() { cf1730B(os.Stdin, os.Stdout) }
