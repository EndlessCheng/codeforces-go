package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf425E(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, K int
	Fscan(in, &n, &K)
	pow2 := make([]int, n*n+1)
	pow2[0] = 1
	for i := 1; i <= n*n; i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	f[0][0] = 1
	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			for k := i + 1; k <= n; k++ {
				f[k][j+1] = (f[k][j+1] + f[i][j]*(pow2[k-i]-1+mod)%mod*pow2[(n-k)*(k-i)]%mod) % mod
			}
		}
	}

	ans := 0
	for i := K; i <= n; i++ {
		ans = (ans + f[i][K]) % mod
	}
	Fprint(out, ans)
}

//func main() { cf425E(bufio.NewReader(os.Stdin), os.Stdout) }
