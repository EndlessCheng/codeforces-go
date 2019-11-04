package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol431C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	const mod = 1e9 + 7
	var n, k, d int
	Fscan(in, &n, &k, &d)
	dp := make([][2]int, n)
	for i := range dp {
		dp[i][0] = -1
		dp[i][1] = -1
	}
	var f func(int, int) int
	f = func(sum int, contain int) int {
		if sum == n {
			return 1
		}
		if contain == 0 && sum+d > n {
			return 0
		}
		if v := dp[sum][contain]; v != -1 {
			return v
		}
		ans := 0
		if contain == 0 {
			for i := 1; i < d; i++ {
				if sum+i > n {
					break
				}
				ans = (ans + f(sum+i, 0)) % mod
			}
			for i := d; i <= k; i++ {
				if sum+i > n {
					break
				}
				ans = (ans + f(sum+i, 1)) % mod
			}
		} else {
			for i := 1; i <= k; i++ {
				if sum+i > n {
					break
				}
				ans = (ans + f(sum+i, 1)) % mod
			}
		}
		dp[sum][contain] = ans
		return ans
	}
	Fprint(out, f(0, 0))
}

//func main() {
//	Sol431C(os.Stdin, os.Stdout)
//}
