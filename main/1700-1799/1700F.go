package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1700F(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var n int
	Fscan(in, &n)
	a := make([][4]int, n)
	for j := range 4 {
		for i := range a {
			Fscan(in, &a[i][j])
		}
	}

	u, d := 0, 0
	ans := 0
	for _, p := range a {
		u += p[0] - p[2]
		d += p[1] - p[3]
		if u > 0 && d < 0 {
			t := min(u, -d)
			ans += t
			u -= t
			d += t
		}
		if u < 0 && d > 0 {
			t := min(-u, d)
			ans += t
			u += t
			d -= t
		}
		ans += abs(u) + abs(d)
	}

	if u != 0 || d != 0 {
		Fprint(out, -1)
	} else {
		Fprint(out, ans)
	}
}

//func main() { cf1700F(bufio.NewReader(os.Stdin), os.Stdout) }
