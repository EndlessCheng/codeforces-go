package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func CF388B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var k uint
	Fscan(in, &k)
	if k == 1 {
		Fprint(out, "2\nNY\nYN")
		return
	}
	m := bits.Len(k)
	n := m * 5
	g := make([][]byte, n)
	for i := range g {
		g[i] = bytes.Repeat([]byte{'N'}, n)
	}
	add := func(v, w int) {
		g[v][w] = 'Y'
		g[w][v] = 'Y'
	}
	add(0, 2)
	add(0, 3)
	if k&1 > 0 {
		add(3, 4)
	}
	for i := 1; i < m; i++ {
		v := i * 5
		if i == 1 {
			add(v, v-2)
		} else {
			add(v, v-4)
		}
		add(v, v-3)
		add(v, v+1)
		add(v, v+2)
		add(v+3, v-1)
		add(v+3, v+4)
		if k>>i&1 > 0 {
			add(v+2, v+4)
		}
	}
	add(1, m*5-1)
	Fprintln(out, n)
	for _, row := range g {
		Fprintf(out, "%s\n", row)
	}
}

//func main() { CF388B(os.Stdin, os.Stdout) }
