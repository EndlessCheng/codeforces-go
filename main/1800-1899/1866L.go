package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1866L(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var n, m int
	Fscan(in, &n, &m)
	mn, ans := 0, 1
	for k := 1; k <= m; k++ {
		n2 := n / gcd(n, k)
		cnt := 0
		if k == 1 {
			continue
		}
		for j := 0; j <= k; j++ {
			x := max((j*n+k)/k, (j*n+k-1)/(k-1))
			y := min((j+1)*n/k, n2)
			if x <= y {
				cnt += (y - x + 1) * ((x+y)*k/2 - j*n)
			}
		}
		if mn < cnt {
			mn = cnt
			ans = k
		}
	}
	Fprint(out, ans)
}

//func main() { cf1866L(bufio.NewReader(os.Stdin), os.Stdout) }
