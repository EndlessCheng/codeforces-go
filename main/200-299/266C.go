package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"sort"
)

// 另一种做法是通过交换空列到第 n 列，以及交换一个含有 1 的行到第 n 行，从而归纳到 n-1 方阵的情况

// github.com/EndlessCheng/codeforces-go
type data struct {
	a     [][]byte
	tp    int
	swaps *[][3]int
}

func (d data) Len() int { return len(d.a) }
func (d data) Less(i, j int) bool {
	if d.tp == 1 {
		return bytes.Count(d.a[i], []byte{1}) < bytes.Count(d.a[j], []byte{1})
	}
	for _, r := range d.a {
		if r[i] > 0 || r[j] > 0 {
			return r[i] > r[j]
		}
	}
	return false
}
func (d data) Swap(i, j int) {
	if i == j {
		return
	}
	*d.swaps = append(*d.swaps, [3]int{d.tp, i, j})
	if d.tp == 1 {
		d.a[i], d.a[j] = d.a[j], d.a[i]
	} else {
		for _, r := range d.a {
			r[i], r[j] = r[j], r[i]
		}
	}
}

func CF266C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x, y int
	Fscan(in, &n)
	d := data{make([][]byte, n), 1, &[][3]int{}}
	for i := range d.a {
		d.a[i] = make([]byte, n)
	}
	for i := 1; i < n; i++ {
		Fscan(in, &x, &y)
		d.a[x-1][y-1] = 1
	}
	sort.Sort(d)
	d.tp = 2
	sort.Sort(d)
	a := *d.swaps
	Fprintln(out, len(a))
	for _, p := range a {
		Fprintln(out, p[0], p[1]+1, p[2]+1)
	}
}

//func main() { CF266C(os.Stdin, os.Stdout) }
