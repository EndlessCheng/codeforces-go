package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF815A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, m, sr, sc int
	Fscan(in, &n, &m)
	a := make([][]int, n)
	minR := make([]int, n)
	minC := make([]int, m)
	for j := range minC {
		minC[j] = 999
	}
	for i := range a {
		a[i] = make([]int, m)
		minR[i] = 999
		for j := range a[i] {
			Fscan(in, &a[i][j])
			if i > 0 && a[i][j]-a[i-1][j] != a[i][0]-a[i-1][0] {
				Fprint(out, -1)
				return
			}
			minR[i] = min(minR[i], a[i][j])
			minC[j] = min(minC[j], a[i][j])
		}
		sr += minR[i]
	}
	for j, v := range a[0] {
		sr += v - minR[0]
		sc += minC[j]
	}
	for _, r := range a {
		sc += r[0] - minC[0]
	}
	if sr < sc {
		Fprintln(out, sr)
		for i, v := range minR {
			Fprint(out, strings.Repeat(Sprintln("row", i+1), v))
		}
		for j, v := range a[0] {
			Fprint(out, strings.Repeat(Sprintln("col", j+1), v-minR[0]))
		}
	} else {
		Fprintln(out, sc)
		for j, v := range minC {
			Fprint(out, strings.Repeat(Sprintln("col", j+1), v))
		}
		for i, r := range a {
			Fprint(out, strings.Repeat(Sprintln("row", i+1), r[0]-minC[0]))
		}
	}
}

//func main() { CF815A(os.Stdin, os.Stdout) }
