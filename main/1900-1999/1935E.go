package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1935E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, q, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		sum := make([][30]int, n+1)
		variable := make([][30]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &l, &r)
			sum[i] = sum[i-1]
			variable[i] = variable[i-1]
			for s := uint(r); s > 0; s &= s - 1 {
				j := bits.TrailingZeros(s)
				sum[i][j]++
				if l>>j < r>>j {
					variable[i][j]++
				}
			}
		}

		Fscan(in, &q)
		for range q {
			Fscan(in, &l, &r)
			l--
			ans := 0
			for i := 29; i >= 0; i-- {
				v := sum[r][i] - sum[l][i]
				if v == 0 {
					continue
				}
				ans |= 1 << i
				if v > 1 && variable[r][i] > variable[l][i] {
					ans |= 1<<i - 1
					break
				}
			}
			Fprint(out, ans, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1935E(bufio.NewReader(os.Stdin), os.Stdout) }
