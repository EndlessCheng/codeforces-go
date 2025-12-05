package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1392F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, s, v int
	Fscan(in, &n)
	for i := range n {
		Fscan(in, &v)
		s += v - i
	}
	for i := range n {
		ans := s/n + i
		if i < s%n {
			ans++
		}
		Fprint(out, ans, " ")
	}
}

//func main() { cf1392F(bufio.NewReader(os.Stdin), os.Stdout) }
