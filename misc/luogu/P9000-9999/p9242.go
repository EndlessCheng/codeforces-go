package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func p9242(in io.Reader, out io.Writer) {
	n, s := 0, ""
	Fscan(in, &n)
	f := [10]int{}
	for range n {
		Fscan(in, &s)
		head, tail := s[0]-'0', s[len(s)-1]-'0'
		f[tail] = max(f[tail], f[head]+1)
	}
	Fprint(out, n-slices.Max(f[:]))
}

//func main() { p9242(bufio.NewReader(os.Stdin), os.Stdout) }
