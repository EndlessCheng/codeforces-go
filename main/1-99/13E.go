package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF13E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, p, v, last int32
	Fscan(in, &n, &q)
	a := make([]int32, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	const B = 500
	type pair struct{ last, jump int32 }
	data := make([]pair, n)
	f := func(i int32) {
		if i+a[i] >= n || i/B != (i+a[i])/B {
			data[i] = pair{i, 0}
		} else {
			data[i] = data[i+a[i]]
			data[i].jump++
		}
	}
	for i := n - 1; i >= 0; i-- {
		f(i)
	}

	for ; q > 0; q-- {
		Fscan(in, &op, &p)
		p--
		if op == 0 {
			Fscan(in, &v)
			a[p] = v
			for i := p; i >= p-p%B; i-- {
				f(i)
			}
		} else {
			jump := int32(0)
			for ; p < n; p = last + a[last] {
				jump += data[p].jump + 1
				last = data[p].last
			}
			Fprintln(out, last+1, jump)
		}
	}
}

//func main() { CF13E(os.Stdin, os.Stdout) }
