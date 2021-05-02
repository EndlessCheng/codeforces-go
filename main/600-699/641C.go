package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF641C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	p, q := int64(0), int64(1)
	var n, m, op, x int64
	for Fscan(in, &n, &m); m > 0; m-- {
		if Fscan(in, &op); op == 1 {
			Fscan(in, &x)
			p += x
			q += x
		} else {
			p ^= 1
			q ^= 1
		}
	}
	ans := make([]interface{}, n)
	p = p%n + n
	for i := int64(0); i < n; i += 2 {
		ans[(p+i)%n] = i + 1
	}
	q = q%n + n
	for i := int64(0); i < n; i += 2 {
		ans[(q+i)%n] = i + 2
	}
	Fprintln(out, ans...)
}

//func main() { CF641C(os.Stdin, os.Stdout) }
