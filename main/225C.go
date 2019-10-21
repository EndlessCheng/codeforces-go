package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func Sol225C(reader io.Reader, writer io.Writer) {
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, m, x, y int
	Fscan(in, &n, &m, &x, &y)
	cnt := make([]int, m)
	pic := make([]string, n)
	for i := range pic {
		Fscan(in, &pic[i])
		for j, c := range pic[i] {
			if c == '.' {
				cnt[j]++
			}
		}
	}

	cache := map[string]int{}
	var f func(int, int8, int) int
	f = func(i int, isDot int8, con int) (res int) {
		if i == m {
			return
		}
		if con > y {
			return 1e8
		}
		hash := Sprintf("%d;%d;%d", i, isDot, con)
		if val, ok := cache[hash]; ok {
			return val
		}
		if isDot == 0 { // #
			res = f(i+1, isDot, con+1) + cnt[i]
			if con >= x {
				res = min(res, f(i+1, isDot^1, 1)+n-cnt[i]) // .
			}
		} else { // .
			res = f(i+1, isDot, con+1) + n - cnt[i]
			if con >= x {
				res = min(res, f(i+1, isDot^1, 1)+cnt[i]) // #
			}
		}
		cache[hash] = res
		return
	}
	Fprint(out, min(f(1, 0, 1)+cnt[0], f(1, 1, 1)+n-cnt[0]))
}

func main() {
	Sol225C(os.Stdin, os.Stdout)
}
