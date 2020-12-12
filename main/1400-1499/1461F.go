package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1461F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	output := func(a []byte, op string) {
		Fprint(out, a[0])
		for _, v := range a[1:] {
			Fprint(out, op, v)
		}
	}

	var n, st int
	Fscan(in, &n)
	a := make([]byte, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	var s string
	Fscan(in, &s)
	if s == "+-" || s == "-+" {
		s = "+"
	}

	if len(s) == 1 {
		output(a, s)
		return
	}

	if s == "*-" || s == "-*" {
		posZ := bytes.IndexByte(a, 0)
		if posZ <= 0 {
			output(a, "*")
		} else {
			output(a[:posZ], "*")
			Fprint(out, "-")
			output(a[posZ:], "*")
		}
		return
	}

	pos := []int{}   // 保存大于 1 的数的位置
	a = append(a, 0) // 方便处理没有 0 的情况
	// 按 0 分割成若干段，分别处理
	for i, v := range a {
		if v > 1 {
			pos = append(pos, i)
			continue
		}
		if v == 1 {
			if len(pos) == 0 { // 一段的开头
				if i > 0 {
					Fprint(out, "+")
				}
				Fprint(out, 1)
			}
			continue
		}
		if m := len(pos); m > 0 {
			addPos := 0
			// 粗略的比较：若乘积的最小值（将所有大于 1 的数视作 2）比加法的最大值大，就说明没有取加号的机会，否则二进制枚举，看看加号放哪能得到最大值
			if m < 20 && 1<<m <= 9*(i-st) {
				max := int64(0)
				for sub := 0; sub < 1<<(m-1); sub++ {
					sum, mul := int64(0), int64(a[pos[0]])
					for i := 1; i < m; i++ {
						if sub>>(i-1)&1 > 0 {
							sum += mul + int64(pos[i]-pos[i-1]-1)
							mul = int64(a[pos[i]])
						} else {
							mul *= int64(a[pos[i]])
						}
						if sum+mul > max {
							max, addPos = sum+mul, sub
						}
					}
				}
			}
			if pos[0] > 0 {
				Fprint(out, "+")
			}
			Fprint(out, a[pos[0]])
			for i := 1; i < m; i++ {
				op := "*"
				if addPos>>(i-1)&1 > 0 {
					op = "+"
				}
				for _, v := range a[pos[i-1]+1 : pos[i]+1] {
					Fprint(out, op, v)
				}
			}
			for _, v := range a[pos[m-1]+1 : i] {
				Fprint(out, "+", v)
			}
		}
		if i < n {
			if i > 0 {
				Fprint(out, "+")
			}
			Fprint(out, 0)
		}
		pos = nil
		st = i + 1
	}
}

//func main() { CF1461F(os.Stdin, os.Stdout) }
