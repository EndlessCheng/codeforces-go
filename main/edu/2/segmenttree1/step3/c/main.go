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
	pos := make([]int, n+1)
	ans := make([]interface{}, n)
	for i := 1; i <= m; i++ {
		Fscan(in, &v)
		p := pos[v]
		if p == 0 {
			pos[v] = i
			continue
		}
		ans[v-1] = query(p+1, i-1)
		add(p)
	}
	Fprint(out, ans...)
}

func main() { run(os.Stdin, os.Stdout) }
