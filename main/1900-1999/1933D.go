package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1933D(in io.Reader, out io.Writer) {
	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		mn, cnt := int(2e9), 0
		for i := range a {
			Fscan(in, &a[i])
			if a[i] < mn {
				mn, cnt = a[i], 1
			} else if a[i] == mn {
				cnt++
			}
		}
		if cnt == 1 {
			Fprintln(out, "YES")
			continue
		}
		for _, v := range a {
			if v%mn > 0 {
				Fprintln(out, "YES")
				continue o
			}
		}
		Fprintln(out, "NO")
	}
}

//func main() { cf1933D(bufio.NewReader(os.Stdin), os.Stdout) }
