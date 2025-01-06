package copypasta

import (
	"fmt"
	"maps"
	"math/bits"
	"slices"
)

/*
我贡献的 LC hack 数据
https://github.com/LeetCode-Feedback/LeetCode-Feedback/issues?q=is%3Aissue+author%3AEndlessCheng+is%3Aclosed

hack 乘法溢出
https://leetcode.cn/problems/frog-position-after-t-seconds/solutions/2281408/dfs-ji-yi-ci-you-qu-de-hack-by-endlessch-jtsr/

用所有的偶数 hack，O(U^logU)
https://leetcode.cn/problems/maximum-total-reward-using-operations-i/solutions/2805422/0-1-bei-bao-by-endlesscheng-702p/comments/2435993

https://yukicoder.me/problems/no/2700

LC2556
https://leetcode.cn/circle/discuss/gXygF3/view/ftWmDG/
考虑这样一个结构
1 1 1
1 0 1
1 1 1
……

*/

// 没找到符合要求的数据
func hackLC3411() {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	lcm := func(a, b int) int { return a / gcd(a, b) * b }
	n := 10
	calc := func(sub int) (l, g int) {
		l = 1
		for _s := uint(sub); _s > 0; _s &= _s - 1 {
			p := bits.TrailingZeros(_s) + 1
			l = lcm(l, p)
			g = gcd(g, p)
		}
		return
	}
	all := map[int]bool{}
	for sub := 1; sub < 1<<n; sub++ {
		l, g := calc(sub)
		all[l*g] = true
	}
	fmt.Println(slices.Sorted(maps.Keys(all)))

	ps := []int{2, 3, 5, 7}
	const mod = 1 << 32
	for k := 1; k <= 100000; k++ {
		for x0 := range all {
			x := x0 + k*mod
			a := []int{}
			for _, p := range ps {
				e := 0
				for x%p == 0 {
					x /= p
					e++
				}
				a = append(a, e)
			}
			if x > 1 {
				continue
			}
			fmt.Println(x0, a)
		}
	}
}
