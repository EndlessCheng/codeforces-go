package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1684C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		k := -1
		a := make([][]int, n)
		for i := range a {
			a[i] = make([]int, m)
			for j := range a[i] {
				Fscan(in, &a[i][j])
			}
			if !sort.IntsAreSorted(a[i]) {
				k = i
			}
		}
		if k < 0 {
			Fprintln(out, 1, 1)
			continue
		}
		b := append([]int{}, a[k]...)
		sort.Ints(b)
		ps := []int{}
		for j, v := range a[k] {
			if v != b[j] {
				if len(ps) == 2 {
					Fprintln(out, -1)
					continue o
				}
				ps = append(ps, j)
			}
		}
		i, j := ps[0], ps[1]
		for _, r := range a {
			r[i], r[j] = r[j], r[i]
			if !sort.IntsAreSorted(r) {
				Fprintln(out, -1)
				continue o
			}
		}
		Fprintln(out, i+1, j+1)
	}
}

// 我的憨憨写法（虽然是线性的，但是明显排序更优雅）
func CF1684C2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		k := -1
		a := make([][]int, n)
		for i := range a {
			a[i] = make([]int, m)
			for j := range a[i] {
				Fscan(in, &a[i][j])
			}
			if !sort.IntsAreSorted(a[i]) {
				k = i
			}
		}
		if k < 0 {
			Fprintln(out, 1, 1)
			continue
		}
		ps := []int{}
		for j := 1; j < m; j++ {
			if a[k][j-1] > a[k][j] {
				if len(ps) == 2 {
					Fprintln(out, -1)
					continue o
				}
				ps = append(ps, j)
			}
		}
		var i, j int
		if len(ps) == 1 {
			j = ps[0]
			i = sort.SearchInts(a[k][:j], a[k][j]+1)
			for _, r := range a {
				r[i], r[j] = r[j], r[i]
				if !sort.IntsAreSorted(r) {
					r[i], r[j] = r[j], r[i]
					goto next
				}
				r[i], r[j] = r[j], r[i]
			}
			Fprintln(out, i+1, j+1)
			continue
		next:
			i = ps[0] - 1
			j = i + sort.SearchInts(a[k][i+1:], a[k][i])
		} else {
			i, j = ps[0]-1, ps[1]
		}
		for _, r := range a {
			r[i], r[j] = r[j], r[i]
			if !sort.IntsAreSorted(r) {
				Fprintln(out, -1)
				continue o
			}
		}
		Fprintln(out, i+1, j+1)
	}
}

//func main() { CF1684C(os.Stdin, os.Stdout) }
