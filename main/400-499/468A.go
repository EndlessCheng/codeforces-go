package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF468A(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	if n < 4 {
		Fprint(out, "NO")
		return
	}
	Fprintln(out, "YES")
	for i := 5 + n&1; i < n; i += 2 {
		Fprintln(out, i+1, "-", i, "= 1\n1 * 1 = 1")
	}
	if n&1 > 0 {
		Fprintln(out, "2 - 1 = 1\n4 * 5 = 20\n20 + 1 = 21\n21 + 3 = 24")
	} else {
		Fprintln(out, "1 * 2 = 2\n2 * 3 = 6\n6 * 4 = 24")
	}
}

//func main() { CF468A(os.Stdin, os.Stdout) }
