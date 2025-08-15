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
然后取 dfs(i-x,j) 和 dfs(i,j-x) 的最小值？
*/

// https://github.com/EndlessCheng
func cf2075D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	dp := [60][60][60]int{}
	var f func(int, int, int) int
	f = func(i, j, low int) int {
		if i == 0 && j == 0 {
			return 0
		}
		if i < 0 || j < 0 || low > max(i, j) {
			return 8e18
		}
		p := &dp[i][j][low]
		if *p > 0 {
			return *p
		}
		res := min(f(i-low, j, low+1)+1<<low, f(i, j-low, low+1)+1<<low, f(i, j, low+1))
		*p = res
		return res
	}

	dp3 := [60][60]int{}
	var f3 func(int, int) int
	f3 = func(i, low int) int {
		if i <= 0 {
			return 0
		}
		if low >= i {
			return 1 << low
		}
		p := &dp3[i][low]
		if *p > 0 {
			return *p
		}
		res := min(f3(i-low, low+1)+1<<low, f3(i, low+1))
		*p = res
		return res
	}

	dp2 := [60][60][60]int{}
	var f2 func(int, int, int) int
	f2 = func(i, j, low int) int {
		if i <= 0 && j <= 0 {
			return 0
		}
		if i <= 0 {
			return f3(j, low)
		}
		if j <= 0 {
			return f3(i, low)
		}
		if low >= max(i, j) {
			return 3 << low
		}
		p := &dp2[i][j][low]
		if *p > 0 {
			return *p
		}
		res := min(f2(i-low, j, low+1)+1<<low, f2(i, j-low, low+1)+1<<low, f2(i, j, low+1))
		*p = res
		return res
	}

	var T, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x, &y)
		if x > y {
			x, y = y, x
		}
		n, m := bits.Len(uint(x)), bits.Len(uint(y))
		ans := f2(n, m, 1)
		for a := bits.Len(uint(x ^ y>>(m-n))); a < n; a++ {
			ans = min(ans, f(a, a+m-n, 1))
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2075D(bufio.NewReader(os.Stdin), os.Stdout) }
