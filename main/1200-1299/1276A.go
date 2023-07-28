package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1276A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		ans := []int{}
		for i := 0; i+2 < len(s); i++ {
			if i+4 < len(s) && s[i:i+5] == "twone" {
				ans = append(ans, i+3)
				i += 4
			} else if s[i:i+3] == "one" || s[i:i+3] == "two" {
				ans = append(ans, i+2)
				i += 2
			}
		}
		Fprintln(out, len(ans))
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1276A(os.Stdin, os.Stdout) }
