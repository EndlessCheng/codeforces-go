package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1217C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		ans, cnt0 := 0, 0
		for i, b := range s {
			if b == '0' {
				cnt0++
				continue
			}
			v := 0
			for j, b := range s[i:] {
				v = v<<1 | int(b&1)
				if v > cnt0+j+1 {
					break
				}
				// 由于 v 最高位为 1，v >= 2^j >= j+1 是恒成立的
				// 因此当 j+1 <= v <= j+1+cnt0 时，一定能添加足够的前导零组成长度为 v 的二进制
				ans++
			}
			cnt0 = 0
		}
		Fprintln(out, ans)
	}
}

// 第一次的做法
func CF1217C_old(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t int
	var s []byte
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &s)
		n, ans := len(s), 0
		pre1 := make([]int, n)
		p := -1
		for i, b := range s {
			pre1[i] = p
			if b == '1' {
				p = i
			}
			s[i] -= '0'
		}
		for i := range s {
			for j, l, v := i, 1, 0; j >= 0; j-- {
				v |= int(s[j]) << (l - 1)
				if v == l {
					ans++
				}
				if l > 20 {
					if l < v && v <= i-pre1[j] {
						ans++
					}
					break
				}
				l++
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1217C(os.Stdin, os.Stdout) }
