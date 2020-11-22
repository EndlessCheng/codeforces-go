package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, t, ans, s int
	Fscan(in, &n, &t)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	if n == 1 {
		if a[0] > t {
			Fprint(out, 0)
		} else {
			Fprint(out, a[0])
		}
		return
	}

	b, end := []int{}, n/2
	var f func(int)
	f = func(p int) {
		if p == end {
			if s <= t {
				b = append(b, s)
			}
			return
		}
		f(p + 1)
		s += a[p]
		f(p + 1)
		s -= a[p]
	}
	f(0)
	l := b
	sort.Ints(l)
	b, end = nil, n
	f(n / 2)
	for _, v := range b {
		p := sort.SearchInts(l, t-v+1) - 1
		if l[p]+v > ans {
			ans = l[p] + v
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
