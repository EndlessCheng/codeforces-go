package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, a, b, i int
	Fscan(in, &n, &a, &b)
	if a+b-1 > n || a*b < n {
		Fprint(out, -1)
		return
	}
	for a--; a >= 0; a-- {
		sz := min(b, n-a) // 后面至少要留下 a 个，这样才能满足 |LIS|=a 的条件
		for j := i + sz; j > i; j-- {
			Fprint(out, j, " ")
		}
		i += sz
		n -= sz
	}
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int { if a > b { return b }; return a }
