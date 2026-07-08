package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
type fenwick []int

func (t fenwick) update(i, v int) {
	for ; i < len(t); i += i & -i {
		t[i] += v
	}
}

func (t fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

func run(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	f := func(k int) (res int) {
		t := make(fenwick, n+1)
		inv, l := 0, 0
		for i, v := range a {
			inv += t.pre(n) - t.pre(v)
			t.update(v, 1)
			for l <= i && inv >= k {
				t.update(a[l], -1)
				inv -= t.pre(a[l])
				l++
			}
			res += l
		}
		return
	}
	Fprint(out, f(k)-f(k+1))
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
