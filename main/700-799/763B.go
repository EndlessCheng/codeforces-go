package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF763B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	Fprintln(out, "YES")
	var n, x1, y1, x2, y2 int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &x1, &y1, &x2, &y2)
		// 保证 (x1,y1) 为左上角
		// 由于边长均为奇数，相邻矩形的左上角的横纵坐标的奇偶性必定不一样，那么按照奇偶性染色即可
		Fprintln(out, x1&1*2+y1&1+1)
	}
}

//func main() { CF763B(os.Stdin, os.Stdout) }
