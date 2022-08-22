package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1016D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, xorA, xorB int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		xorA ^= a[i]
	}
	b := make([]int, m)
	for i := range b {
		Fscan(in, &b[i])
		xorB ^= b[i]
	}

	if xorA != xorB {
		Fprint(out, "NO")
		return
	}
	Fprintln(out, "YES")
	Fprint(out, xorA^a[0]^b[0]) // 左上角
	for _, x := range b[1:] {
		Fprint(out, " ", x)
	}
	Fprintln(out)
	zeros := strings.Repeat("0 ", m-1)
	for _, x := range a[1:] {
		Fprintln(out, x, zeros)
	}
}

//func main() { CF1016D(os.Stdin, os.Stdout) }
