package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// dp[i][3] 表示由 i 位数字构成的魔鬼数的个数
	// dp[i][j] (j<3) 表示 i 位数字构成的、开头有连续 j 个 6 的非魔鬼数的个数
	const mx = 11
	const cont = 3
	dp := [mx][cont + 1]int{}
	dp[0][0] = 1
	for i := 1; i < mx; i++ {
		for j := 0; j < cont; j++ {
			dp[i][0] += dp[i-1][j] * 9 // 开头无 6，直接转移（0-9 中除去 6 共 9 个数）
			dp[i][j+1] = dp[i-1][j]    // 开头有 j+1 个 6，下一个有 j 个 6
		}
		dp[i][cont] += dp[i-1][cont] * 10
	}

	var T, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x)
		n := 1
		for ; dp[n][3] < x; n++ {
		}
		has6 := 0
		for i := 1; i <= n; i++ {
			for d := 0; d <= 9; d++ { // 试填
				need6 := 3
				if has6 == 3 {
					need6 = 0
				} else if d == 6 {
					need6 = 3 - 1 - has6
				}
				sum := 0
				for j := need6; j <= 3; j++ {
					sum += dp[n-i][j]
				}
				if sum >= x {
					Fprint(out, d)
					if has6 < 3 {
						if d == 6 {
							has6++
						} else {
							has6 = 0
						}
					}
					break
				}
				x -= sum
			}
		}
		Fprintln(out)
	}
}

func main() { run(os.Stdin, os.Stdout) }
