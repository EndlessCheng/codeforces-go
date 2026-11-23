package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2289C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n+1)
		st := []int{-1}
		ans := 0
		for r := range a {
			if r < n {
				Fscan(in, &a[r])
			}
			for len(st) > 1 && a[st[len(st)-1]] > a[r] {
				i := st[len(st)-1]
				st = st[:len(st)-1]
				l := st[len(st)-1]
				if i-l > 1 && r-i > 1 {
					ans += 2
				} else {
					ans += r - l - 1
				}
			}
			st = append(st, r)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2289C(bufio.NewReader(os.Stdin), os.Stdout) }
