package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF484B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mx int = 1e6
	a := [mx + 1]bool{}
	var n, v, ans int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		a[v] = true
	}
	pre := [mx + 1]int{}
	for i := 2; i <= mx; i++ {
		if a[i] {
			pre[i] = i
		} else {
			pre[i] = pre[i-1]
		}
	}
	for i := 2; i <= mx; i++ {
		if a[i] {
			for j := 2 * i; j <= mx; j += i {
				if x := pre[j-1] % i; x > ans {
					ans = x
				}
			}
			if pre[mx]%i > ans {
				ans = pre[mx] % i
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF484B(os.Stdin, os.Stdout) }
