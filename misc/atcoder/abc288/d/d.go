package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, q, l, r int
	Fscan(in, &n, &k)
	a := make([]int, n)
	d := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if i > 0 {
			d[i] = a[i] - a[i-1]
		}
	}

	// 同余前缀和模板
	for len(d)%k > 0 {
		d = append(d, 0)
	}
	sum := make([]int, len(d)+k)
	for i, v := range d {
		sum[i+k] = sum[i] + v
	}
	pre := func(x, t int) int {
		if x%k <= t {
			return sum[x/k*k+t]
		}
		return sum[(x+k-1)/k*k+t]
	}
	query := func(l, r, t int) int {
		t %= k
		return pre(r, t) - pre(l, t)
	}

o:
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &l, &r)
		l--
		if (r-l)%k > 0 && a[l]+query(l+k, r, l) != 0 {
			Fprintln(out, "No")
			continue
		}
		for i := 1; i < k; i++ {
			l++
			if (r-l)%k > 0 && query(l, r, l) != 0 {
				Fprintln(out, "No")
				continue o
			}
		}
		Fprintln(out, "Yes")
	}
}

func main() { run(os.Stdin, os.Stdout) }
