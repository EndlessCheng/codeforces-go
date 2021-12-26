package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF625C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	k--
	Fprintln(out, ((n+k-1)*n+k+2)*n/2)
	v, w := 1, n*k+1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j < k {
				Fprint(out, v, " ")
				v++
			} else {
				Fprint(out, w, " ")
				w++
			}
		}
		Fprintln(out)
	}
}

//func main() { CF625C(os.Stdin, os.Stdout) }
