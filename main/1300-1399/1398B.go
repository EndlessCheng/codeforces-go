package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1398B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	T, s := 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		a := strings.Split(s, "0")
		sort.Strings(a)
		ans := 0
		for i := len(a) - 1; i >= 0; i -= 2 {
			ans += len(a[i])
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1398B(os.Stdin, os.Stdout) }
