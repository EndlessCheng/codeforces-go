package main

import "math/rand"

/* 中位数

要使任意两元素最终相等，这两个元素之差必须是 $x$ 的倍数，否则无法通过加减 $x$ 来相等。

假设要让所有元素均为 $y$，可以发现：
- $y$ 每增加 $x$，小于或等于 $y$ 的元素要多操作一次，大于 $y$ 的元素要少操作一次；
- $y$ 每减小 $x$，大于或等于 $y$ 的元素要多操作一次，小于 $y$ 的元素要少操作一次。

因此 $y$ 选在所有元素的中位数上是最「均衡」的。

*/

// github.com/EndlessCheng/codeforces-go
func minOperations(g [][]int, x int) (ans int) {
	n := len(g) * len(g[0])
	for _, r := range g {
		for _, v := range r {
			if (v-g[0][0])%x != 0 {
				return -1
			}
		}
	}
	a := make([]int, 0, n)
	for _, r := range g {
		a = append(a, r...)
	}
	quickSelect(a)
	for _, v := range a {
		ans += abs(v-a[n/2]) / x
	}
	return
}

func quickSelect(a []int) int {
	k := len(a) / 2
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	for l, r := 0, len(a)-1; l < r; {
		v := a[l]
		i, j := l, r+1
		for {
			for i++; i < r && a[i] < v; i++ {
			}
			for j--; j > l && a[j] > v; j-- {
			}
			if i >= j {
				break
			}
			a[i], a[j] = a[j], a[i]
		}
		a[l], a[j] = a[j], v
		if j == k {
			break
		} else if j < k {
			l = j + 1
		} else {
			r = j - 1
		}
	}
	return a[k]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
