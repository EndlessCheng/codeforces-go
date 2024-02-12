package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p1823(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	type pair struct{ h, c int }
	st := []pair{}
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && a[i] > st[len(st)-1].h {
			ans += st[len(st)-1].c
			st = st[:len(st)-1]
		}
		if len(st) > 0 && a[i] == st[len(st)-1].h {
			ans += st[len(st)-1].c
			if len(st) > 1 {
				ans++
			}
			st[len(st)-1].c++
		} else {
			if len(st) > 0 {
				ans++
			}
			st = append(st, pair{a[i], 1})
		}
	}
	Fprint(out, ans)
}

//func main() { p1823(os.Stdin, os.Stdout) }
