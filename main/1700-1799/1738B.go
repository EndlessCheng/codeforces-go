package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1738B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		s := make([]int, k)
		for i := range s {
			Fscan(in, &s[i])
		}
		if k > 1 && s[0] > (n-k+1)*(s[1]-s[0]) {
			Fprintln(out, "No")
			continue
		}
		for i := 2; i < k; i++ {
			if s[i-1]*2 > s[i]+s[i-2] {
				Fprintln(out, "No")
				continue o
			}
		}
		Fprintln(out, "Yes")
	}
}

//func main() { cf1738B(os.Stdin, os.Stdout) }
