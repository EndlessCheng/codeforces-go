package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/big"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1536E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		c := 0
		for i := 0; i < n; i++ {
			Fscan(in, &s)
			c += strings.Count(s, "#")
		}
		ans := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(c)), big.NewInt(1e9+7)).Int64()
		if c == n*m {
			ans--
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1536E(os.Stdin, os.Stdout) }
