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
	divisors := [mx][]int{}
	for i := 1; i < mx; i++ {
		for j := i * 2; j < mx; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		pos := [mx][]int{}
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
			l, r := leftHi[i], rightHi[i]
			ans += min(r, rightLo[i]) - i
			for _, d := range divisors[v] {
				ps := pos[d]
				if len(ps) == 0 {
					continue
				}
				firstDi := ps[0]
				if firstDi >= r {
					continue
				}
				lf, rf := leftLo[firstDi], rightLo[firstDi]
				lf = max(lf, l)
				rf = min(rf, r)
				if firstDi > i {
					if lf > i {
						continue
					}
					// 左边 (lf, i] 右边 [firstDi, rf)
					ans += (i - lf) * (rf - firstDi)
				} else if len(ps) == 1 {
					if firstDi <= l || rf < i {
						continue
					}
					// 左边 (lf, firstDi] 右边 [i, rf)
					ans += (firstDi - lf) * (rf - i)
				} else {
					secondDi := ps[1]
					if firstDi <= l && secondDi >= r {
						continue
					}
					ls, rs := leftLo[secondDi], rightLo[secondDi]
					if ls > i {
						// 由 firstDi 负责
						if firstDi > l && rf > i {
							// 左边 (lf, firstDi] 右边 [i, rf)
							ans += (firstDi - lf) * (rf - i)
						}
						continue
					}
					ls = max(ls, l)
					rs = min(rs, r)
					if rf < i {
						// 由 secondDi 负责
						if secondDi < r {
							// 左边 (ls, i] 右边 [secondDi, rs)
							ans += (i - ls) * (rs - secondDi)
						}
						continue
					}
					if firstDi <= l {
						// 左边 (ls, i] 右边 [secondDi, rs)
						ans += (i - ls) * (rs - secondDi)
					} else if secondDi >= r {
						// 左边 (lf, firstDi] 右边 [i, rf)
						ans += (firstDi - lf) * (rf - i)
					} else {
						// 左边 (lf, i] 右边 [i, rs) 减去
						// 左边 (firstDi, i] 右边 [i, secondDi)
						ans += (i-lf)*(rs-i) - (i-firstDi)*(secondDi-i)
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
