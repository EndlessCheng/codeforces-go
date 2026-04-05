package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1793B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, r, l int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &r, &l)
		Fprintln(out, (r-l)*2)
		for x := l; x <= r; x++ {
			Fprint(out, x, " ")
		}
		for x := r - 1; x > l; x-- {
			Fprint(out, x, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1793B(bufio.NewReader(os.Stdin), os.Stdout) }
