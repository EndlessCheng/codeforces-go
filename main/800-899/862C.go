package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF862C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, x int
	Fscan(bufio.NewReader(in), &n, &x)

	var a []int
	f := func(m int) {
		for i := 0; i < m-1; i++ {
			v := 1 << (18 - i) // 一开始写成了 19，WA 了一发！
			x ^= v
			a = append(a, v)
		}
		a = append(a, x)
		n -= m
	}
	switch n % 4 {
	case 0:
		f(4)
	case 2:
		if n == 2 {
			if x == 0 {
				Fprint(out, "NO")
			} else {
				Fprint(out, "YES\n0 ", x)
			}
			return
		}
		f(6)
	default:
		f(n % 4)
	}

	Fprintln(out, "YES")
	has := map[int]bool{}
	for _, v := range a {
		has[v] = true
		Fprint(out, v, " ")
	}
	for i := 0; n > 0; i += 4 {
		if has[i] || has[i+1] || has[i+2] || has[i+3] {
			continue
		}
		Fprint(out, i, i+1, i+2, i+3, " ")
		n -= 4
	}
}

//func main() { CF862C(os.Stdin, os.Stdout) }
