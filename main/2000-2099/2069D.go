package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2069D(in io.Reader, out io.Writer) {
	T, s := 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		l := 0
		for l < n/2 && s[l] == s[n-1-l] {
			l++
		}
		if l == n/2 {
			Fprintln(out, 0)
			continue
		}
		r := n/2 - 1
		for s[r] == s[n-1-r] {
			r--
		}

		var tot, cnt [26]int
		for i := l; i <= r; i++ {
			cnt[s[i]-'a']++
			tot[s[i]-'a']++
			tot[s[n-1-i]-'a']++
		}
		bal := true
		for i, c := range cnt {
			if c*2 != tot[i] {
				bal = false
				break
			}
		}
		if bal {
			Fprintln(out, r-l+1)
			continue
		}

		clear(tot[:])
		for _, v := range s[l : n-l] {
			tot[v-'a']++
		}

		clear(cnt[:])
		rr := n - 1 - l
		for ; ; rr-- {
			v := s[rr] - 'a'
			cnt[v]++
			if cnt[v] > tot[v]/2 {
				break
			}
		}

		clear(cnt[:])
		ll := l
		for ; ; ll++ {
			v := s[ll] - 'a'
			cnt[v]++
			if cnt[v] > tot[v]/2 {
				break
			}
		}
		Fprintln(out, min(rr-l+1, n-l-ll))
	}
}

//func main() { cf2069D(bufio.NewReader(os.Stdin), os.Stdout) }
