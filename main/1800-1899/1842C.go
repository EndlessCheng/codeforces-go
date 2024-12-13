package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1842C(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		mx := make([]int, n+1)
		for i := range mx {
			mx[i] = -1e9
		}
		f := 0
		for i := range n {
			Fscan(in, &v)
			f, mx[v] = max(f, mx[v]+i+1), max(mx[v], f-i)
		}
		Fprintln(out, f)
	}
}

//func main() { cf1842C(bufio.NewReader(os.Stdin), os.Stdout) }
