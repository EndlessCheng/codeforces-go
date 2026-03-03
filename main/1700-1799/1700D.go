package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf1700D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, t int
	Fscan(in, &n)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		s[i] += s[i-1]
	}

	time := make([]int, n)
	preMax := 0
	for i := 1; i <= n; i++ {
		preMax = max(preMax, (s[i]-1)/i+1)
		time[i-1] = max(preMax, (s[n]-1)/i+1)
	}

	Fscan(in, &q)
	for range q {
		Fscan(in, &t)
		i := sort.Search(n, func(i int) bool { return time[i] <= t })
		if i == n {
			Fprintln(out, -1)
		} else {
			Fprintln(out, i+1)
		}
	}
}

//func main() { cf1700D(bufio.NewReader(os.Stdin), os.Stdout) }
