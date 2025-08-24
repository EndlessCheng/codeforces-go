package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1194C(in io.Reader, out io.Writer) {
	var T int
	var s, t, p []byte
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &t, &p)
		cnt := [26]int{}
		j := 0
		for i, b := range t {
			if s[j] != b {
				cnt[b-'a']++
				continue
			}
			j++
			if j == len(s) {
				for _, b := range t[i+1:] {
					cnt[b-'a']++
				}
				for _, b := range p {
					cnt[b-'a']--
				}
				for _, c := range cnt {
					if c > 0 {
						Fprintln(out, "NO")
						continue o
					}
				}
				Fprintln(out, "YES")
				continue o
			}
		}
		Fprintln(out, "NO")
	}
}

//func main() { cf1194C(bufio.NewReader(os.Stdin), os.Stdout) }
