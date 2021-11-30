package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1408C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, pj int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &pj)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := 0.
		i, j, di, dj := 0, n-1, float64(a[0]), float64(pj-a[n-1])
		for {
			if ti, tj := di/float64(i+1), dj/float64(n-j); ti < tj {
				ans += ti
				dj -= float64(n-j) * ti
				if i == j {
					ans += dj / float64(n+2)
					break
				}
				di = float64(a[i+1] - a[i])
				i++
			} else {
				ans += tj
				di -= float64(i+1) * tj
				if i == j {
					ans += di / float64(n+2)
					break
				}
				dj = float64(a[j] - a[j-1])
				j--
			}
		}
		Fprintf(out, "%.15f\n", ans)
	}
}

//func main() { CF1408C(os.Stdin, os.Stdout) }
