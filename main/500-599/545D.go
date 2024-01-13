package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf545D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, s, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	for _, t := range a {
		if s <= t {
			s += t
			ans++
		}
	}
	Fprint(out, ans)
}

//func main() { cf545D(os.Stdin, os.Stdout) }
