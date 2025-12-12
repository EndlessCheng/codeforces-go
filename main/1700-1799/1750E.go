package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1750E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		sum := make([]int, n+1)
		for i, b := range s {
			sum[i+1] = sum[i] + 1 - int(b%2*2)
		}

		l := make([]int, n+1)
		r := make([]int, n+1)
		st := []int{-1}
		for i, v := range sum {
			for len(st) > 1 && sum[st[len(st)-1]] > v {
				r[st[len(st)-1]] = i
				st = st[:len(st)-1]
			}
			l[i] = st[len(st)-1]
			st = append(st, i)
		}
		for _, i := range st[1:] {
			r[i] = n + 1
		}

		ans := 0
		b := slices.Clone(sum)
		slices.Sort(b)
		for i, v := range b {
			ans += v*i - sum[i]*((i-l[i])*(r[i]-i)-1)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1750E(bufio.NewReader(os.Stdin), os.Stdout) }
