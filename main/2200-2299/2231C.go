package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2231C(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		cnt := map[int]int{}
		sum := map[int]int{}
		for i := range n {
			Fscan(in, &v)
			x := v
			step := 0
			for {
				if i == 0 || cnt[v] > 0 {
					cnt[v]++
					sum[v] += step
				}
				step++
				if v&1 > 0 {
					v++
				} else {
					v >>= 1
				}
				if v == 1 {
					break
				}
			}
			if x > 1 {
				cnt[1]++
				sum[1] += step
			}
		}

		ans := int(1e18)
		for v, c := range cnt {
			if c == n {
				ans = min(ans, sum[v])
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2231C(bufio.NewReader(os.Stdin), os.Stdout) }
