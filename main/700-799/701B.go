package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF701B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, x, y int
	Fscan(in, &n, &m)
	rows, cols := map[int]bool{}, map[int]bool{}
	for ; m > 0; m-- {
		Fscan(in, &x, &y)
		rows[x] = true
		cols[y] = true
		Fprint(out, int64(n)*int64(n-len(rows)-len(cols))+int64(len(rows))*int64(len(cols)), " ")
	}
}

//func main() { CF701B(os.Stdin, os.Stdout) }
