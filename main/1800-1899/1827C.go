package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1827C(in io.Reader, out io.Writer) {
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		var ans, boxM, boxR int
		halfLen := make([]int, n-1)
		f := make([]int, n+1)
		st := []int{}
		for i := range halfLen {
			hl := 0
			if i < boxR {
				hl = min(halfLen[boxM*2-i], boxR-i)
			}
			for i >= hl && i+hl+1 < n && s[i-hl] == s[i+hl+1] {
				hl++
				boxM, boxR = i, i+hl
			}
			r := i + 1
			for len(st) > 0 && st[len(st)-1]+halfLen[st[len(st)-1]] < r {
				st = st[:len(st)-1]
			}
			if hl > 0 {
				for len(st) > 0 && halfLen[st[len(st)-1]] <= hl {
					st = st[:len(st)-1]
				}
				st = append(st, i)
				halfLen[i] = hl
			}
			if len(st) > 0 {
				f[r+1] = f[st[len(st)-1]*2+1-r] + 1
			}
			ans += f[r+1]
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1827C(bufio.NewReader(os.Stdin), os.Stdout) }
