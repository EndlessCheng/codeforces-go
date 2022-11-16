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
	const inf int = 1e18
	var n, m, c int
	Fscan(in, &n, &m, &c)
	pre := make([]int, m+1)
	suf := make([]int, m+1)
	for i := range pre {
		pre[i] = inf
		suf[i] = inf
	}
	ans := inf
	a := make([]int, m)
	for i := 0; i < n; i++ {
		for j := range a {
			Fscan(in, &a[j])
			ans = min(ans, a[j]+c*(i+j)+min(pre[j+1], pre[j]))
			pre[j+1] = min(min(pre[j+1], pre[j]), a[j]-c*(i+j))
		}
		for j := m - 1; j >= 0; j-- {
			ans = min(ans, a[j]+c*(i-j)+suf[j])
			suf[j] = min(min(suf[j], suf[j+1]), a[j]-c*(i-j))
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int { if a > b { return b }; return a }
