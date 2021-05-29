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
			a[i]--
		}
		vis := make([]bool, n)
		ans := make([]interface{}, n)
		for v, b := range vis {
			if !b {
				vs := []int{}
				for !vis[v] {
					vis[v] = true
					vs = append(vs, v)
					v = a[v]
				}
				for _, v := range vs {
					ans[v] = len(vs)
				}
			}
		}
		Fprintln(out, ans...)
	}
}

func main() { run(os.Stdin, os.Stdout) }
