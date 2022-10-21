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

	var n, mi int
	Fscan(in, &n)
	s := make([]int, n)
	for i := range s {
		Fscan(in, &s[i])
	}
	t := make([]int, n)
	for i := range t {
		Fscan(in, &t[i])
		if t[i] < t[mi] {
			mi = i // 寻找 mi，使 dp 满足无后效性
		}
	}

	ans := make([]int, n)
	ans[mi] = t[mi]
	for i := 1; i < n; i++ {
		cur := (mi + i) % n
		pre := (cur + n - 1) % n
		ans[cur] = min(t[cur], ans[pre]+s[pre])
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
