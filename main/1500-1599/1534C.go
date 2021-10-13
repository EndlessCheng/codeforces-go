package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1534C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		pos := make([]int, n+1)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			pos[v] = i
		}
		b := make([]int, n)
		for i := range b {
			Fscan(in, &b[i])
		}
		ans := 1
		vis := make([]bool, n)
		for i, s := range vis {
			if !s {
				ans = ans * 2 % (1e9 + 7)
				for ; !vis[i]; i = pos[b[i]] {
					vis[i] = true
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1534C(os.Stdin, os.Stdout) }
