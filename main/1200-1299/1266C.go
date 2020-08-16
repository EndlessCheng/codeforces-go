package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1266C(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(_r, &n, &m)
	if n == 1 && m == 1 {
		Fprint(out, 0)
		return
	}
	if m == 1 {
		for i := 1; i <= n; i++ {
			Fprintln(out, i+1)
		}
		return
	}
	for i := 1; i <= n; i++ {
		for j := n + 1; j <= n+m; j++ {
			Fprint(out, i*j, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1266C(os.Stdin, os.Stdout) }
