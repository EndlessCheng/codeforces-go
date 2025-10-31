package main

import (
	. "fmt"
	"io"
	"math"
	"math/big"
)

// https://space.bilibili.com/206214
func p1287(in io.Reader, out io.Writer) {
	var n, k, ans int
	Fscan(in, &n, &k)
	powN := 1
	for i := 0; i <= k; i++ {
		c := int(new(big.Int).Binomial(int64(k), int64(i)).Int64()) * int(math.Pow(float64(i), float64(n)))
		if (k-i)%2 == 0 {
			ans += c
		} else {
			ans -= c
		}
		powN *= n
	}
	Fprint(out, ans)
}

//func main() { p1287(os.Stdin, os.Stdout) }
