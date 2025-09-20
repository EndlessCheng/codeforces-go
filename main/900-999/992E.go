package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type fenwick92 []int

func (t fenwick92) update(i, x int) {
	for ; i < len(t); i += i & -i {
		t[i] += x
	}
}

func (t fenwick92) pre(i int) (s int) {
	for ; i > 0; i &= i - 1 {
		s += t[i]
	}
	return
}

func (t fenwick92) lowerBound(sum int) (res int) {
	for b := 1 << 17; b > 0; b >>= 1 {
		if nxt := res | b; nxt < len(t) && t[nxt] < sum {
			sum -= t[nxt]
			res = nxt
		}
	}
	return res + 1
}

func cf992E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, p, x int
	Fscan(in, &n, &q)
	a := make([]int, n+1)
	t := make(fenwick92, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		t.update(i, a[i])
	}

o:
	for range q {
		Fscan(in, &p, &x)
		t.update(p, x-a[p])
		a[p] = x

		i := 1
		for i <= n {
			s := t.pre(i)
			if a[i]*2 == s {
				Fprintln(out, i)
				continue o
			}
			i = t.lowerBound(s * 2)
		}
		Fprintln(out, -1)
	}
}

//func main() { cf992E(bufio.NewReader(os.Stdin), os.Stdout) }
