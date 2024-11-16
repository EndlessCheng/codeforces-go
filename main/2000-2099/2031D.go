package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2031D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		suf := make([]int, n)
		suf[n-1] = a[n-1]
		for i := n - 2; i >= 0; i-- {
			suf[i] = min(suf[i+1], a[i])
		}

		for i := 0; i < n; {
			st := i
			mx := a[i]
			for i++; i < n && mx > suf[i]; i++ {
				mx = max(mx, a[i])
			}
			for range i - st {
				Fprint(out, mx, " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { cf2031D(bufio.NewReader(os.Stdin), os.Stdout) }
