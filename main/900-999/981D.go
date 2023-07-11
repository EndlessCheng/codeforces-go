package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func CF981D(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	sum := make([]uint64, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &sum[i])
		sum[i] += sum[i-1]
	}

	ans := uint64(0)
	for i := bits.Len64(sum[n]) - 1; i >= 0; i-- {
		bit := uint64(1) << i
		target := ans | bit
		f := make([]bool, n+1)
		f[0] = true
		for i := 0; i < k; i++ {
			for r := n; r > 0; r-- {
				f[r] = false
				for l, ok := range f[:r] {
					if ok && (sum[r]-sum[l])&target == target {
						f[r] = true
						break
					}
				}
			}
			f[0] = false
		}
		if f[n] {
			ans = target
		}
	}
	Fprint(out, ans)
}

//func main() { CF981D(os.Stdin, os.Stdout) }
