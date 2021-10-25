package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1512E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, l, r, s int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &l, &r, &s)
		l--
		k := r - l
		mi := (1 + k) * k / 2
		mx := (2*n - k + 1) * k / 2
		if s < mi || s > mx {
			Fprintln(out, -1)
			continue
		}

		a := make([]int, n)
		for i := l; i < r; i++ {
			a[i] = i - l + 1
		}
		cur := mi
		for i := r - 1; i >= l && cur < s; i-- {
			d := r - 1 - i
			mx := n - d
			if cur+mx-a[i] >= s {
				a[i] += s - cur
				break
			}
			cur += mx - a[i]
			a[i] = mx
		}

		vis := make([]bool, n+1)
		for i := l; i < r; i++ {
			vis[a[i]] = true
		}
		cur = 1
		for i, v := range a {
			if v == 0 {
				for ; vis[cur]; cur++ {
				}
				a[i] = cur
				cur++
			}
		}
		for _, v := range a {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1512E(os.Stdin, os.Stdout) }
