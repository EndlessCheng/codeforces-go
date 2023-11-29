package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF920C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	n, s := 0, ""
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	Fscan(in, &s)

	for i := 0; i < n-1; i++ {
		if s[i] == '0' {
			continue
		}
		st := i
		for ; i < n-1 && s[i] == '1'; i++ {
		}
		sort.Ints(a[st : i+1])
	}
	if sort.IntsAreSorted(a) {
		Fprint(out, "YES")
	} else {
		Fprint(out, "NO")
	}
}

//func main() { CF920C(os.Stdin, os.Stdout) }
