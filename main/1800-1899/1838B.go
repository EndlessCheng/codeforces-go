package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1838B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		p := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &v)
			p[v] = i
		}
		b := []int{p[1], p[2], p[n]}
		sort.Ints(b)
		Fprintln(out, b[1], p[n])
	}
}

//func main() { CF1838B(os.Stdin, os.Stdout) }
