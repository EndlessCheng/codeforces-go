package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2033D(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		ans, s := 0, 0
		has := map[int]bool{0: true}
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			s += v
			if has[s] {
				ans++
				has = map[int]bool{}
			}
			has[s] = true
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2033D(bufio.NewReader(os.Stdin), os.Stdout) }
