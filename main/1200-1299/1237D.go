package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1237D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	a := make([]int, n, n*3)
	for i := range a {
		Fscan(in, &a[i])
	}
	a = append(append(a, a...), a...)
	q := []int{}
	for l, r := 0, 0; l < n; l++ {
		for ; r < n*3 && (len(q) == 0 || a[r]*2 >= a[q[0]]); r++ {
			for len(q) > 0 && a[q[len(q)-1]] <= a[r] {
				q = q[:len(q)-1]
			}
			q = append(q, r)
		}
		if r == n*3 {
			Fprint(out, "-1 ")
		} else {
			Fprint(out, r-l, " ")
		}
		if q[0] == l {
			q = q[1:]
		}
	}
}

//func main() { cf1237D(bufio.NewReader(os.Stdin), os.Stdout) }
