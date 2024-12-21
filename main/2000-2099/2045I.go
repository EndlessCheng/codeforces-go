package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type fenwick45 []int

func (f fenwick45) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

func (f fenwick45) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func (f fenwick45) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}

func cf2045I(in io.Reader, out io.Writer) {
	var n, m, v, ans int
	Fscan(in, &n, &m)
	pre := make([]int, m)
	f := make(fenwick45, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		v--
		if pre[v] == 0 {
			ans += m - 1 // v 的贡献
		} else {
			ans += f.query(pre[v]+1, i) // v 的贡献
			f.update(pre[v], -1)
		}
		f.update(i, 1)
		pre[v] = i
	}
	Fprint(out, ans)
}

//func main() { cf2045I(bufio.NewReader(os.Stdin), os.Stdout) }
