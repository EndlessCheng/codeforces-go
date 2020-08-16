package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/big"
)

func exgcd(a, b int64) (gcd, x, y int64) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, y, x = exgcd(b, a%b)
	y -= a / b * x
	return
}

// github.com/EndlessCheng/codeforces-go
func Sol1244C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, p64, a, b int64
	Fscan(in, &n, &p64, &a, &b)
	gcd, x64, y64 := exgcd(a, b)
	if p64%gcd != 0 {
		Fprint(out, -1)
		return
	}

	ni := big.NewInt
	x, y := ni(x64), ni(y64)
	x.Mul(x, ni(p64/gcd))
	y.Mul(y, ni(p64/gcd))

	a0, b0 := ni(a/gcd), ni(b/gcd)
	newY := (&big.Int{}).Set(y)
	newY.Mod(newY, a0).Add(newY, a0).Mod(newY, a0)
	tmpY := (&big.Int{}).Set(newY)
	x.Sub(x, tmpY.Sub(tmpY, y).Div(tmpY, a0).Mul(tmpY, b0))

	if !x.IsInt64() || !newY.IsInt64() {
		Fprint(out, -1)
		return
	}
	x64, y64 = x.Int64(), newY.Int64()
	if x64 < 0 {
		Fprint(out, -1)
		return
	}
	if x64+y64 > n {
		Fprint(out, -1)
		return
	}
	Fprint(out, x64, y64, n-x64-y64)
}

//func main() {
//	Sol1244C(os.Stdin, os.Stdout)
//}
