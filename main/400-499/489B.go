package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf489B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	Fscan(in, &m)
	b := make([]int, m)
	for i := range b {
		Fscan(in, &b[i])
	}
	sort.Ints(b)

	j := 0
	for _, v := range a {
		for j < m && b[j] < v-1 {
			j++
		}
		if j < m && b[j] <= v+1 {
			ans++
			j++
		}
	}
	Fprint(out, ans)
}

//func main() { cf489B(os.Stdin, os.Stdout) }
