package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1736C1(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			a[i] -= i + 1
		}

		ans := 0
		type pair struct{ v, r int }
		q := []pair{}
		for i := n - 1; i >= 0; i-- {
			r := i
			for len(q) > 0 && a[i] <= q[len(q)-1].v {
				r = q[len(q)-1].r
				q = q[:len(q)-1]
			}
			q = append(q, pair{a[i], r})
			if q[0].v < -i {
				q = q[1:]
			}
			ans += q[0].r - i + 1
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1736C1(bufio.NewReader(os.Stdin), os.Stdout) }
