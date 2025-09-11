package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// https://github.com/EndlessCheng
func cf1194B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		col := make([]int, m)
		a := make([]string, n)
		for i := range a {
			Fscan(in, &a[i])
			for j, b := range a[i] {
				col[j] += int(b >> 2 & 1)
			}
		}

		ans := n + m
		for _, s := range a {
			r := strings.Count(s, ".")
			for j, c := range col {
				ans = min(ans, r+c-int(s[j]>>2&1))
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1194B(bufio.NewReader(os.Stdin), os.Stdout) }
