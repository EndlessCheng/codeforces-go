package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1174D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, x int
	Fscan(in, &n, &x)
	if m := 1 << n; x < m {
		m >>= 1
		Fprintln(out, m-1)
		mask := x&-x - 1
		for i := 1; i < m; i++ {
			v := i & -i
			Fprint(out, v&^mask<<1|v&mask, " ")
		}
	} else {
		Fprintln(out, m-1)
		for i := 1; i < m; i++ {
			Fprint(out, i&-i, " ")
		}
	}
}

//func main() { CF1174D(os.Stdin, os.Stdout) }
