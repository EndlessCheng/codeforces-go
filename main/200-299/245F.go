package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
	"time"
)

// github.com/EndlessCheng/codeforces-go
func CF245F(_r io.Reader, out io.Writer) {
	in := bufio.NewScanner(_r)
	var n, m int
	in.Scan()
	Sscan(in.Text(), &n, &m)
	a := []int{}
	for in.Scan() {
		s := in.Text()[:19]
		t, _ := time.Parse("2006-01-02 15:04:05", s)
		v := int(t.Unix())
		a = append(a, v)
		if len(a)-sort.SearchInts(a, v-n+1) >= m {
			Fprint(out, s)
			return
		}
	}
	Fprint(out, -1)
}

//func main() { CF245F(os.Stdin, os.Stdout) }
