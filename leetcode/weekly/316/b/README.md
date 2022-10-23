[视频讲解](https://www.bilibili.com/video/BV1ne4y1e7nu) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

# 方法一：暴力枚举

暴力枚举即可，可以在不包含因子 $k$ 时提前退出循环。

```py [sol1-Python3]
class Solution:
    def subarrayGCD(self, nums: List[int], k: int) -> int:
        ans = 0
        for i in range(len(nums)):
            g = 0
            for j in range(i, len(nums)):
                g = gcd(g, nums[j])
                if g % k: break
                if g == k: ans += 1
        return ans
```

```go [sol1-Go]
func subarrayGCD(nums []int, k int) (ans int) {
	for i := range nums {
		g := 0
		for _, x := range nums[i:] {
			g = gcd(g, x)
			if g%k > 0 {
				break
			}
			if g == k {
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

- 时间复杂度：$O(n^2\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=max(\textit{nums})$。
- 空间复杂度：$O(1)$，仅用到若干变量。

# 方法二：利用 GCD 的性质

原理见我之前写的 [这篇题解的方法二](https://leetcode.cn/problems/smallest-subarrays-with-maximum-bitwise-or/solution/by-endlesscheng-zai1/)，[视频讲解](https://www.bilibili.com/video/BV1ne4y1e7nu) 详细介绍了这种算法的原理。

```py [sol1-Python3]
class Solution:
    def subarrayGCD(self, nums: List[int], k: int) -> int:
        ans = 0
        a = []  # [GCD，相同 GCD 区间的右端点]
        i0 = -1
        for i, x in enumerate(nums):
            if x % k:  # 保证后续求的 GCD 都是 k 的倍数
                a = []
                i0 = i
                continue
            a.append([x, i])
            # 原地去重，因为相同的 GCD 都相邻在一起
            j = 0
            for p in a:
                p[0] = gcd(p[0], x)
                if a[j][0] != p[0]:
                    j += 1
                    a[j] = p
                else:
                    a[j][1] = p[1]
            del a[j + 1:]
            if a[0][0] == k:  # a[0][0] >= k
                ans += a[0][1] - i0
        return ans
```

```go [sol1-Go]
func subarrayGCD(nums []int, k int) (ans int) {
	type result struct{ v, i int }
	var a []result
	i0 := -1
	for i, v := range nums {
		if v%k > 0 {
			a = nil
			i0 = i
			continue
		}
		for j, p := range a {
			a[j].v = gcd(p.v, v)
		}
		a = append(a, result{v, i})
		j := 0
		for _, q := range a[1:] {
			if a[j].v != q.v {
				j++
				a[j] = q
			} else {
				a[j].i = q.i
			}
		}
		a = a[:j+1]
		if a[0].v == k {
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

- 时间复杂度：$O(n\log^2 U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=max(\textit{nums})$。
- 空间复杂度：$O(\log U)$。

#### 相似题目

- [Codeforces 475D. CGCDSSQ](https://codeforces.com/problemset/problem/475/D)
