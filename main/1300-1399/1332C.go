package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1332C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, k int
	var s []byte
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &k, &s)
		ans := 0
		for i := 0; i < k/2; i++ {
			cnt := [26]int{}
			for j := i; j < n; j += k {
				cnt[s[j]-'a']++
			}
			for j := k - 1 - i; j < n; j += k {
				cnt[s[j]-'a']++
			}
			max := 0
			for _, c := range cnt {
				if c > max {
					max = c
				}
			}
			ans += 2*n/k - max
		}
		if k&1 == 1 {
			cnt := [26]int{}
			for j := k / 2; j < n; j += k {
				cnt[s[j]-'a']++
			}
			max := 0
			for _, c := range cnt {
				if c > max {
					max = c
				}
			}
			ans += n/k - max
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1332C(os.Stdin, os.Stdout) }
