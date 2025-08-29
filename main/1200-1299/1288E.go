package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type fenwick88 []int

func (f fenwick88) update(i, val int) {
	for i++; i < len(f); i += i & -i {
		f[i] += val
	}
}

func (f fenwick88) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func cf1288E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, v int
	Fscan(in, &n, &m)
	mn := make([]int, n)
	mx := make([]int, n)
	pos := make([]int, n)
	t := make(fenwick88, m+n+1)
	for i := range mn {
		mn[i] = i
		pos[i] = m + i
		t.update(m+i, 1)
	}

	for i := m - 1; i >= 0; i-- {
		Fscan(in, &v)
		v--
		mn[v] = 0
		old := pos[v]
		pos[v] = i
		mx[v] = max(mx[v], t.pre(old)) // 在移动前达到最大
		t.update(old, -1)
		t.update(i, 1)
	}

	for i, v := range mn {
		Fprintln(out, v+1, max(mx[i], t.pre(pos[i]))+1)
	}
}

//func main() { cf1288E(bufio.NewReader(os.Stdin), os.Stdout) }
