package _00_399

import (
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF322B(_r io.Reader, _w io.Writer) {
	a := make([]int, 3)
	ans := 0
	has3k := false
	for i := range a {
		Fscan(_r, &a[i])
		if a[i] > 0 && a[i]%3 == 0 {
			has3k = true
		}
		ans += a[i] / 3
		a[i] %= 3
	}
	sort.Ints(a)
	if has3k && a[0] == 0 && a[1] == 2 && a[2] == 2 {
		ans++
	} else {
		ans += a[0]
	}
	Fprint(_w, ans)
}

//func main() { CF322B(os.Stdin, os.Stdout) }
