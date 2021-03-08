package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1493C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s)
		if n%k > 0 {
			Fprintln(out, -1)
			continue
		}
		cnt := make([]int, 26)
		for _, b := range s {
			cnt[b-'a']++
		}
	O:
		for _, c := range cnt {
			if c%k == 0 {
				continue
			}
			for i := n - 1; i >= 0; i-- {
			o:
				for s[i]++; s[i] <= 'z'; s[i]++ {
					cnt[s[i]-1-'a']--
					cnt[s[i]-'a']++
					left := n - 1 - i
					for _, c := range cnt {
						need := (k - c%k) % k
						if left < need {
							continue o
						}
						left -= need
					}
					for i++; left > 0; left-- { // 没注意到上面循环结束时还有 left > 0 的情况，RE 了一次
						s[i] = 'a'
						i++
					}
					for p := byte(0); i < n; i++ {
						for ; cnt[p]%k == 0; p++ {
						}
						s[i] = 'a' + p
						cnt[p]++
					}
					break O
				}
				cnt[25]--
			}
		}
		Fprintln(out, string(s))
	}
}

//func main() { CF1493C(os.Stdin, os.Stdout) }
