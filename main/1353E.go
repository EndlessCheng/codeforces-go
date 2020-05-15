package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1353E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, k int
	var s []byte
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &k, &s)
		ones := 0
		for i := range s {
			s[i] -= '0'
			ones += int(s[i])
		}
		if ones <= 1 {
			Fprintln(out, 0)
			continue
		}
		ans := n
		for st := 0; st < k; st++ {
			for i := st; i < n; i += k {
				if s[i] > 0 {
					cnt := [2]int{}
					for ; i < n; i += k {
						cnt[s[i]]++
						if cnt[0] > cnt[1] {
							break
						}
						if cost := cnt[0] + ones - cnt[1]; cost < ans {
							ans = cost
						}
					}
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1353E(os.Stdin, os.Stdout) }
