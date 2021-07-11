package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func CF814D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]struct{ x, y, r int64 }, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y, &a[i].r)
	}
	ans := int64(0)
	for _, p := range a {
		dep := 0
		for _, q := range a {
			if p.r < q.r && (p.x-q.x)*(p.x-q.x)+(p.y-q.y)*(p.y-q.y) <= (q.r-p.r)*(q.r-p.r) {
				dep++
			}
		}
		if dep == 0 || dep&1 > 0 {
			ans += p.r * p.r
		} else {
			ans -= p.r * p.r
		}
	}
	Fprintf(out, "%.11f", float64(ans)*math.Pi)
}

//func main() { CF814D(os.Stdin, os.Stdout) }
