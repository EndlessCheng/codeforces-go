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

	var T, n, m, r, c int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &r, &c)
		ans := 0
		ans = max(ans, abs(1-r)+abs(1-c))
		ans = max(ans, abs(n-r)+abs(1-c))
		ans = max(ans, abs(1-r)+abs(m-c))
		ans = max(ans, abs(n-r)+abs(m-c))
		Fprintln(out, ans)
	}
}

func main() { run(os.Stdin, os.Stdout) }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
