package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1788D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 1_000_000_007
	const mx = 3000
	pow2 := [mx + 1]int64{1}
	for i := 1; i <= mx; i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	ans := int64(0)
	for i, ai := range a {
		l, r := i-1, i+2
		for j := i + 1; j < n; j++ {
			aj := a[j]
			for l >= 0 && a[l] >= ai*2-aj {
				l--
			}
			for r <= j || r < n && a[r] < aj*2-ai {
				r++
			}
			ans += pow2[l+1+n-r]
		}
	}
	Fprint(out, ans%mod)
}

//func main() { CF1788D(os.Stdin, os.Stdout) }
