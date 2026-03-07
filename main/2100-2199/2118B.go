package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2118B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fprintln(out, n*2-1)
		Fprintln(out, 1, 1, n)
		for i := 2; i <= n; i++ {
			Fprintln(out, i, 1, i-1)
			Fprintln(out, i, i, n)
		}
	}
}

//func main() { cf2118B(os.Stdin, os.Stdout) }
