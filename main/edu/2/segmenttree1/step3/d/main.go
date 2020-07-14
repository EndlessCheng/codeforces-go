package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v int
	Fscan(in, &n)
	m := 2 * n
	tree := make([]int, m+1)
	add := func(i int) {
		for ; i <= m; i += i & -i {
			tree[i]++
		}
	}
	sub := func(i int) {
		for ; i <= m; i += i & -i {
			tree[i]--
		}
	}
	query := func(l, r int) (s int) {
		l--
		for ; r > l; r &= r - 1 {
			s += tree[r]
		}
		for ; l > r; l &= l - 1 {
			s -= tree[l]
		}
		return
	}
	ans := make([]int, n+1)
	a := make([]int, m+1)
	pos := make([]int, n+1)
	for i := 1; i <= m; i++ {
		Fscan(in, &v)
		a[i] = v
		p := pos[v]
		if p == 0 {
			pos[v] = i
			add(i)
		} else {
			ans[v] = query(p+1, i-1)
			sub(p)
		}
	}
	tree = make([]int, m+1)
	pos = make([]int, n+1)
	for i := m; i > 0; i-- {
		v := a[i]
		p := pos[v]
		if p == 0 {
			pos[v] = i
			add(i)
		} else {
			ans[v] += query(i+1, p-1)
			sub(p)
		}
	}
	for _, v := range ans[1:] {
		Fprint(out, v, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
