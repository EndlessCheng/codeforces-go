package main

import (
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func cf293C(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	if n%3 > 0 {
		Fprint(out, 0)
		return
	}
	n /= 3
	for i := 2; i*i*i <= n; i++ {
		if n%i > 0 {
			continue
		}
		m := n / i
		for j := int(math.Sqrt(float64(m))); j >= i && j > m/j-i; j-- {
			if m%j > 0 {
				continue
			}
			k := m / j
			if (i+j+k)%2 > 0 {
				continue
			}
			if i == k {
				ans++
			} else if i == j || j == k {
				ans += 3
			} else {
				ans += 6
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf293C(os.Stdin, os.Stdout) }
