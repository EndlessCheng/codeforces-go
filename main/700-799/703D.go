package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type fenwick03 []int

func (f fenwick03) update(i, v int) {
	for ; i < len(f); i += i & -i {
		f[i] ^= v
	}
}

func (f fenwick03) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res ^= f[i]
	}
	return
}

func cf703D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, l, r int
	Fscan(in, &n)
	a := make([]int, n+1)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		s[i] = s[i-1] ^ a[i]
	}
	Fscan(in, &m)
	type pair struct{ l, i int }
	qs := make([][]pair, n+1)
	for i := range m {
		Fscan(in, &l, &r)
		qs[r] = append(qs[r], pair{l, i})
	}

	ans := make([]int, m)
	t := make(fenwick03, n+1)
	last := map[int]int{}
	for i := 1; i <= n; i++ {
		v := a[i]
		if j, ok := last[v]; ok {
			t.update(j, v)
		}
		last[v] = i
		t.update(i, v)
		for _, p := range qs[i] {
			ans[p.i] = s[i] ^ s[p.l-1] ^ t.pre(i) ^ t.pre(p.l-1)
		}
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { cf703D(bufio.NewReader(os.Stdin), os.Stdout) }
