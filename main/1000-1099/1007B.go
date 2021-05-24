package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1007B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	const MX int = 1e5
	d := [MX + 1]int{}
	for i := 1; i <= MX; i++ {
		for j := i; j <= MX; j += i {
			d[j]++
		}
	}
	const mx = 140
	C := [mx][4]int64{}
	C[0][0] = 1
	for i := 1; i < mx; i++ {
		C[i][0] = 1
		for j := 1; j < 4; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}
	perm3 := [][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}}

	var T, a, b, c int
	x := [8]int{}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &c)
		a, b, c, ab, ac, bc, abc := d[a], d[b], d[c], d[gcd(a, b)], d[gcd(a, c)], d[gcd(b, c)], d[gcd(gcd(a, b), c)]
		x[1] = a - ab - ac + abc // 某个类型的因子个数
		x[2] = b - ab - bc + abc
		x[4] = c - ac - bc + abc
		x[3] = ab - abc
		x[5] = ac - abc
		x[6] = bc - abc
		x[7] = abc

		ans := int64(0)
		t := [3]int{}
		for t[0] = 1; t[0] < 8; t[0]++ {
			for t[1] = t[0]; t[1] < 8; t[1]++ {
				for t[2] = t[1]; t[2] < 8; t[2]++ {
					for _, p := range perm3 {
						if t[p[0]]&1 > 0 && t[p[1]]&2 > 0 && t[p[2]]&4 > 0 {
							goto calc
						}
					}
					continue
				calc:
					c := [8]int{}
					c[t[0]]++
					c[t[1]]++
					c[t[2]]++
					res := int64(1)
					for i := 1; i < 8; i++ {
						if c[i] > 0 {
							res *= C[x[i]+c[i]-1][c[i]] // 重复选择
						}
					}
					ans += res
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1007B(os.Stdin, os.Stdout) }
