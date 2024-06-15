package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1715C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, ans, i, x int
	Fscan(in, &n, &m)
	a := make([]int, n+1)
	for i = 1; i <= n; i++ {
		Fscan(in, &a[i])
		k := n + 1 - i
		if a[i] == a[i-1] {
			ans += k
		} else {
			ans += k * i // 增量的贡献
		}
	}
	for ; m > 0; m-- {
		Fscan(in, &i, &x)
		k := n + 1 - i
		if a[i] == a[i-1] && x != a[i-1] {
			ans += k * (i - 1)
		} else if a[i] != a[i-1] && x == a[i-1] {
			ans -= k * (i - 1)
		}
		if i < n {
			k--
			if a[i] == a[i+1] && x != a[i+1] {
				ans += k * i
			} else if a[i] != a[i+1] && x == a[i+1] {
				ans -= k * i
			}
		}
		a[i] = x
		Fprintln(out, ans)
	}
}

//func main() { cf1715C(bufio.NewReader(os.Stdin), os.Stdout) }
