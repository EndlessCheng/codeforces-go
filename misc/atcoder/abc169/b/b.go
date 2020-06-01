package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/big"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] == 0 {
			Fprint(out, 0)
			return
		}
	}

	upper := big.NewInt(1e18)
	ans := big.NewInt(1)
	for _, v := range a {
		ans.Mul(ans, big.NewInt(v))
		if ans.Cmp(upper) > 0 {
			Fprint(out, -1)
			return
		}
	}
	Fprint(out, ans.Int64())
}

func main() { run(os.Stdin, os.Stdout) }
