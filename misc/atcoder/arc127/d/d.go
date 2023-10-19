package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
	}

	f := func(a, p, q []int) (res int) {
		if p == nil || q == nil {
			return
		}
		for j := 0; j < 18; j++ {
			p1, q1 := 0, 0
			for _, i := range p {
				p1 += a[i] >> j & 1
			}
			for _, i := range q {
				q1 += a[i] >> j & 1
			}
			res += (p1*(len(q)-q1) + (len(p)-p1)*q1) << j
		}
		return
	}

	var dfs func([]int, int)
	dfs = func(idx []int, j int) {
		if len(idx) <= 1 {
			return
		}
		if j < 0 {
			ans += f(a, idx, idx) / 2
			return
		}
		p := [4][]int{}
		for _, i := range idx {
			m := a[i]>>j&1<<1 | b[i]>>j&1
			p[m] = append(p[m], i)
		}
		ans += f(a, p[0], p[1])
		ans += f(a, p[2], p[3])
		ans += f(b, p[0], p[2])
		ans += f(b, p[1], p[3])
		dfs(append(p[0], p[3]...), j-1)
		dfs(append(p[1], p[2]...), j-1)
	}

	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	dfs(idx, 17)
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
