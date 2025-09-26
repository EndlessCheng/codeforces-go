package main

import (
	. "fmt"
	"io"
)

// (n-1)*n*(n+1)/6 的两种组合视角 https://chatgpt.com/c/68d4e480-0634-8324-ad29-15877e539f81

// https://github.com/EndlessCheng
func cf1827B2(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		ll := make([]int, n)
		l := make([]int, n)
		r := make([]int, n)
		for i := range n {
			ll[i] = -1
			l[i] = -1
			r[i] = n
		}
		var st, st2 []int
		for i := n - 1; i >= 0; i-- {
			for len(st2) > 0 && a[st2[len(st2)-1]] < a[i] {
				ll[st2[len(st2)-1]] = i
				st2 = st2[:len(st2)-1]
			}
			for len(st) > 0 && a[st[len(st)-1]] > a[i] {
				j := st[len(st)-1]
				st = st[:len(st)-1]
				st2 = append(st2, j)
				l[j] = i
			}
			if len(st) > 0 {
				r[i] = st[len(st)-1]
			}
			st = append(st, i)
		}

		ans := (n - 1) * n * (n + 1) / 6
		for i := range n {
			ans -= (l[i] - ll[i]) * (r[i] - i)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1827B2(bufio.NewReader(os.Stdin), os.Stdout) }
