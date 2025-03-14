package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF446A(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	ans := 1
	f := make([][2]int, n)
	f[0][0] = 1
	f[0][1] = 1
	for i := 1; i < n; i++ {
		if a[i] > a[i-1] {
			f[i][0] = f[i-1][0] + 1
			f[i][1] = f[i-1][1] + 1
		} else {
			f[i][0] = 1
			f[i][1] = 2 // 把 a[i-1] 改成任意 <= a[i]-1 的数
		} 
		if i > 1 && a[i] > a[i-2]+1 { // 把 a[i-1] 改成 [a[i-2]+1, a[i]-1] 中的任意数字
			f[i][1] = max(f[i][1], f[i-2][0]+2) // 在 f[i-2][0] 的基础上多了两个数
		}
		ans = max(ans, f[i][0], f[i][1], f[i-1][0]+1) // 在 f[i-1][0] 的基础上，修改 a[i] 为任意 >= a[i-1]+1 的数
	}
	Fprint(out, ans)
}

//func main() { CF446A(bufio.NewReader(os.Stdin), os.Stdout) }
