package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF382C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	if n == 1 {
		Fprint(out, -1)
		return
	}

	// 判断 a 是否为等差数列
	isAP := func(a []int) bool {
		for i := 2; i < len(a); i++ {
			if a[i]-a[i-1] != a[1]-a[0] {
				return false
			}
		}
		return true
	}

	ans := []int{}
	sort.Ints(a)
	if isAP(a) {
		if a[0] == a[1] { // 公差为 0
			ans = append(ans, a[0])
		} else {
			ans = append(ans, 2*a[0]-a[1], 2*a[n-1]-a[n-2]) // 最左 / 最右加一个
			if n == 2 && (a[0]+a[1])%2 == 0 {
				ans = append(ans, (a[0]+a[1])/2) // 中间加一个
			}
		}
	} else {
		if (a[0]+a[1])%2 == 0 && isAP(append([]int{a[0], (a[0] + a[1]) / 2}, a[1:]...)) {
			ans = append(ans, (a[0]+a[1])/2) // 中间加一个
		}
		d0 := a[1] - a[0]
		for i := 2; i < n; i++ {
			d := a[i] - a[i-1]
			if d != d0 {
				if d == d0*2 && isAP(append(append(append([]int{}, a[:i]...), (a[i-1]+a[i])/2), a[i:]...)) {
					ans = append(ans, (a[i-1]+a[i])/2) // 中间加一个
				}
				break
			}
		}
	}

	sort.Ints(ans)
	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF382C(os.Stdin, os.Stdout) }
