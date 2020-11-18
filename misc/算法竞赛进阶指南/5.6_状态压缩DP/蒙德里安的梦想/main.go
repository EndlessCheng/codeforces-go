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

	var n, m int
	for {
		Fscan(in, &n, &m)
		if n == 0 {
			break
		}
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, 1<<m)
		}
		dp[0][0] = 1
		for i, d := range dp[:n] {
			for cur, v := range d {
				if v == 0 {
					continue
				}
				var f func(j, next int)
				f = func(j, next int) {
					if j == m {
						dp[i+1][next] += dp[i][cur]
						return
					}
					// 下一行该列必须是 0
					if cur>>j&1 > 0 {
						f(j+1, next)
						return
					}
					// 下一行该列可以是 1
					f(j+1, next|1<<j)
					// 也可以是 00
					if j+1 < m && cur>>(j+1)&1 == 0 {
						f(j+2, next)
					}
				}
				// 搜索下一行的所有可能状态，进行转移
				f(0, 0)
			}
		}
		Fprintln(out, dp[n][0])
	}
}

func main() { run(os.Stdin, os.Stdout) }
