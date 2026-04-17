package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2072C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, or int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &or)
		lb0 := (or + 1) &^ or
		if n >= lb0 {
			for i := range lb0 {
				Fprint(out, i, " ")
			}
			for range n - lb0 {
				Fprint(out, or, " ")
			}
			Fprintln(out)
		} else {
			s := 0
			for i := range n - 1 {
				s |= i
				Fprint(out, i, " ")
			}
			if s|(n-1) == or {
				Fprintln(out, n-1)
			} else {
				Fprintln(out, or)
			}
		}
	}
}

//func main() { cf2072C(bufio.NewReader(os.Stdin), os.Stdout) }
