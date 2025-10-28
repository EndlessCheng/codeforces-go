package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1806D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353
	var T, n, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		f, g := 0, 1
		for i := 1; i < n; i++ {
			Fscan(in, &x)
			f = (f*i + g*(1-x)) % mod
			g = g * (i - x) % mod
			Fprint(out, f, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1806D(bufio.NewReader(os.Stdin), os.Stdout) }
