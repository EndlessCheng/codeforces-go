package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://www.luogu.com.cn/problem/P4552

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, x, y int
	Fscan(in, &n)
	d := make([]int, n) // 32 位下建议用 int64
	for i := range d {
		Fscan(in, &d[i])
	}
	for i := n - 1; i > 0; i-- {
		d[i] -= d[i-1]
	}
	for _, v := range d[1:] {
		if v > 0 {
			x += v
		} else {
			y -= v
		}
	}
	Fprintln(out, min(x, y)+abs(x-y))
	Fprintln(out, abs(x-y)+1)
}

func main() { run(os.Stdin, os.Stdout) }

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
