package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf600B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	for ; m > 0; m-- {
		Fscan(in, &v)
		Fprint(out, sort.SearchInts(a, v+1), " ")
	}
}

//func main() { cf600B(os.Stdin, os.Stdout) }
