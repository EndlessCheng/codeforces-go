package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1278A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T int
	var s, t []byte
	O := [26]int{}
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &t)
		if len(t) < len(s) {
			Fprintln(out, "NO")
			continue
		}
		cnt := [26]int{}
		for _, b := range s {
			cnt[b-'a']++
		}
		for _, b := range t[:len(s)] {
			cnt[b-'a']--
		}
		if cnt == O {
			Fprintln(out, "YES")
			continue
		}
		for i, b := range t[len(s):] {
			cnt[t[i]-'a']++
			cnt[b-'a']--
			if cnt == O {
				Fprintln(out, "YES")
				continue o
			}
		}
		Fprintln(out, "NO")
	}
}

//func main() { CF1278A(os.Stdin, os.Stdout) }
