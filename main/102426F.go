package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF102426F(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	r := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}

	n, m, v := r(), r(), r()
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, m)
		for j := range mat[i] {
			mat[i][j] = r()
		}
	}
	Fprint(_w, sort.Search(100, func(H int) bool {
		sum := 0
		for _, row := range mat {
			for _, h := range row {
				if h > H {
					sum += h - H
				}
			}
		}
		return sum <= v
	}))
}

//func main() { CF102426F(os.Stdin, os.Stdout) }
