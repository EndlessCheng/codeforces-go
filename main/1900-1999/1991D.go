package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1991D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	a := []any{1, 2, 2, 3, 3}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n < 6 {
			Fprintln(out, n/2+1)
			Fprintln(out, a[:n]...)
		} else {
			Fprintln(out, 4)
			for i := range n {
				Fprint(out, i%4+1, " ")
			}
			Fprintln(out)
		}
	}
}

//func main() { cf1991D(bufio.NewReader(os.Stdin), os.Stdout) }
