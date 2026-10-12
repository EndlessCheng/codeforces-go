package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func p9588(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var m, op, x, pop int
	Fscan(in, &m, &m)
	a := []int{}
	s := []int{0}
	q := []int{}

	for range m {
		Fscan(in, &op)
		if op == 4 {
			for s[q[0]+1] <= pop {
				q = q[1:]
			}
			Fprintln(out, a[q[0]])
			continue
		}
		Fscan(in, &x)
		if op == 1 {
			for len(q) > 0 && a[q[len(q)-1]] <= x {
				q = q[:len(q)-1]
			}
			q = append(q, len(a))
			a = append(a, x)
			s = append(s, s[len(s)-1]+x)
		} else if op == 2 {
			pop += x
		} else {
			x += pop
			i := sort.SearchInts(s, x) - 1
			Fprintln(out, x-s[i])
		}
	}
}

//func main() { p9588(bufio.NewReader(os.Stdin), os.Stdout) }
