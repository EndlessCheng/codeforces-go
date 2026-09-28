package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf200A(in io.Reader, _w io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, k int
	Fscan(in, &n, &m, &k)
	d := make([][]int, n+1)
	v := make([][]int, n+1)
	for i := range d {
		d[i] = make([]int, m+1)
		v[i] = make([]int, m+1)
	}

	solve := func(x, y, k int) (int, int, bool) {
		l := max(1, x-k)
		r := min(x+k, n)
		for i := l; i <= r; i++ {
			t := k - abs(i-x)
			if y-t > 0 && v[i][y-t] == 0 {
				return i, y - t, true
			}
			if y+t <= m && v[i][y+t] == 0 {
				return i, y + t, true
			}
		}
		return 0, 0, false
	}

	for range k {
		var x, y int
		Fscan(in, &x, &y)
		for i := -2; i <= 2; i++ {
			for j := -2; j <= 2; j++ {
				if x+i < 1 || x+i > n || y+j < 1 || y+j > m {
					continue
				}
				d[x][y] = max(d[x][y], d[x+i][y+j]-abs(i)-abs(j))
			}
		}

		var a, b int
		for {
			var ok bool
			a, b, ok = solve(x, y, d[x][y])
			if ok {
				break
			}
			d[x][y]++
		}

		Fprintln(out, a, b)
		v[a][b] = 1
	}
}

//func main() { cf200A(bufio.NewReader(os.Stdin), os.Stdout) }
