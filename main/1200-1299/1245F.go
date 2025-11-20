package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1245F(in io.Reader, out io.Writer) {
	var T, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &l, &r)
		type args struct {
			p          int
			a, b, c, d bool
		}
		memo := map[args]int{}
		var dfs func(int, bool, bool, bool, bool) int
		dfs = func(p int, limHigh1, limLow1, limHigh2, limLow2 bool) (res int) {
			if p < 0 {
				return 1
			}
			t := args{p, limHigh1, limLow1, limHigh2, limLow2}
			if v, ok := memo[t]; ok {
				return v
			}

			hi1 := 1
			if limHigh1 {
				hi1 = r >> p & 1
			}
			lo1 := 0
			if limLow1 {
				lo1 = l >> p & 1
			}

			hi2 := 1
			if limHigh2 {
				hi2 = r >> p & 1
			}
			lo2 := 0
			if limLow2 {
				lo2 = l >> p & 1
			}

			for i := lo1; i <= hi1; i++ {
				for j := lo2; j <= hi2; j++ {
					if i == 0 || j == 0 {
						res += dfs(p-1, limHigh1 && i == hi1, limLow1 && i == lo1, limHigh2 && j == hi2, limLow2 && j == lo2)
					}
				}
			}
			memo[t] = res
			return
		}
		Fprintln(out, dfs(bits.Len(uint(r))-1, true, true, true, true))
	}
}

//func main() { cf1245F(bufio.NewReader(os.Stdin), os.Stdout) }
