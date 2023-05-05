package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1763D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod int64 = 1e9 + 7
	const mx = 100
	F := [mx + 1]int64{1}
	pow2 := [mx + 1]int64{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * int64(i) % mod
		pow2[i] = pow2[i-1] * 2 % mod
	}
	pow := func(x, n int64) (res int64) {
		res = 1
		for ; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return
	}
	invF := [...]int64{mx: pow(F[mx], mod-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * int64(i) % mod
	}
	C := func(n, k int) int64 {
		if k < 0 || k > n {
			return 0
		}
		return F[n] * invF[k] % mod * invF[n-k] % mod
	}

	var T, n, i, j, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &i, &j, &x, &y)
		i--
		j--
		if x > y {
			x, y = y, x
			i, j = n-1-j, n-1-i
		}
		if y == n {
			if j == n-1 {
				Fprintln(out, 0)
			} else {
				Fprintln(out, C(x-1, i)*C(y-x-1, j-i-1)%mod)
			}
		} else {
			ans := C(x-1, i) * (C(y-x-1, j-i-1) + C(y-x-1, n-1-j-(x-1-i))) % mod * pow2[n-1-y]
			if x == i+1 && y == j+1 {
				ans += mod - 1
			}
			Fprintln(out, ans%mod)
		}
	}
}

//func main() { CF1763D(os.Stdin, os.Stdout) }
