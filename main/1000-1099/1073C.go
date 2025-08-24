package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf1073C(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	dirs := []struct{ x, y int }{'L': {-1, 0}, 'R': {1, 0}, 'D': {0, -1}, 'U': {0, 1}}
	var n, tx, ty int
	var s string
	Fscan(in, &n, &s, &tx, &ty)
	sd := abs(tx) + abs(ty)
	if n < sd || n%2 != sd%2 {
		Fprint(out, -1)
		return
	}
	sum := make([]struct{ x, y int }, n+1)
	for i, b := range s {
		d := dirs[b]
		sum[i+1].x = sum[i].x + d.x
		sum[i+1].y = sum[i].y + d.y
	}
	ans := sort.Search(n, func(sz int) bool {
		for i := sz; i <= n; i++ {
			if abs(sum[i-sz].x+sum[n].x-sum[i].x-tx)+
				abs(sum[i-sz].y+sum[n].y-sum[i].y-ty) <= sz {
				return true
			}
		}
		return false
	})
	Fprint(out, ans)
}

//func main() { cf1073C(bufio.NewReader(os.Stdin), os.Stdout) }
