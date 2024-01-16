package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1921F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const B = 150
	var sum, sum2 [B][1e5 + B]int // 分块题目推荐用定长数组

	var T, n, q, s, d, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		for d := 1; d < B; d++ {
			for i, v := range a {
				sum[d][i+d] = sum[d][i] + v
				sum2[d][i+d] = sum2[d][i] + v*(i/d+1)
			}
		}

		for ; q > 0; q-- {
			Fscan(in, &s, &d, &k)
			s--
			if d < B {
				r := s + d*k
				Fprint(out, sum2[d][r]-sum2[d][s]-(sum[d][r]-sum[d][s])*(s/d), " ")
			} else {
				res := 0
				for i := 0; i < k; i++ {
					res += a[s+d*i] * (i + 1)
				}
				Fprint(out, res, " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { cf1921F(os.Stdin, os.Stdout) }
