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

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		vis := map[int]bool{}
		ans := []int{}
		for i := n - 1; i >= 0; i-- {
			if v := a[i]; !vis[v] {
				vis[v] = true
				ans = append(ans, v)
			}
		}
		Fprintln(out, len(ans))
		for i := len(ans) - 1; i >= 0; i-- {
			Fprint(out, ans[i], " ")
		}
		Fprintln(out)
	}
}

func main() { run(os.Stdin, os.Stdout) }
