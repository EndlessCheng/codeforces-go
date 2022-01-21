package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF814B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	diff := []int{}
	vis := make([]bool, n+1)
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
		if b[i] != a[i] {
			diff = append(diff, i)
		} else {
			vis[a[i]] = true
		}
	}
	i := 1
	for vis[i] {
		i++
	}
	if len(diff) == 1 {
		a[diff[0]] = i
	} else {
		j := i + 1
		for vis[j] {
			j++
		}
		if a[diff[0]] == i != (a[diff[1]] == j) && b[diff[0]] == i != (b[diff[1]] == j) {
			a[diff[0]] = i
			a[diff[1]] = j
		} else {
			a[diff[0]] = j
			a[diff[1]] = i
		}
	}
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

//func main() { CF814B(os.Stdin, os.Stdout) }
