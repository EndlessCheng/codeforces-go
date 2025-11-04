package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf2168B(in io.Reader, out io.Writer) {
	var tp string
	var T, n, v int
	Fscan(in, &tp, &T)
	if tp[0] == 'f' {
		for range T {
			Fscan(in, &n)
			pos := make([]int, n+1)
			for i := range n {
				Fscan(in, &v)
				pos[v] = i
			}
			if pos[1] < pos[n] {
				Fprintln(out, 0)
			} else {
				Fprintln(out, 1)
			}
		}
	} else {
		q := func(l, r int) int {
			Println("?", l, r)
			Fscan(in, &r)
			return r
		}
		for range T {
			Fscan(in, &n, &v)
			if v == 0 {
				Fprintln(out, "!", sort.Search(n, func(i int) bool { return q(1, i+1) == n-1 })+1)
			} else {
				Fprintln(out, "!", sort.Search(n, func(i int) bool { return q(i+1, n) < n-1 }))
			}
		}
	}
}

//func main() { cf2168B(bufio.NewReader(os.Stdin), os.Stdout) }
