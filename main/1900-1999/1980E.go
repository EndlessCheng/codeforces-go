package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1980E(in io.Reader, out io.Writer) {
	var T, n, m int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		f := func() (rows, cols [][]int) {
			rows = make([][]int, n)
			cols = make([][]int, m)
			for i := range cols {
				cols[i] = make([]int, n)
			}
			for i := range rows {
				rows[i] = make([]int, m)
				for j := range rows[i] {
					Fscan(in, &rows[i][j])
					cols[j][i] = rows[i][j]
				}
				slices.Sort(rows[i])
			}
			for _, col := range cols {
				slices.Sort(col)
			}
			cmp := func(a, b []int) int { return slices.Compare(a, b) }
			slices.SortFunc(rows, cmp)
			slices.SortFunc(cols, cmp)
			return
		}

		rows1, cols1 := f()
		rows2, cols2 := f()

		for i, r1 := range rows1 {
			if !slices.Equal(r1, rows2[i]) {
				Fprintln(out, "NO")
				continue o
			}
		}
		for i, c1 := range cols1 {
			if !slices.Equal(c1, cols2[i]) {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { cf1980E(bufio.NewReader(os.Stdin), os.Stdout) }
