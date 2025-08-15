package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

/*
10 = a+b+c
5 = d+e
1+2+3+6 = 12
1+2+4+5 = 12
看分解后的最大的数

dfs(i,j)
找最小的数 x，满足 1+2+...+x >= i+j
然后取 dfs(i-x,j) 和 dfs(i,j-x) 的最小值
*/

// https://github.com/EndlessCheng
func cf2075D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	dp := [60][60][60]int{}
	var f func(int, int, int) int
	f = func(i, j, low int) (res int) {
		if i == 0 && j == 0 {
			return
		}
		if i < 0 || j < 0 || low > max(i, j) {
			return 8e18
		}
		dv := &dp[i][j][low]
		if *dv > 0 {
			return *dv
		}
		res = f(i-low, j, low+1) + 1<<low
		res2 := f(i, j-low, low+1) + 1<<low
		res3 := f(i, j, low+1)
		res = min(res, res2, res3)
		*dv = res
		return
	}

	var T, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x, &y)
		if x > y {
			x, y = y, x
		}
		n, m := bits.Len(uint(x)), bits.Len(uint(y))
		d := bits.Len(uint(x<<(m-n) ^ y))
		d2 := d - (m - n)
		if d == d2 && (d == 1 || d == 2) {
			d, d2 = 3, 3
		}
		Fprintln(out, f(d, d2, 1))
	}
}

//func main() { cf2075D(bufio.NewReader(os.Stdin), os.Stdout) }
