package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1610D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int64 = 1e9 + 7
	pow := func(x int64, n uint) (res int64) {
		res = 1
		for ; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return
	}

	cnt := [30]uint{}
	var n, v, c uint
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		cnt[bits.TrailingZeros(v)]++
	}
	ans := int64(0)
	for i := 29; i > 0; i-- {
		if cnt[i] > 0 {
			ans += pow(2, c) * (pow(2, cnt[i]-1) - 1)
			c += cnt[i]
		}
	}
	ans += pow(2, c) * (pow(2, cnt[0]) - 1)
	Fprint(out, ans%mod)
}

//func main() { CF1610D(os.Stdin, os.Stdout) }
