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
	var n, m, ans int
	Fscan(in, &n, &m)
	a := make([]int, n)
	cnt := make([]int, n+1)
	for i := range a {
		Fscan(in, &a[i])
		if i < m {
			cnt[a[i]]++
		}
	}
	for cnt[ans] > 0 {
		ans++
	}
	for i, v := range a[m:] {
		cnt[v]++
		w := a[i] // 离开窗口的数
		cnt[w]--
		if cnt[w] == 0 && w < ans {
			ans = w
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
