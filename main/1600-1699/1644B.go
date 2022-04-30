package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1644B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fprint(out, 1)
		for i := n; i > 1; i-- {
			Fprint(out, " ", i)
		}
		Fprintln(out)
		for i := n; i > 1; i-- {
			j := n
			for ; j > i; j-- {
				Fprint(out, j, " ")
			}
			Fprint(out, j-1, j, " ")
			for j -= 2; j > 0; j-- {
				Fprint(out, j, " ")
			}
			Fprintln(out)
		}
	}
}

//func main() { CF1644B(os.Stdin, os.Stdout) }
