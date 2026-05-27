package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	a := make([]struct{ w, v int }, n)
	pre := make([][]int, n+1)
	for i := range pre {
		pre[i] = make([]int, m+1)
	}
	for i := range a {
		Fscan(in, &a[i].w, &a[i].v)
		p := a[i]
		copy(pre[i+1], pre[i][:p.w])
		for j := p.w; j <= m; j++ {
			pre[i+1][j] = max(pre[i][j], pre[i][j-p.w]+p.v)
		}
	}

	maxS := pre[n][m]
	ans := make([]byte, n)
	suf := make([]int, m+1)
	for i := n - 1; i >= 0; i-- {
		p := a[i]

		notChoose, choose := 0, 0
		for j, v := range pre[i] {
			notChoose = max(notChoose, v+suf[m-j])
			if j <= m-p.w {
				choose = max(choose, v+p.v+suf[m-p.w-j])
			}
		}
		if notChoose < maxS {
			ans[i] = 'A'
		} else if choose < maxS {
			ans[i] = 'C'
		} else {
			ans[i] = 'B'
		}

		for j := m; j >= p.w; j-- {
			suf[j] = max(suf[j], suf[j-p.w]+p.v)
		}
	}
	Fprintf(out, "%s", ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
