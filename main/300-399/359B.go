package _00_399

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF359B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	a := make([]interface{}, 2*n)
	for i := range a {
		a[i] = 2*n - i
	}
	for i := 0; i < k; i++ {
		a[2*i], a[2*i+1] = a[2*i+1], a[2*i]
	}
	Fprint(out, a...)
}

//func main() {
//	CF359B(os.Stdin, os.Stdout)
//}
