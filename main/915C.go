package main

import (
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF915C(in io.Reader, out io.Writer) {
	var a []byte
	var upp string
	Fscan(in, &a, &upp)
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
	ans := ""
	for len(a) > 0 {
		for i, b := range a {
			aa := append(append([]byte{}, a[:i]...), a[i+1:]...)
			tmp := make([]byte, len(aa))
			copy(tmp, aa)
			sort.Slice(aa, func(i, j int) bool { return aa[i] < aa[j] })
			if s := ans + string(b) + string(aa); len(s) < len(upp) || s <= upp {
				ans += string(b)
				a = tmp
				break
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF915C(os.Stdin, os.Stdout) }
