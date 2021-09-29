package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1573B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		pos := make([]int, n)
		for i := range pos {
			Fscan(in, &v)
			pos[v/2] = i
		}
		minP := make([]int, n)
		minP[0] = pos[0]
		for i := 1; i < n; i++ {
			minP[i] = min(minP[i-1], pos[i])
		}
		ans := int(1e9)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			ans = min(ans, i+minP[v/2-1])
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1573B(os.Stdin, os.Stdout) }
