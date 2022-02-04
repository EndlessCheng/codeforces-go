package main

import (
	. "fmt"
	"io"
	"math/big"
)

// github.com/EndlessCheng/codeforces-go
func CF888D(in io.Reader, out io.Writer) {
	var n, k, ans int64
	Fscan(in, &n, &k)
	for i, d := range []int64{1, 0, 1, 2, 9}[:k+1] {
		ans += new(big.Int).Binomial(n, int64(i)).Int64() * d
	}
	Fprint(out, ans)
}

//func main() { CF888D(os.Stdin, os.Stdout) }
