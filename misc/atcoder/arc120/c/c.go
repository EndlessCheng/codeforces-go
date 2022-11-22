package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func f(a, b []int) (res int) {
	tree := make([]int, len(a)+1)
	add := func(i int) {
		for i++; i < len(tree); i += i & -i {
			tree[i]++
		}
	}
	sum := func(i int) (res int) {
		for ; i > 0; i &= i - 1 {
			res += tree[i]
		}
		return
	}

	pos := map[int][]int{}
	for i, v := range a {
		pos[v] = append(pos[v], i)
	}
	for i, v := range b {
		p := pos[v]
		if len(p) == 0 {
			return -1
		}
		j := p[0]
		pos[v] = p[1:]
		res += i - sum(j)
		add(j)
	}
	return
}

func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		a[i] += i
	}
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
		b[i] += i
	}
	Fprint(out, f(a, b))
}

func main() { run(os.Stdin, os.Stdout) }
