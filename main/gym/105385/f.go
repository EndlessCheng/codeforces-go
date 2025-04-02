package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func runF(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		s := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &s[i])
			s[i] += s[i-1]
		}

		slices.Sort(s[1:n])
		ans := 0
		for _, v := range s[:n] {
			ans += s[n] - v
			Fprint(out, ans, " ")
		}
		Fprintln(out)
	}
}

//func main() { runF(bufio.NewReader(os.Stdin), os.Stdout) }
