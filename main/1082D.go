package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1082D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, d int
	Fscan(in, &n)
	d1 := []int{}
	d2 := [][2]int{}
	for i := 1; i <= n; i++ {
		Fscan(in, &d)
		if d == 1 {
			d1 = append(d1, i)
		} else {
			d2 = append(d2, [2]int{i, d})
		}
	}
	for len(d1) < 2 {
		d1 = append(d1, d2[0][0])
		d2 = d2[1:]
	}
	if len(d2) == 0 {
		Fprint(out, "NO")
		return
	}
	ans := make([][2]int, 0, n-1)
	ans = append(ans,
		[2]int{d1[0], d2[0][0]},
		[2]int{d1[1], d2[len(d2)-1][0]},
	)
	for i := 1; i < len(d2); i++ {
		ans = append(ans, [2]int{d2[i-1][0], d2[i][0]})
	}
o:
	for _, v := range d1[2:] {
		for i, p := range d2 {
			if p[1] > 2 {
				ans = append(ans, [2]int{v, p[0]})
				d2[i][1]--
				continue o
			}
		}
		Fprintln(out, "NO")
		return
	}
	Fprintln(out, "YES", len(d2)+1)
	Fprintln(out, n-1)
	for _, v := range ans {
		Fprintln(out, v[0], v[1])
	}
}

//func main() { CF1082D(os.Stdin, os.Stdout) }
