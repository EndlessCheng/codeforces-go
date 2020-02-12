package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1303D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func() (ans int) {
		var n, v int64
		var m int
		Fscan(in, &n, &m)
		sum := int64(0)
		cnts := map[int64]int{}
		for i := 0; i < m; i++ {
			Fscan(in, &v)
			sum += v
			cnts[v]++
		}
		if sum < n {
			return -1
		}

		for i := int64(1); i <= n; i <<= 1 {
			if n&i > 0 {
				if cnts[i] == 0 {
					j := i << 1
					for ; cnts[j] == 0; j <<= 1 {
					}
					for ; j > i; j >>= 1 {
						cnts[j]--
						cnts[j>>1] += 2
						ans++
					}
				}
				cnts[i]--
			}
			cnts[i<<1] += cnts[i] >> 1
			cnts[i] &= 1
		}
		return
	}

	var t int
	for Fscan(in, &t); t > 0; t-- {
		Fprintln(out, solve())
	}
}

//func main() { CF1303D(os.Stdin, os.Stdout) }
