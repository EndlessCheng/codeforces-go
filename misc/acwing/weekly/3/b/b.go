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

	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		pos := make([][]int, 100)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			pos[v-1] = append(pos[v-1], i)
		}
		ans := n
		for _, p := range pos {
			if p == nil {
				continue
			}
			res := 0
			p = append(append([]int{-1}, p...), n)
			for i, l := 1, 0; i < len(p); i++ {
				if p[i] <= l || p[i-1]+1 == p[i] {
					continue
				}
				l = max(l, p[i-1]+1)
				c := (p[i]-l-1)/k + 1
				res += c
				l += c * k
			}
			ans = min(ans, res)
		}
		Fprintln(out, ans)
	}
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
