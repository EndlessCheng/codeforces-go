package main

import (
	"bufio"
	. "fmt"
	"io"
)

func Sol981D(reader io.Reader, writer io.Writer) {
	found := func(bits uint64, sum []uint64, k int) bool {
		n := len(sum) - 1
		var dp [51][51]bool
		for r := 1; r <= n; r++ {
			dp[1][r] = bits&sum[r] == bits
		}
		for i := 2; i <= k; i++ {
			for l := i - 1; l < n; l++ {
				if dp[i-1][l] {
					for r := l + 1; r <= n; r++ {
						dp[i][r] = dp[i][r] || bits&(sum[r]-sum[l]) == bits
					}
				}
			}
		}
		return dp[k][n]
	}

	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	sum := make([]uint64, n+1)
	for i := 1; i <= n; i++ {
		var a uint64
		Fscan(in, &a)
		sum[i] = sum[i-1] + a
	}

	ans := uint64(0)
	for i := 55; i >= 0; i-- {
		pow2 := uint64(1) << uint(i)
		if found(ans|pow2, sum, k) {
			ans |= pow2
		}
	}
	Fprintln(out, ans)
}

//func main() {
//	Sol981D(os.Stdin, os.Stdout)
//}
