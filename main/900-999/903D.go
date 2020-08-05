package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/big"
)

// github.com/EndlessCheng/codeforces-go
func CF903D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, y, s int64
	ans := &big.Int{}
	cnt := map[int64]int64{}
	Fscan(in, &n)
	for i := int64(0); i < n; i++ {
		Fscan(in, &y)
		ans.Add(ans, big.NewInt(i*y-s))
		for x := y - 1; x <= y+1; x++ {
			ans.Sub(ans, big.NewInt(cnt[x]*(y-x)))
		}
		s += y
		cnt[y]++
	}
	Fprint(out, ans)
}

//func main() { CF903D(os.Stdin, os.Stdout) }
