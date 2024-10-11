package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf1795C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		s := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &s[i])
			s[i] += s[i-1]
		}

		d := make([]int, n+1)
		ex := make([]int, n+1)
		sd := 0
		for i, v := range a {
			j := i + sort.SearchInts(s[i:], s[i]+v+1) - 1
			d[i]++
			d[j]--
			sd += d[i]
			ex[j] += v - (s[j] - s[i])
			Fprint(out, sd*(s[i+1]-s[i])+ex[i], " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1795C(bufio.NewReader(os.Stdin), os.Stdout) }
