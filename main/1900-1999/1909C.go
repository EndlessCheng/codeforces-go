package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1909C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		l := make([]int, n)
		for i := range l {
			Fscan(in, &l[i])
		}
		sort.Ints(l)
		r := make([]int, n)
		for i := range r {
			Fscan(in, &r[i])
		}
		sort.Ints(r)

		st := l[:0]
		j := 0
		for i, v := range r {
			for j < n && l[j] < v {
				st = append(st, l[j])
				j++
			}
			r[i] -= st[len(st)-1]
			st = st[:len(st)-1]
		}
		sort.Ints(r)

		c := l
		for i := range c {
			Fscan(in, &c[i])
		}
		sort.Ints(c)

		ans := 0
		for i, v := range c {
			ans += v * r[n-1-i]
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1909C(os.Stdin, os.Stdout) }
