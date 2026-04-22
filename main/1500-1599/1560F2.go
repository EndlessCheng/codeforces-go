package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1560F2(in io.Reader, out io.Writer) {
	var T, k int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &k)
		cnt := [10]int{}
		d := 0
		for _, b := range s {
			if cnt[b-'0'] == 0 {
				d++
			}
			cnt[b-'0']++
		}
		if d <= k {
			Fprintf(out, "%s\n", s)
			continue
		}

	o:
		for i := len(s) - 1; ; i-- {
			cnt[s[i]-'0']--
			if cnt[s[i]-'0'] == 0 {
				d--
			}
			for s[i]++; s[i] <= '9'; s[i]++ {
				c := s[i] - '0'
				if cnt[c] == 0 {
					d++
				}
				cnt[c]++
				if d <= k {
					v := 0
					for d == k && cnt[v] == 0 {
						v++
					}
					for j := i + 1; j < len(s); j++ {
						s[j] = '0' + byte(v)
					}
					break o
				}
				cnt[c]--
				if cnt[c] == 0 {
					d--
				}
			}
		}
		Fprintf(out, "%s\n", s)
	}
}

//func main() { cf1560F2(bufio.NewReader(os.Stdin), os.Stdout) }
