package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf696B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	p := make([]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &p[i])
		p[i]--
	}

	size := make([]int, n)
	for i := n - 1; i > 0; i-- {
		size[i]++
		size[p[i]] += size[i]
	}
	size[0]++

	depth := make([]int, n)
	for i, sz := range size {
		depth[i] = depth[p[i]] + 1
		Fprintf(out, "%.1f ", float64(n-sz+depth[i]+1)/2)
	}
}

//func main() { cf696B(os.Stdin, os.Stdout) }
