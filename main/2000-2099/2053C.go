package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2053C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		m := n + 1
		ans := 0
		if n&1 > 0 {
			ans = m / 2
		}

		n /= 2
		if n >= k {
			d := bits.Len(uint(n)) - bits.Len(uint(k))
			if n>>d >= k {
				d++
			}
			ans += n & (1<<d - 1) * m
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2053C(bufio.NewReader(os.Stdin), os.Stdout) }
