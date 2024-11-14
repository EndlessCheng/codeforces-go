package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1736C1(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := n * (n + 1) / 2
		l := 0
		for i := range n {
			Fscan(in, &v)
			v -= i + 1
			for l < i && -l > v {
				ans -= n - i
				l++
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1736C1(bufio.NewReader(os.Stdin), os.Stdout) }
