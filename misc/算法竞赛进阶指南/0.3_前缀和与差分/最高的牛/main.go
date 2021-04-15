package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://www.luogu.com.cn/problem/P2879

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, p, h, m, a, b int
	Fscan(in, &n, &p, &h, &m)
	d := make([]int, n+1)
	use := map[[2]int]bool{}
	for ; m > 0; m-- {
		Fscan(in, &a, &b)
		if a > b {
			a, b = b, a
		}
		p := [2]int{a, b}
		if use[p] {
			continue
		}
		use[p] = true
		d[a+1]--
		d[b]++
	}
	for _, d := range d[1:] {
		h += d
		Fprintln(out, h)
	}
}

func main() { run(os.Stdin, os.Stdout) }
