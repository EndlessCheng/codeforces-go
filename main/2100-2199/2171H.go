package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type fenwick71 []int

func (t fenwick71) update(i, v int) {
	for ; i < len(t); i += i & -i {
		t[i] = max(t[i], v)
	}
}

func (t fenwick71) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res = max(res, t[i])
	}
	return
}

func cf2171H(in io.Reader, out io.Writer) {
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		t := make(fenwick71, m+1)
		for i := 2; i <= n; i++ {
			for v := m - m%i; v >= i; v -= i { // v 必须是 i 的倍数，贡献 e 才能是正的
				e := 0
				for x := v; x%i == 0; x /= i {
					e++
				}
				w := v - i + 1
				t.update(w, t.pre(w)+e)
			}
		}
		Fprintln(out, t.pre(m-n+1))
	}
}

//func main() { cf2171H(bufio.NewReader(os.Stdin), os.Stdout) }
