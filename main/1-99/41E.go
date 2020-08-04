package __99

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF41E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	Fprintln(out, n*n/4)
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if i&1 != j&1 {
				Fprintln(out, i, j)
			}
		}
	}
}

//func main() { CF41E(os.Stdin, os.Stdout) }
