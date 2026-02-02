package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1474E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)

		c := 0
		for i := n/2 + 1; i < n; i++ {
			c += (i - 1) * (i - 1)
		}
		for i := 1; i <= n/2; i++ {
			c += (n - i) * (n - i)
		}
		Fprintln(out, c)

		Fprint(out, n/2+1, " ")
		for i := 1; i < n/2; i++ {
			Fprint(out, i, " ")
		}
		for i := n/2 + 2; i <= n; i++ {
			Fprint(out, i, " ")
		}
		Fprintln(out, n/2)

		Fprintln(out, n-1)
		for i := n/2 + 1; i < n; i++ {
			Fprintln(out, i, 1)
		}
		for i := n / 2; i > 0; i-- {
			Fprintln(out, i, n)
		}
	}
}

//func main() { cf1474E(bufio.NewReader(os.Stdin), os.Stdout) }
