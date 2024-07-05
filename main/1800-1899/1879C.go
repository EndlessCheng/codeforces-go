package main

import (
	. "fmt"
	"io"
)

func cf1879C(in io.Reader, out io.Writer) {
	const mod = 998244353
	T, s := 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		m, ans, size := 0, 1, 0
		for i := range s {
			size++
			if i == len(s)-1 || s[i] != s[i+1] {
				ans = ans * size % mod
				size = 0
			} else {
				m++
				ans = ans * m % mod
			}
		}
		Fprintln(out, m, ans)
	}
}

//func main() { cf1879C(bufio.NewReader(os.Stdin), os.Stdout) }
