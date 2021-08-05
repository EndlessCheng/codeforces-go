package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF545C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	if n == 1 {
		Fprint(out, 1)
		return
	}
	a := make([]struct{ x, h int }, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].h)
	}
	ans := 2
	for i := 1; i < n-1; i++ {
		if a[i].x-a[i].h > a[i-1].x { // 左倒
			ans++
		} else if a[i].x+a[i].h < a[i+1].x { // 右倒
			a[i].x += a[i].h
			ans++
		}
	}
	Fprint(out, ans)
}

//func main() { CF545C(os.Stdin, os.Stdout) }
