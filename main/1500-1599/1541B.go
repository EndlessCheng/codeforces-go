package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1541B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, ai int
	for Fscan(in, &T); T > 0; T-- {
		ans := 0
		Fscan(in, &n)
		pos := make([]int, n*2+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &ai)
			for aj := 1; ai*aj < i*2; aj++ {
				if j := pos[aj]; j > 0 && ai*aj == i+j {
					ans++
				}
			}
			pos[ai] = i
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1541B(os.Stdin, os.Stdout) }
