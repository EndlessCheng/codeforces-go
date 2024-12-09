package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2047B(in io.Reader, out io.Writer) {
	var T, n int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		cnt := [26]int{}
		for _, b := range s {
			cnt[b-'a']++
		}
		mx, mn := -1, -1
		for i, c := range cnt {
			if c == 0 {
				continue
			}
			if mx < 0 || c > cnt[mx] {
				mx = i
			}
			if mn < 0 || c <= cnt[mn] { // = 可以让 mn 尽量靠后
				mn = i
			}
		}
		for i, b := range s {
			if b == 'a'+byte(mn) {
				s[i] = 'a' + byte(mx)
				break
			}
		}
		Fprintln(out, string(s))
	}
}

//func main() { cf2047B(bufio.NewReader(os.Stdin), os.Stdout) }
