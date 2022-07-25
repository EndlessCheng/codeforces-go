package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1119C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m int
	var v byte
	Fscan(in, &n, &m)
	var row, col, zero [500]byte
	for i := 0; i < n*2; i++ {
		for j := 0; j < m; j++ {
			Fscan(in, &v)
			row[i%n] ^= v
			col[j] ^= v
		}
	}
	if row == zero && col == zero {
		Fprint(out, "Yes")
	} else {
		Fprint(out, "No")
	}
}

//func main() { CF1119C(os.Stdin, os.Stdout) }
