package main

import (
	. "fmt"
	"io"
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1029F(_r io.Reader, _w io.Writer) {
	ds := func(n int64) (ds [][2]int64) {
		for d := int64(1); d*d <= n; d++ {
			if n%d == 0 {
				ds = append(ds, [2]int64{d, n / d})
			}
		}
		return
	}
	var a, b int64
	Fscan(_r, &a, &b)
	da, db := ds(a), ds(b)
	n := a + b
	for d := int64(math.Sqrt(float64(n))); ; d-- {
		if n%d == 0 {
			d2 := n / d
			i := sort.Search(len(da), func(i int) bool { return da[i][0] > d }) - 1
			j := sort.Search(len(db), func(i int) bool { return db[i][0] > d }) - 1
			if da[i][1] <= d2 || db[j][1] <= d2 {
				Fprint(_w, 2*(d+d2))
				break
			}
		}
	}
}

//func main() { CF1029F(os.Stdin, os.Stdout) }
