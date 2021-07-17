package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, ans int
	Fscan(in, &n)
	c := map[int]int{}
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		c[v-i] += v
		ans = max(ans, c[v-i])
	}
	Fprint(out, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { run(os.Stdin, os.Stdout) }
