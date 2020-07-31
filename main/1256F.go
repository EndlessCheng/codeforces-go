package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1256F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s, t []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s, &t)
		same := false
		cnt := [26]int{}
		for _, b := range s {
			if cnt[b-'a']++; cnt[b-'a'] > 1 {
				same = true
			}
		}
		for _, b := range t {
			cnt[b-'a']--
		}
		if cnt != [26]int{} {
			Fprintln(out, "NO")
		} else if same {
			Fprintln(out, "YES")
		} else {
			c := 0
			for i, b := range s {
				for _, b2 := range s[:i] {
					if b2 > b {
						c ^= 1
					}
				}
			}
			for i, b := range t {
				for _, b2 := range t[:i] {
					if b2 > b {
						c ^= 1
					}
				}
			}
			if c > 0 {
				Fprintln(out, "NO")
			} else {
				Fprintln(out, "YES")
			}
		}
	}
}

//func main() { CF1256F(os.Stdin, os.Stdout) }
