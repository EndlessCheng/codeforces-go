package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1358D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	var x, ans int64
	Fscan(in, &n, &x)
	d := make([]int64, n)
	for i := range d {
		Fscan(in, &d[i])
	}
	d = append(d, d...)
	n *= 2
	sumD := make([]int64, n+1)
	sumH := make([]int64, n+1)
	for i, c := range d {
		sumD[i+1] = sumD[i] + c
		sumH[i+1] = sumH[i] + c*(c+1)/2
	}
	for r := 1; r <= n; r++ {
		l := sort.Search(r, func(i int) bool { return sumD[r]-sumD[i] <= x })
		sd := sumD[r] - sumD[l]
		sh := sumH[r] - sumH[l]
		if l > 0 {
			c := x - sd
			sh += c * (2*d[l-1] - c + 1) / 2
		}
		if sh > ans {
			ans = sh
		}
	}
	Fprint(_w, ans)
}

//func main() { CF1358D(os.Stdin, os.Stdout) }
