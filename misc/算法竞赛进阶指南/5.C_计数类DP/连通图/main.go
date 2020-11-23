package main

import (
	. "fmt"
	"io"
	"math/big"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	f := [51]*big.Int{nil, big.NewInt(1), big.NewInt(1)}
	for i := int64(3); i <= 50; i++ {
		f[i] = new(big.Int)
		for j := int64(1); j < i; j++ {
			v := new(big.Int).Binomial(i-2, j-1)
			f[i].Add(f[i], v.Mul(v, f[j]).Mul(v, f[i-j]).Mul(v, big.NewInt(1<<j-1)))
		}
	}
	for {
		var n int
		Fscan(in, &n)
		if n == 0 {
			break
		}
		Fprintln(out, f[n])
	}
}

func main() { run(os.Stdin, os.Stdout) }
