package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1675D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		p := make([]int, n)
		vis := make([]bool, n)
		for i := range p {
			Fscan(in, &p[i])
			p[i]--
			if p[i] != i {
				vis[p[i]] = true
			}
		}

		ls := []int{}
		for v, b := range vis {
			if !b {
				ls = append(ls, v)
			}
		}
		Fprintln(out, len(ls))
		for _, v := range ls {
			path := []int{v}
			for v = p[v]; vis[v]; v = p[v] {
				path = append(path, v)
				vis[v] = false
				if p[v] == v {
					break
				}
			}
			Fprintln(out, len(path))
			for i := len(path) - 1; i >= 0; i-- {
				Fprint(out, path[i]+1, " ")
			}
			Fprintln(out)
		}
		Fprintln(out)
	}
}

//func main() { CF1675D(os.Stdin, os.Stdout) }
