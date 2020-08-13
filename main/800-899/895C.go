package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF895C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod int64 = 1e9 + 7
	cnt := [71]int64{}
	var n, v int
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		cnt[v]++
	}
	pow2 := make([]int64, n)
	pow2[0] = 1
	for i := 1; i < n; i++ {
		pow2[i] = pow2[i-1] << 1 % mod
	}
	ps := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67}

	dp := [2][1 << 19]int64{}
	dp[0][0] = 1
	cur := 0
	for i, c := range cnt {
		if c == 0 {
			continue
		}
		e, x := 0, i
		for j, p := range ps {
			for ; x%p == 0; x /= p {
				e ^= 1 << j
			}
		}
		cur ^= 1
		dp[cur] = [1 << 19]int64{}
		for s := 0; s < 1<<19; s++ {
			v := dp[cur^1][s] * pow2[cnt[i]-1]
			dp[cur][s^e] = (dp[cur][s^e] + v) % mod
			dp[cur][s] = (dp[cur][s] + v) % mod
		}
	}
	Fprint(out, (dp[cur][0]-1+mod)%mod)
}

//func main() { CF895C(os.Stdin, os.Stdout) }