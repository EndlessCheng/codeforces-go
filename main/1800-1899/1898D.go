package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1898D(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var T, n, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		base := 0
		mn, mx := int(1e9), 0
		for _, v := range a {
			Fscan(in, &w)
			base += abs(v - w)
			mn = min(mn, max(v, w))
			mx = max(mx, min(v, w))
		}
		Fprintln(out, base+max(mx-mn, 0)*2)
	}
}

//func main() { cf1898D(bufio.NewReader(os.Stdin), os.Stdout) }
