package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1279D(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Split(bufio.ScanWords)
	r := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b&15)
		}
		return
	}
	const mod int64 = 998244353
	pow := func(x int64) int64 {
		x %= mod
		res := int64(1)
		for n := mod - 2; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	div := func(a, b int64) int64 { return a * pow(b) % mod }

	n := r()
	cnts := [1e6 + 1]int{}
	ps := make([][]int, n)
	for i := range ps {
		ps[i] = make([]int, r())
		for j := range ps[i] {
			ps[i][j] = r()
			cnts[ps[i][j]]++
		}
	}
	ans := int64(0)
	for _, p := range ps {
		s := 0
		for _, v := range p {
			s += cnts[v]
		}
		ans = (ans + div(int64(s), int64(len(p)))) % mod
	}
	Fprint(_w, div(ans, int64(n)*int64(n)))
}

//func main() { CF1279D(os.Stdin, os.Stdout) }
