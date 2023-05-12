package copypasta

/* 快速沃尔什变换 fast Walsh–Hadamard transform, FWT, FWHT
在算法竞赛中，FWT 是用于解决对下标进行【位运算卷积】问题的方法
一个常见的应用场景是对频率数组求 FWT
例如把 a 里面所有元素的出现次数统计到一个 cnt 数组中，cnt[x] 表示 x 在 a 中的出现次数
然后求 FWT(cnt, cnt)，就得到了 a 中任意两个元素的【异或/或/和】的结果的出现次数（特殊地，这得到了任意两个数的【异或/或/和】的最大值）
又例如，求一个数组的三个元素的最大异或和，在值域不大的情况下，
可以先求出该数组的频率数组与频率数组的 FWT，即得到两个元素的所有异或和（及组成该异或和的元素对数），
然后枚举两元素异或和，在原数组的异或字典树上查询最大异或和
具体到名称，OR 上的 FWT 也叫 fast zeta transform，AND 上的 FWT 也叫 fast mobius transform
https://en.wikipedia.org/wiki/Fast_Walsh%E2%80%93Hadamard_transform
https://oi-wiki.org/math/poly/fwt/
OR Convolution for Common People https://codeforces.com/blog/entry/115438
https://www.luogu.com.cn/blog/CJL/fwt-xue-xi-bi-ji
https://www.luogu.com.cn/blog/command-block/wei-yun-suan-juan-ji-yu-ji-kuo-zhan
一些习题 https://blog.csdn.net/weixin_38686780/article/details/81912853

模板题 https://www.luogu.com.cn/problem/P4717
LC1803 https://leetcode-cn.com/problems/count-pairs-with-xor-in-a-range/
与仙人掌结合 https://codeforces.com/problemset/problem/1218/D
todo https://codeforces.com/problemset/problem/662/C
todo https://atcoder.jp/contests/abc212/tasks/abc212_h

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
const _mod = 1_000_000_007

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

// 将 (-mod,mod) 范围内的 a 变成 [0,mod) 范围内
// 原理是负数右移会不断补 1，所以最后二进制都是 1，因此返回值等价于 a+_mod
// 而对于非负数，右移后二进制全为 0，所以返回结果仍然是 a
func norm32(a int32) int32 {
	return a + a>>31&_mod
}

func norm64(a int64) int64 {
	return a + a>>63&_mod
}
