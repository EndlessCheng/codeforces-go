package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF739A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, l, r int
	Fscan(in, &n, &m)
	mi := n
	for ; m > 0; m-- {
		if Fscan(in, &l, &r); r-l+1 < mi {
			mi = r - l + 1
		}
	}
	Fprintln(out, mi)
	for i := 0; i < n; i++ {
		Fprint(out, i%mi, " ")
	}
}

//func main() { CF739A(os.Stdin, os.Stdout) }
