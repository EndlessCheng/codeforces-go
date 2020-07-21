package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF999D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ i, v int }

	var n, m int
	Fscan(in, &n, &m)
	sz := n / m
	a := make([]int, n)
	id := make([][]int, m)
	for i := range a {
		Fscan(in, &a[i])
		id[a[i]%m] = append(id[a[i]%m], i)
	}

	ans := int64(0)
	var todo []pair
	for v := 0; v < 2*m; v++ {
		i := v % m
		for len(id[i]) > sz {
			j := id[i][0]
			id[i] = id[i][1:]
			todo = append(todo, pair{j, v})
		}
		for len(todo) > 0 && len(id[i]) < sz {
			p := todo[0]
			todo = todo[1:]
			id[i] = append(id[i], p.i)
			d := v - p.v
			a[p.i] += d
			ans += int64(d)
		}
	}
	Fprintln(out, ans)
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

//func main() { CF999D(os.Stdin, os.Stdout) }
