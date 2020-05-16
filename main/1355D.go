package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1355D(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, s int
	Fscan(_r, &n, &s)
	if s < 2*n {
		Fprint(out, "NO")
	} else {
		Fprintln(out, "YES")
		for i := 1; i < n; i++ {
			Fprint(out, "1 ")
		}
		Fprintln(out, s-n+1)
		Fprint(out, n)
	}
}

//func main() { CF1355D(os.Stdin, os.Stdout) }
