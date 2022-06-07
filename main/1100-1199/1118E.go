package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1118E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var m, n int
	Fscan(in, &m, &n)
	if int64(n)*int64(n-1) < int64(m) {
		Fprint(out, "NO")
		return
	}
	Fprintln(out, "YES")
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			Fprintln(out, i, j)
			if m--; m == 0 {
				return
			}
			Fprintln(out, j, i)
			if m--; m == 0 {
				return
			}
		}
	}
}

//func main() { CF1118E(os.Stdin, os.Stdout) }
