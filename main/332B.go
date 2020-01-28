package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF332B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct {
		v int64
		i int
	}
	const mx = 18
	var st [][mx]pair
	stInit := func(a []int64) {
		n := len(a)
		st = make([][mx]pair, n)
		for i := range st {
			st[i][0] = pair{a[i], i}
		}
		for j := uint(1); 1<<j <= n; j++ {
			for i := 0; i+(1<<j)-1 < n; i++ {
				st0, st1 := st[i][j-1], st[i+(1<<(j-1))][j-1]
				if st0.v >= st1.v {
					st[i][j] = st0
				} else {
					st[i][j] = st1
				}
			}
		}
	}
	stQuery := func(l, r int) pair {
		k := uint(bits.Len(uint(r-l)) - 1)
		st0, st1 := st[l][k], st[r-(1<<k)][k]
		if st0.v >= st1.v {
			return st0
		}
		return st1
	}

	var n, k, ansA, ansB int
	var max int64
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sumK := make([]int64, n+1-k)
	for _, v := range a[:k] {
		sumK[0] += int64(v)
	}
	for i := k; i < n; i++ {
		sumK[i-k+1] = sumK[i-k] + int64(a[i]-a[i-k])
	}
	stInit(sumK)
	for i, sumA := range sumK[:n-2*k+1] {
		p := stQuery(i+k, n+1-k)
		if sumA+p.v > max {
			max = sumA + p.v
			ansA = i
			ansB = p.i
		}
	}
	Fprint(out, ansA+1, ansB+1)
}

//func main() {
//	CF332B(os.Stdin, os.Stdout)
//}
