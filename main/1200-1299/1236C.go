package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1236C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j&1 > 0 {
				Fprint(out, j*n+n-i, " ")
			} else {
				Fprint(out, j*n+i+1, " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { CF1236C(os.Stdin, os.Stdout) }
