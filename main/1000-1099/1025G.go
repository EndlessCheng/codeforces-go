package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1025G(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, x int
	Fscan(in, &n)
	pow2 := make([]int, n+1)
	pow2[0] = 1
	cnt := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &x)
		pow2[i] = pow2[i-1] * 2 % mod
		if x > 0 {
			cnt[x]++
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans += pow2[cnt[i]] - 1
	}
	Fprint(out, (pow2[n-1]-1-ans%mod+mod)%mod)
}

//func main() { cf1025G(bufio.NewReader(os.Stdin), os.Stdout) }
