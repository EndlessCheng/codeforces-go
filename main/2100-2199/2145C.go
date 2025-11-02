package main

import (
	. "fmt"
	"io"
	"strings"
)

// https://github.com/EndlessCheng
func cf2145C(in io.Reader, out io.Writer) {
	T, n, s := 0, 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		d := strings.Count(s, "a")*2 - n
		pos := map[int]int{0: -1}
		ans, sum := n, 0
		for i, b := range s {
			sum += 1 - int(b-'a')*2
			pos[sum] = i
			if j, ok := pos[sum-d]; ok {
				ans = min(ans, i-j)
			}
		}
		if ans == n {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2145C(bufio.NewReader(os.Stdin), os.Stdout) }
