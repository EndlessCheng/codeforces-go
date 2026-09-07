package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1687D(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	l := make([]int, n+1)
	r := make([]int, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
		l[i] = i - 1
		r[i] = i + 1
	}

	pos := [][]int{}
	for i := 1; i < n; i++ {
		d := a[i] - a[i-1]
		for len(pos) <= d {
			pos = append(pos, nil)
		}
		pos[d] = append(pos[d], i)
	}

	for i := 0; ; i++ {
		if i < len(pos) {
			for _, j := range pos[i] {
				r[l[j]] = r[j]
				l[r[j]] = l[j]
			}
		}

		s := i*i - a[0]
		x, y := 0, s+i
		j := i
		for l := 0; l < n; l = r[l] {
			for j*(j+1) < a[l]+s {
				j++
			}
			x = max(x, j*j-a[l])
			y = min(y, j*(j+1)-a[r[l]-1])
		}

		if x <= y {
			Fprint(out, x)
			return
		}
	}
}

//func main() { cf1687D(bufio.NewReader(os.Stdin), os.Stdout) }
