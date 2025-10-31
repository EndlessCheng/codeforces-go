package P2000_2999

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p2513(in io.Reader, out io.Writer) {
	const mod = 10000
	var n, k int
	Fscan(in, &n, &k)
	if n*(n-1)/2 < k {
		Fprint(out, 0)
		return
	}
	f := make([]int, k+1)
	f[0] = 1
	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			f[j] = (f[j] + f[j-1]) % mod
		}
		for j := k; j > i; j-- {
			f[j] = (f[j] - f[j-i-1] + mod) % mod
		}
	}
	Fprint(out, f[k])
}

//func main() { p2513(os.Stdin, os.Stdout) }
