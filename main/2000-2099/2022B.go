package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2022B(in io.Reader, out io.Writer) {
	var T, n, x, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		s, mx := 0, 0
		for range n {
			Fscan(in, &v)
			s += v
			mx = max(mx, v)
		}
		Fprintln(out, max((s-1)/x+1, mx))
	}
}

//func main() { cf2022B(bufio.NewReader(os.Stdin), os.Stdout) }
