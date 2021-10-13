package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1515D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var T, n, l, r, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &l, &r)
		c := make([]int, n+1)
		for i := 0; i < l; i++ {
			Fscan(in, &v)
			c[v]++
		}
		for i := 0; i < r; i++ {
			Fscan(in, &v)
			c[v]--
		}
		ans := 0
		for i := 1; i <= n; i++ {
			for l > r && c[i] > 1 { // 左袜子多，且有相同颜色，左变右，与自己匹配
				l -= 2
				c[i] -= 2
				ans++
			}
			for l < r && c[i] < -1 { // 右袜子多，且有相同颜色，右变左，与自己匹配
				r -= 2
				c[i] += 2
				ans++
			}
		}
		ans += abs(l-r) / 2 // 左变右或右变左
		left := 0
		for _, c := range c {
			left += abs(c) // 变更颜色
		}
		Fprintln(out, ans+left/2) // 每次变更颜色都会配对两只袜子，所以要除以 2
	}
}

//func main() { CF1515D(os.Stdin, os.Stdout) }
