package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF459C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, b, d int
	Fscan(in, &n, &b, &d)
	ans := [][]int{}
	a := make([]int, d)
	var f func(int)
	f = func(p int) {
		if p == d {
			ans = append(ans, append([]int(nil), a...))
			return
		}
		for a[p] = 1; a[p] <= b && len(ans) < n; a[p]++ {
			f(p + 1)
		}
	}
	f(0)
	if len(ans) < n {
		Fprint(out, -1)
		return
	}
	for j := 0; j < d; j++ {
		for _, row := range ans {
			Fprint(out, row[j], " ")
		}
		Fprintln(out)
	}
}

//func main() { CF459C(os.Stdin, os.Stdout) }
