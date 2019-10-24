package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
// 题解：https://www.luogu.org/blog/endlesscheng/solution-cf472d
func Sol472D(reader io.Reader, writer io.Writer) {
	in := bufio.NewScanner(reader)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(writer)
	defer out.Flush()
	readInt := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}

	n := readInt()
	d := make([][]int, n)
	for i := range d {
		d[i] = make([]int, n)
		for j := range d[i] {
			d[i][j] = readInt()
			if (i == j) == (d[i][j] != 0) {
				Fprint(out, "NO")
				return
			}
		}
	}
	for i := range d {
		for j := i + 1; j < n; j++ {
			if d[i][j] != d[j][i] {
				Fprint(out, "NO")
				return
			}
		}
	}
	if n <= 2 {
		Fprint(out, "YES")
		return
	}

	type pair struct{ d, idx int }
	ps := make([]pair, n)
	for i, dis := range d[0] {
		ps[i] = pair{dis, i}
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].d < ps[j].d })
	checked := make([]int, 2, n)
	checked[1] = ps[1].idx
	for i := 2; i < n; i++ {
		dis, newIdx := ps[i].d, ps[i].idx
		oldIdx := 0
		minD2 := int(2e9)
		for _, idx := range checked[1:] {
			d2 := d[idx][newIdx]
			if d[0][idx]+d2 == dis && d2 < minD2 { // 离的最近的才能真正成为新加的边
				oldIdx = idx
				minD2 = d2
			}
		}
		for _, idx := range checked {
			if d[idx][oldIdx]+d[oldIdx][newIdx] != d[idx][newIdx] {
				Fprint(out, "NO")
				return
			}
		}
		checked = append(checked, newIdx)
	}
	Fprint(out, "YES")
}

//func main() {
//	Sol472D(os.Stdin, os.Stdout)
//}
