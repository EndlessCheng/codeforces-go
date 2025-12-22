package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p8773(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, x, v, l, r int
	Fscan(in, &n, &m, &x)

	pos := [1 << 20]int{}
	maxI := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		maxI[i] = max(maxI[i-1], pos[v^x])
		pos[v] = i
	}

	for i := 1; i <= m; i++ {
		Fscan(in, &l, &r)
		if maxI[r] >= l {
			Fprintln(out, "yes")
		} else {
			Fprintln(out, "no")
		}
	}
}

//func main() { p8773(bufio.NewReader(os.Stdin), os.Stdout) }
