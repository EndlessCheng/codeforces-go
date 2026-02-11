package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1772C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &k, &n)
		v := 1
		for i := range k {
			v = min(v+i, n-(k-1-i))
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1772C(os.Stdin, os.Stdout) }
