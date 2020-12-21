package main

import (
	. "fmt"
	"io"
	"math/big"
)

// github.com/EndlessCheng/codeforces-go
func CF1244C(in io.Reader, out io.Writer) {
	var n, a, b, c, g, r, x0, y0 big.Int
	Fscan(in, &n, &c, &a, &b)
	g.GCD(&x0, &y0, &a, &b)
	c.QuoRem(&c, &g, &r)
	if r.Sign() != 0 {
		Fprint(out, -1)
		return
	}

	a.Quo(&a, &g)
	b.Quo(&b, &g)
	x0.Mul(&x0, &c)
	y0.Mul(&y0, &c)

	y1 := new(big.Int).Mod(&y0, &a) // 最小的非负 y
	k := new(big.Int).Quo(new(big.Int).Sub(&y0, y1), &a)
	x1 := new(big.Int).Add(&x0, new(big.Int).Mul(k, &b))

	left := new(big.Int).Sub(&n, x1)
	left.Sub(left, y1)
	if x1.Sign() < 0 || left.Sign() < 0 {
		Fprint(out, -1)
		return
	}

	Fprint(out, x1, y1, left)
}

//func main() { CF1244C(os.Stdin, os.Stdout) }
