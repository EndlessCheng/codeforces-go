package main

import (
	. "fmt"
	"io"
	"math/big"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var n int64
	Fscan(in, &n)
	Fprint(out, new(big.Int).Binomial(n-1, 11))
}

func main() { run(os.Stdin, os.Stdout) }
