package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// O(n) 做法
func CF432D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	Fscan(in, &s)
	n := len(s)
	z := make([]int, n)
	cnt := make([]int, n+2)
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(z[i-l], r-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
		cnt[z[i]]++
	}
	z[0] = n
	cnt[n]++

	ans := [][2]int{}
	for i := n; i > 0; i-- {
		cnt[i] += cnt[i+1]
		if z[n-i] == i {
			ans = append(ans, [2]int{i, cnt[i]})
		}
	}
	Fprintln(out, len(ans))
	for i := len(ans) - 1; i >= 0; i-- {
		p := ans[i]
		Fprintln(out, p[0], p[1])
	}
}

// O(nlogn) 做法
func CF432D2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	Fscan(in, &s)
	n := len(s)
	z := make([]int, n)
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(z[i-l], r-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
	}
	z[0] = n

	sortedZ := append([]int{}, z...)
	sort.Ints(sortedZ)

	ans := [][2]int{}
	for i := 1; i <= n; i++ {
		if z[n-i] == i {
			j := sort.SearchInts(sortedZ, i)
			ans = append(ans, [2]int{i, n - j})
		}
	}
	Fprintln(out, len(ans))
	for _, p := range ans {
		Fprintln(out, p[0], p[1])
	}
}

//func main() { CF432D(os.Stdin, os.Stdout) }
