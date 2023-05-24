package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF711D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 1_000_000_007
	pow := func(x int64, n int) (res int64) {
		x %= mod
		res = 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return
	}

	var n int
	Fscan(in, &n)
	to := make([]int, n)
	for i := range to {
		Fscan(in, &to[i])
		to[i]--
	}

	ans := int64(1)
	time := make([]int, n)
	clock := 1
	for x, t := range time {
		if t > 0 {
			continue
		}
		for t0 := clock; x >= 0; x = to[x] {
			if time[x] > 0 {
				if time[x] >= t0 {
					sz := clock - time[x]
					n -= sz
					ans = ans * (pow(2, sz) - 2) % mod
				}
				break
			}
			time[x] = clock
			clock++
		}
	}
	Fprint(out, (ans*pow(2, n)%mod+mod)%mod)
}

//func main() { CF711D(os.Stdin, os.Stdout) }
