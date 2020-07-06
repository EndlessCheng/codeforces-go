package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1000F(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	buf := make([]byte, 4096)
	_i := len(buf)
	rc := func() byte {
		if _i == len(buf) {
			_r.Read(buf)
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	ri := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}
	type query struct{ bid, l, r, idx int }

	n := ri()
	a := make([]int, n)
	for i := range a {
		a[i] = ri()
	}
	q := ri()
	qs := make([]query, q)
	blockSize := int(math.Round(math.Sqrt(float64(n))))
	for i := range qs {
		l := ri()
		qs[i] = query{l / blockSize, l, ri() + 1, i}
	}
	sort.Slice(qs, func(i, j int) bool {
		qi, qj := qs[i], qs[j]
		if qi.bid != qj.bid {
			return qi.bid < qj.bid
		}
		if qi.bid&1 == 0 {
			return qi.r < qj.r
		}
		return qi.r > qj.r
	})

	cnt := [5e5 + 1]int{}
	del := [5e5 + 1]int{} // 懒删除标记
	s := []int{0}
	update := func(i, d int) {
		v := a[i-1]
		if cnt[v] == 1 {
			del[v]++ // 从栈中删除 v
		}
		cnt[v] += d
		if cnt[v] == 1 {
			// 把 v 插入栈
			if del[v] > 0 {
				del[v]--
			} else {
				s = append(s, v)
			}
		}
	}
	ans := make([]int, q)
	l, r := 1, 1
	for _, q := range qs {
		for ; r < q.r; r++ {
			update(r, 1)
		}
		for ; l < q.l; l++ {
			update(l, -1)
		}
		for l > q.l {
			l--
			update(l, 1)
		}
		for r > q.r {
			r--
			update(r, -1)
		}
		// 获取栈顶元素
		for del[s[len(s)-1]] > 0 {
			del[s[len(s)-1]]--
			s = s[:len(s)-1]
		}
		ans[q.idx] = s[len(s)-1]
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { CF1000F(os.Stdin, os.Stdout) }
