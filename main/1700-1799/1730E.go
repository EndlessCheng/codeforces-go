package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1730E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mx = 1000001
	divisors := [mx][]uint32{}
	for i := uint32(1); i < mx; i++ {
		for j := i * 2; j < mx; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}

	var T, n int
	pos := [mx][]int{}
	for Fscan(in, &T); T > 0; T-- {
		for i := range pos {
			pos[i] = pos[i][:0]
		}
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			pos[a[i]] = append(pos[a[i]], i)
		}

		leftHi := make([]int, n)  // >= a[i]
		rightHi := make([]int, n) // > a[i]
		leftLo := make([]int, n)  // < a[i]
		rightLo := make([]int, n) // < a[i]
		s := []int{-1}
		t := []int{-1}
		for i, v := range a {
			for len(s) > 1 && v > a[s[len(s)-1]] {
				rightHi[s[len(s)-1]] = i
				s = s[:len(s)-1]
			}
			leftHi[i] = s[len(s)-1]
			s = append(s, i)

			for len(t) > 1 && v <= a[t[len(t)-1]] {
				t = t[:len(t)-1]
			}
			leftLo[i] = t[len(t)-1]
			t = append(t, i)
		}
		for _, i := range s[1:] {
			rightHi[i] = n
		}

		t = append(t[:0], n)
		for i := n - 1; i >= 0; i-- {
			v := a[i]
			for len(t) > 1 && v <= a[t[len(t)-1]] {
				t = t[:len(t)-1]
			}
			rightLo[i] = t[len(t)-1]
			t = append(t, i)
		}

		ans := 0
		for i, v := range a {
			r := rightHi[i]
			ans += min(rightLo[i], r) - i // 全为 v 的子数组个数
			for _, d := range divisors[v] {
				ps := pos[d]
				l := leftHi[i]
				if len(ps) > 0 && ps[0] < i {
					j := ps[0]
					ps = ps[1:]
					if j > l && rightLo[j] > i {
						ans += (j - max(leftLo[j], l)) * (min(rightLo[j], r) - i)
					}
					l = max(l, j) // 避免重复统计
				}
				if len(ps) > 0 {
					j := ps[0]
					if j < r && leftLo[j] < i {
						ans += (i - max(leftLo[j], l)) * (min(rightLo[j], r) - j)
					}
				}
			}
			if len(pos[v]) > 1 && pos[v][1] == i {
				pos[v] = pos[v][1:]
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1730E(os.Stdin, os.Stdout) }
