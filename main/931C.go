package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF931C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	c := map[int]int{}
	for i := range a {
		Fscan(in, &a[i])
		c[a[i]]++
	}
	b := []int{}
	for k := range c {
		b = append(b, k)
	}
	sort.Ints(b)
	if len(b) == 2 && b[0]+2 == b[1] {
		b = []int{b[0], b[0] + 1, b[1]}
	}
	if len(b) == 3 {
		min02 := c[b[0]]
		if c[b[2]] < min02 {
			min02 = c[b[2]]
		}
		if c[b[1]]/2 > min02 {
			c[b[1]] &^= 1
			n -= c[b[1]]
			for i, v := range a {
				if v == b[1] && c[v] > 0 {
					a[i] = b[c[v]&1<<1]
					c[v]--
				}
			}
		} else {
			n -= 2 * min02
			c[b[0]] = min02
			c[b[2]] = min02
			for i, v := range a {
				if v != b[1] && c[v] > 0 {
					a[i] = b[1]
					c[v]--
				}
			}
		}
	}
	Fprintln(out, n)
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

//func main() { CF931C(os.Stdin, os.Stdout) }
