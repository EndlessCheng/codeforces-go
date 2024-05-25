package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1301E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, q, r1, c1, r2, c2 int
	Fscan(in, &n, &m, &q)
	size := make([][]int16, n)
	f := make([][]int16, n)
	for i := range size {
		size[i] = make([]int16, m)
		f[i] = make([]int16, m)
	}
	a := make([][]byte, n)
	mxSz := 0
	for i := range a {
		Fscan(in, &a[i])
		for j, c := range a[i] {
			if i == 0 || j == 0 || c != a[i][j-1] || c != a[i-1][j] || c != a[i-1][j-1] {
				f[i][j] = 1
			} else {
				f[i][j] = min(f[i][j-1], f[i-1][j], f[i-1][j-1]) + 1
			}
			sz16 := f[i][j]
			sz := int(sz16)
			if i-sz >= sz-1 && j-sz >= sz-1 &&
				a[i][j] == 'B' &&
				a[i][j-sz] == 'Y' && f[i][j-sz] == sz16 &&
				a[i-sz][j] == 'G' && f[i-sz][j] == sz16 &&
				a[i-sz][j-sz] == 'R' && f[i-sz][j-sz] >= sz16 {
				size[i][j] = sz16
				mxSz = max(mxSz, sz)
			}
		}
	}

	sums := make([][][]int32, mxSz+1)
	for k := 1; k <= mxSz; k++ {
		sum := make([][]int32, n+1)
		for i := range sum {
			sum[i] = make([]int32, m+1)
		}
		for i, row := range size {
			for j, sz := range row {
				sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j]
				if sz == int16(k) {
					sum[i+1][j+1]++
				}
			}
		}
		sums[k] = sum
	}
	query := func(k, r1, c1, r2, c2 int) int32 {
		sum := sums[k]
		return sum[r2][c2] - sum[r2][c1] - sum[r1][c2] + sum[r1][c1]
	}

	for ; q > 0; q-- {
		Fscan(in, &r1, &c1, &r2, &c2)
		r1--
		c1--
		ans := sort.Search(min((r2-r1)/2, (c2-c1)/2, mxSz), func(low int) bool {
			low++
			return query(low, r1+low*2-1, c1+low*2-1, r2, c2) == 0
		})
		Fprintln(out, ans*ans*4)
	}
}

//func main() { cf1301E(bufio.NewReader(os.Stdin), os.Stdout) }
