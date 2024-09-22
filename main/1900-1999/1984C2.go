package main

import (
	. "fmt"
	"io"
	"slices"
)

func cf1984C2(in io.Reader, out io.Writer) {
	const mod = 998244353
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := 1
		s := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &s[i])
			s[i] += s[i-1]
			if s[i] >= 0 {
				ans = ans * 2 % mod
			}
		}
		minS := slices.Min(s)
		if minS >= 0 {
			Fprintln(out, ans)
			continue
		}
		sumNeg := 0
		neg := 1
		for i := n; i > 0; i-- {
			if s[i] == minS {
				sumNeg += neg
			}
			if s[i] < 0 {
				neg = neg * 2 % mod
			}
		}
		Fprintln(out, sumNeg%mod*ans%mod)
	}
}

//func main() { cf1984C2(bufio.NewReader(os.Stdin), os.Stdout) }
