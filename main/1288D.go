package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1288D(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	read := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}

	n, m := read(), uint(read())
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, m)
		for j := range mat[i] {
			mat[i][j] = read()
		}
	}
	var ansI, ansJ int
	sort.Search(1e9+1, func(min int) bool {
		id := make([]int, 1<<m)
		for i, mi := range mat {
			v := 0
			for j, a := range mi {
				if a >= min {
					v |= 1 << uint(j)
				}
			}
			id[v] = i + 1
		}
		for i := range id {
			if id[i] > 0 {
				for j := i; j < 1<<m; j++ {
					if i|j == 1<<m-1 && id[j] > 0 {
						ansI, ansJ = id[i], id[j]
						return false
					}
				}
			}
		}
		return true
	})
	Fprint(out, ansI, ansJ)
}

//func main() {
//	CF1288D(os.Stdin, os.Stdout)
//}
