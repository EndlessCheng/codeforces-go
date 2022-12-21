package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const inf int = 1e9

	var n, m, v int
	Fscan(in, &n, &m)
	f := make([][2]int, m+1)
	for i := range f {
		f[i] = [2]int{inf, inf}
	}
	f[0][1] = 0
	for ; n > 0; n-- {
		Fscan(in, &v)
		for j := m; j > 0; j-- {
			f[j][0] = min(f[j][0], f[j][1]+1)
			if j >= v {
				f[j][1] = min(f[j-v][0], f[j-v][1])
			} else {
				f[j][1] = inf
			}
		}
		f[0] = [2]int{1, inf}
	}
	for _, p := range f[1:] {
		res := min(p[0], p[1])
		if res == inf {
			res = -1
		}
		Fprintln(out, res)
	}
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int { if b < a { return b }; return a }
