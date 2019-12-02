package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

type vec552 struct {
	x, y int
}

func (a vec552) sub(b vec552) vec552   { return vec552{a.x - b.x, a.y - b.y} }
func (a vec552) cross(b vec552) int { return a.x*b.y - a.y*b.x }
func (a vec552) reverse() vec552    { return vec552{-a.x, -a.y} }
func (a vec552) up() vec552 {
	if a.y < 0 || a.y == 0 && a.x < 0 {
		return a.reverse()
	}
	return a
}

// github.com/EndlessCheng/codeforces-go
func Sol552D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	if n < 3 {
		Fprint(out, 0)
		return
	}
	ps := make([]vec552, n)
	for i := range ps {
		Fscan(in, &ps[i].x, &ps[i].y)
	}
	ans := int64(n) * int64(n-1) * int64(n-2) / 6
	for i, pi := range ps {
		ls := make([]vec552, 0, n-i-1)
		for j := i + 1; j < n; j++ {
			ls = append(ls, ps[j].sub(pi).up())
		}
		sort.Slice(ls, func(i, j int) bool { return ls[i].cross(ls[j]) > 0 })
		for j := 0; j < len(ls); {
			j0 := j
			for j++; j < len(ls) && ls[j].cross(ls[j0]) == 0; j++ {
			}
			ans -= int64(j-j0) * int64(j-j0-1) / 2
		}
	}
	Fprint(out, ans)
}

//func main() {
//	Sol552D(os.Stdin, os.Stdout)
//}
