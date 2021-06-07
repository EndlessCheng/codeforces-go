package copypasta

/* 快速沃尔什变换 fast Walsh–Hadamard transform, FWT, FWHT
在算法竞赛中，FWT 是用于解决对下标进行【位运算卷积】问题的方法
https://en.wikipedia.org/wiki/Fast_Walsh%E2%80%93Hadamard_transform
https://oi-wiki.org/math/poly/fwt/
https://www.luogu.com.cn/blog/CJL/fwt-xue-xi-bi-ji
https://www.luogu.com.cn/blog/command-block/wei-yun-suan-juan-ji-yu-ji-kuo-zhan
一些习题 https://blog.csdn.net/weixin_38686780/article/details/81912853

模板题 https://www.luogu.com.cn/problem/P4717
LC1803/周赛233D https://leetcode-cn.com/problems/count-pairs-with-xor-in-a-range/
与仙人掌结合 https://codeforces.com/problemset/problem/1218/D
todo https://codeforces.com/problemset/problem/662/C

EXTRA: 子集卷积
todo https://www.cnblogs.com/yijan/p/12387352.html
 https://www.luogu.com.cn/problem/P6097
*/

// 取模的写法见 https://www.luogu.com.cn/record/51397587
func fwtOR(a []int, op int) []int {
	n := len(a)
	for l, k := 2, 1; l <= n; l, k = l<<1, k<<1 {
		for i := 0; i < n; i += l {
			for j := 0; j < k; j++ {
				a[i+j+k] += a[i+j] * op // 注意负数
			}
		}
	}
	return a
}

func fwtAND(a []int, op int) []int {
	n := len(a)
	for l, k := 2, 1; l <= n; l, k = l<<1, k<<1 {
		for i := 0; i < n; i += l {
			for j := 0; j < k; j++ {
				a[i+j] += a[i+j+k] * op // 注意负数
			}
		}
	}
	return a
}

func fwtXOR(a []int, op int) []int {
	n := len(a)
	for l, k := 2, 1; l <= n; l, k = l<<1, k<<1 {
		for i := 0; i < n; i += l {
			for j := 0; j < k; j++ {
				// 若题目没有取模，IFWT 时 *op 改成 /2
				a[i+j], a[i+j+k] = (a[i+j]+a[i+j+k])*op, (a[i+j]-a[i+j+k])*op // 注意负数
			}
		}
	}
	return a
}

// 求 OR 和 AND 时 invOp = -1
// 求 XOR 时 invOp = inv(2)
func fwt(a, b []int, fwtFunc func([]int, int) []int, invOp int) []int {
	// 不修改原始数组
	a = fwtFunc(append([]int(nil), a...), 1)
	b = fwtFunc(append([]int(nil), b...), 1)
	for i, v := range b {
		a[i] *= v // mod
	}
	c := fwtFunc(a, invOp)
	return c
}

// 注：若代码性能瓶颈在 FWT 上，可以通过以下方式消除比较慢的乘法和取模（CF 上需要将 int64 改成 int）
// 优化前 1575ms https://codeforces.com/contest/1218/submission/118700754
// 优化后  748ms https://codeforces.com/contest/1218/submission/118704484
const _mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= _mod {
		a -= _mod
	}
	return a
}

func sub(a, b int) int {
	a -= b
	if a < 0 {
		a += _mod
	}
	return a
}

func div2(a int) int {
	if a&1 > 0 {
		a += _mod
	}
	return a >> 1
}
