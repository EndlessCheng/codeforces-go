package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1566C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s0, s1 []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s0, &s1)
		ans := 0
		use := -1
		for i, b := range s0 {
			if b != s1[i] {
				ans += 2
			} else if b == '0' || i > 0 && s0[i-1] == '0' && s1[i-1] == '0' && use != i-1 {
				ans++
			} else if i < n-1 && s0[i+1] == '0' && s1[i+1] == '0' {
				ans++
				use = i + 1
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1566C(os.Stdin, os.Stdout) }
