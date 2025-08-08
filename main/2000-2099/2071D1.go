package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2071D1(in io.Reader, out io.Writer) {
	var T, n, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &r, &r)
		a := make([]int8, n*2+3)
		pre := make([]int8, n+2)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			pre[i] = pre[i-1] ^ a[i]
		}
		if n%2 == 0 {
			n++
			a[n] = pre[n/2]
			pre[n] = pre[n-1] ^ a[n]
		}
		for i := n + 1; i <= n*2; i++ {
			a[i] = pre[i/2]
		}

		ans := int8(0)
		if r > n*4 {
			r /= 2
			z := bits.TrailingZeros(uint(r))
			d := bits.Len(uint(r)) - bits.Len(uint(n*2))
			if r>>d > n*2 {
				d++
			}
			if z <= d {
				ans = pre[n] & int8(z&1^1)
			} else {
				ans = pre[n]&int8(d&1^1) ^ a[r>>d]
			}
		} else if r > n*2 {
			ans = pre[n]
			r /= 2
			if r%2 == 0 {
				ans ^= a[r]
			}
		} else {
			ans = a[r]
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2071D1(bufio.NewReader(os.Stdin), os.Stdout) }
