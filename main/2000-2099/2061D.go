package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2061D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		cnt := map[int]int{}
		for range n {
			Fscan(in, &v)
			cnt[v]++
		}

		op := n - m
		var f func(int) bool
		f = func(v int) bool {
			if v == 0 || op < 0 {
				return false
			}
			if cnt[v] > 0 {
				cnt[v]--
				return true
			}
			op--
			return f(v/2) && f(v-v/2)
		}

		ok := true
		for range m {
			Fscan(in, &v)
			if ok && !f(v) {
				ok = false
			}
		}

		if ok && op == 0 {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}

//func main() { cf2061D(bufio.NewReader(os.Stdin), os.Stdout) }
