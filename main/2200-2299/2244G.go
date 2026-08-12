package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type fenwick []int

func (t fenwick) update(i, v int) {
	for ; i < len(t); i += i & -i {
		t[i] = max(t[i], v)
	}
}

func (t fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res = max(res, t[i])
	}
	return
}

func cf2244G(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		t := make(fenwick, n+1)
		type pair struct{ i, f int }
		todo := make([][]pair, n+1)
		ans := 0
		for i := 1; i <= n; i++ {
			for _, p := range todo[i] {
				t.update(p.i, p.f)
			}
			Fscan(in, &v)
			res := t.pre(i-v-1) + v
			ans = max(ans, res)
			if s := i + v + 1; s <= n {
				todo[s] = append(todo[s], pair{i, res})
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2244G(bufio.NewReader(os.Stdin), os.Stdout) }
