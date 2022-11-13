思路和 [2447. 最大公因数等于 K 的子数组数目](https://leetcode.cn/problems/number-of-subarrays-with-gcd-equal-to-k/) 完全一致。

# 方法一：暴力枚举

```py [sol1-Python3]
class Solution:
    def subarrayLCM(self, nums: List[int], k: int) -> int:
        ans, n = 0, len(nums)
        for i in range(n):
            res = 1
            for j in range(i, n):
                res = lcm(res, nums[j])
                if k % res: break  # 剪枝：LCM 必须是 k 的因子
                if res == k: ans += 1
        return ans
```

```go [sol1-Go]
func subarrayLCM(nums []int, k int) (ans int) {
	for i := range nums {
		lcm := 1
		for _, x := range nums[i:] {
			lcm = lcm / gcd(lcm, x) * x
			if k%lcm > 0 { // 剪枝：lcm 必须是 k 的因子
				break
			}
			if lcm == k {
				ans++
			}
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$O(n(n+\log k))$，其中 $n$ 为 $\textit{nums}$ 的长度。LCM 倍增的次数是 $O(\log k)$ 次，因此内层循环的时间复杂度为 $O(n+\log k)$，所以总的时间复杂度为 $O(n(n+\log k))$。
- 空间复杂度：$O(1)$，仅用到若干变量。

# 方法二：利用 LCM 的性质

参考我之前的那篇 [题解](https://leetcode.cn/problems/number-of-subarrays-with-gcd-equal-to-k/solutions/1917454/by-endlesscheng-1f1r/)。

最小公倍数要么不变，要么至少 $\times 2$，因此在遍历 $\textit{nums}$ 的同时，维护最小公倍数集合（数组），这至多有 $O(\log k)$ 个。

注意最小公倍数必须是 $k$ 的因子。

由于切片的原因，实现起来 Go 应该是最舒服的。

```go [sol2-Go]
func subarrayLCM(nums []int, k int) (ans int) {
	type result struct{ lcm, i int }
	var a []result
	i0 := -1
	for i, x := range nums {
		if k%x > 0 {
			a = nil
			i0 = i
			continue
		}
		for j, p := range a {
			a[j].lcm = p.lcm / gcd(p.lcm, x) * x
		}
		for len(a) > 0 && k%a[0].lcm > 0 { // 去除不合法的 LCM
			a = a[1:]
		}
		a = append(a, result{x, i})
		j := 0
		for _, q := range a[1:] {
			if a[j].lcm != q.lcm {
				j++
				a[j] = q
			} else {
				a[j].i = q.i
			}
		}
		a = a[:j+1]
		if a[0].lcm == k {
			ans += a[0].i - i0
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log k)$，其中 $n$ 为 $\textit{nums}$ 的长度。LCM 倍增的次数是 $O(\log k)$ 次，并且每次去重的时间复杂度也为 $O(\log k)$，因此时间复杂度为 $O(n\log k)$。
- 空间复杂度：$O(\log k)$。

#### 相似题目

- [Codeforces 475D. CGCDSSQ](https://codeforces.com/problemset/problem/475/D)
