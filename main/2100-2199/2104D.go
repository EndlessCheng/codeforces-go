package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2104D(in io.Reader, out io.Writer) {
	const mx = 5800079
	primes := make([]int, 0, 4e5)
	np := [mx + 1]bool{}
	for i := 2; i <= mx; i++ {
		if !np[i] {
			primes = append(primes, i)
			for j := i * i; j <= mx; j += i {
				np[j] = true
			}
		}
	}
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.SortFunc(a, func(a, b int) int { return b - a })
		ans, s := n, 0
		for i, v := range a {
			s += v - primes[i]
			if s < 0 {
				break
			}
			ans--
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2104D(bufio.NewReader(os.Stdin), os.Stdout) }
