package main

import (
	. "fmt"
	"io"
)

func cf1937B(in io.Reader, out io.Writer) {
	var T, n int
	var s, t string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s, &t)
		i := 1
		for i < n && s[i] <= t[i-1] {
			i++
		}
		j := i - 1
		for j > 0 && s[j] == t[j-1] {
			j--
		}
		Fprintln(out, s[:i]+t[i-1:])
		Fprintln(out, i-j)
	}
}

//func main() { cf1937B(bufio.NewReader(os.Stdin), os.Stdout) }
