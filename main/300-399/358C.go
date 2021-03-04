package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF358C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	op := [3]string{"pushStack", "pushQueue", "pushFront"}

	var n, v int
	type pair struct{ v, i int }
	a := []pair{}
	for Fscan(in, &n); n > 0; n-- {
		if Fscan(in, &v); v > 0 {
			a = append(a, pair{v, len(a)})
			continue
		}
		if len(a) == 0 {
			Fprintln(out, 0)
		} else if len(a) == 1 {
			Fprintln(out, `pushStack
1 popStack`)
		} else if len(a) == 2 {
			Fprintln(out, `pushStack
pushQueue
2 popStack popQueue`)
		} else {
			b := append([]pair(nil), a...)
			sort.Slice(b, func(i, j int) bool { return b[i].v > b[j].v })
			c := 0
			for i := range a {
				if i == b[0].i || i == b[1].i || i == b[2].i {
					Fprintln(out, op[c])
					c++
				} else {
					Fprintln(out, "pushBack")
				}
			}
			Fprintln(out, "3 popStack popQueue popFront")
		}
		a = nil
	}
	for range a { // WA 了一次，没判断 a 还有数据的情况
		Fprintln(out, op[0])
	}
}

//func main() { CF358C(os.Stdin, os.Stdout) }
