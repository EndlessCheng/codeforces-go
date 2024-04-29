package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func CF1365C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v int
	Fscan(in, &n)
	cnt := make([]int, n)
	pos := make([]int, n+1)
	for i := range cnt {
		Fscan(in, &v)
		pos[v] = i
	}
	for i := range cnt {
		Fscan(in, &v)
		p := pos[v]
		if p < i {
			p += n
		}
		cnt[p-i]++
	}
	Fprint(out, slices.Max(cnt))
}

//func main() { CF1365C(os.Stdin, os.Stdout) }
