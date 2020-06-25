package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1370B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var t, n, v int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		id := [2][]int{}
		for i := 1; i <= 2*n; i++ {
			Fscan(in, &v)
			id[v&1] = append(id[v&1], i)
		}
		if len(id[0])&1 == 0 {
			if len(id[0]) > 0 {
				id[0] = id[0][2:]
			} else {
				id[1] = id[1][2:]
			}
		} else {
			id[0] = id[0][1:]
			id[1] = id[1][1:]
		}
		for _, v := range id {
			for i := 0; i < len(v); i += 2 {
				Fprintln(out, v[i], v[i+1])
			}
		}
	}
}

//func main() { CF1370B(os.Stdin, os.Stdout) }
