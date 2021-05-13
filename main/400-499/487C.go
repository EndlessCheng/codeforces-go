package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF487C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int64
	Fscan(in, &n)
	for i := int64(2); n > 4 && i*i <= n; i++ {
		if n%i == 0 {
			Fprint(out, "NO")
			return
		}
	}
	Fprintln(out, "YES\n1")
	if n == 1 {
		return
	}
	if n == 4 {
		Fprint(out, "3\n2\n4")
		return
	}
	inv := make([]int64, n)
	inv[1] = 1
	for i := int64(2); i < n; i++ {
		Fprintln(out, i*inv[i-1]%n)
		inv[i] = (n - n/i) * inv[n%i] % n
	}
	Fprint(out, n)
}

//func main() { CF487C(os.Stdin, os.Stdout) }
