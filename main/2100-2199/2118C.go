package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2118C(in io.Reader, out io.Writer) {
	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		cnt := [60]int{}
		ans := 0
		for range n {
			Fscan(in, &v)
			ans += bits.OnesCount(uint(v))
			for i := range cnt {
				cnt[i] += v>>i&1 ^ 1
			}
		}
		for i, c := range cnt {
			if k>>i <= c {
				ans += k >> i
				break
			}
			ans += c
			k -= c << i
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2118C(bufio.NewReader(os.Stdin), os.Stdout) }
