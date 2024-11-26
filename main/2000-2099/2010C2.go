package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2010C2(in io.Reader, out io.Writer) {
	var s string
	Fscan(in, &s)
	n := len(s)
	z := make([]int, n)
	boxL, boxR := 0, 0
	for i := 1; i < (n+1)/2; i++ {
		if i <= boxR {
			z[i] = min(z[i-boxL], boxR-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			boxL, boxR = i, i+z[i]
			z[i]++
		}
		if i+z[i] >= n {
			Fprintln(out, "YES", s[i:])
			return
		}
	}
	Fprint(out, "NO")
}

//func main() { cf2010C2(bufio.NewReader(os.Stdin), os.Stdout) }
