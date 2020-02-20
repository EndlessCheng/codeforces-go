package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF651C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x, y int
	Fscan(in, &n)
	rows, cols := map[int]int{}, map[int]int{}
	cnts := map[[2]int]int{}
	ans := int64(0)
	for ; n > 0; n-- {
		Fscan(in, &x, &y)
		p := [2]int{x, y}
		ans += int64(rows[y] + cols[x] - cnts[p])
		rows[y]++
		cols[x]++
		cnts[p]++
	}
	Fprint(out, ans)
}

//func main() { CF651C(os.Stdin, os.Stdout) }
