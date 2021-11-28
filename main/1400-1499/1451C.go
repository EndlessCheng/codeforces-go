package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1451C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	var s string
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s)
		c := [26]int{}
		for _, b := range s {
			c[b-'a']++
		}
		Fscan(in, &s)
		c2 := [26]int{}
		for _, b := range s {
			c2[b-'a']++
		}
		tot := 0
		for i := range c {
			tot += c[i] - c2[i]
			if tot < 0 || tot%k > 0 {
				Fprintln(out, "No")
				continue o
			}
		}
		Fprintln(out, "Yes")
	}
}

//func main() { CF1451C(os.Stdin, os.Stdout) }
