package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1648B(in io.Reader, out io.Writer) {
	var T, n, mx, v int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &mx)
		s := make([]int, mx+1)
		for range n {
			Fscan(in, &v)
			s[v] = 1
		}
		for i := 2; i <= mx; i++ {
			s[i] += s[i-1]
		}
		if s[1] == 0 {
			Fprintln(out, "No")
			continue
		}
		for i := 2; i <= mx/2; i++ {
			if s[i] == s[i-1] {
				continue
			}
			for j := i; j <= mx; j += i {
				if s[j-1] < s[min(j+i, mx)] && s[j/i] == s[j/i-1] {
					Fprintln(out, "No")
					continue o
				}
			}
		}
		Fprintln(out, "Yes")
	}
}

//func main() { cf1648B(bufio.NewReader(os.Stdin), os.Stdout) }
