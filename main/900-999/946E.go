package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"strings"
)

// https://github.com/EndlessCheng
func cf946E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	T, s := 0, []byte{}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		if s[0] == '1' && s[n-1] <= '1' && bytes.Count(s[1:n-1], []byte{'0'}) == n-2 {
			Fprintln(out, strings.Repeat("9", n-2))
			continue
		}

		for i := n - 1; i >= 0; i-- {
			if s[i] > '0' {
				s[i]--
				for i++; i < n; i++ {
					s[i] = '9'
				}
				break
			}
		}

		cnt := [10]int{}
		for _, b := range s {
			cnt[b-'0'] ^= 1
		}
		odd := 0
		for _, c := range cnt[:] {
			odd += c
		}

		if odd > 0 {
		o:
			for i := n - 1; i >= 0; i-- {
				cnt[s[i]-'0'] ^= 1
				odd += cnt[s[i]-'0']*2 - 1
				for s[i]--; s[i] >= '0'; s[i]-- {
					v := s[i] - '0'
					cnt[v] ^= 1
					odd += cnt[v]*2 - 1
					if odd <= n-1-i {
						odd -= cnt[9]
						for j := i + 1; j < n-odd; j++ {
							s[j] = '9'
						}
						j := n - odd
						for k := 8; k >= 0; k-- {
							if cnt[k] > 0 {
								s[j] = '0' + byte(k)
								j++
							}
						}
						break o
					}
					odd -= cnt[v]*2 - 1
					cnt[v] ^= 1
				}
			}
		}
		Fprintf(out, "%s\n", s)
	}
}

//func main() { cf946E(bufio.NewReader(os.Stdin), os.Stdout) }
