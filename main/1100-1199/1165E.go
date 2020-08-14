package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func Sol1165E(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	const mod int64 = 998244353
	var n int
	Fscan(in, &n)
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
		// https://oeis.org/A003991
		a[i] *= int64(n-i) * int64(i+1)
	}
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
	}
	sort.Ints(b)

	ans := int64(0)
	for i, ai := range a {
		ans = (ans + ai%mod*int64(b[i])) % mod
	}
	Fprint(out, ans)
}

//func main() {
//	Sol1165E(os.Stdin, os.Stdout)
//}
