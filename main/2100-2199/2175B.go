package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2175B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &l, &r)
		get := func(i int) int {
			if i == r {
				return l - 1
			}
			return i
		}
		for i := range n {
			Fprint(out, get(i)^get(i+1), " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2175B(bufio.NewReader(os.Stdin), os.Stdout) }
