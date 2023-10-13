package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1792C(_r io.Reader, _w io.Writer) {
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
		x := n / 2
		for x > 0 && pos[x] < pos[x+1] && pos[n-x] < pos[n-x+1] {
			x--
		}
		Fprintln(out, x)
	}
}

//func main() { CF1792C(os.Stdin, os.Stdout) }
