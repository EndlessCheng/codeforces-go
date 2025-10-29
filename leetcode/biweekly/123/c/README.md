[视频讲解](https://www.bilibili.com/video/BV14C411r7nN/) 第三题。

## 前置知识：前缀和

对于数组 $a$，定义它的前缀和 $\textit{s}[0]=0$，$\textit{s}[i+1] = \sum\limits_{j=0}^{i}a[j]$。

根据这个定义，有 $s[i+1]=s[i]+a[i]$。

例如 $a=[1,2,1,2]$，对应的前缀和数组为 $s=[0,1,3,4,6]$。

通过前缀和，我们可以把**连续子数组的元素和转换成两个前缀和的差**，$a[\textit{left}]$ 到 $a[\textit{right}]$ 的元素和等于

$$
\sum_{j=\textit{left}}^{\textit{right}}a[j] = \sum\limits_{j=0}^{\textit{right}}a[j] - \sum\limits_{j=0}^{\textit{left}-1}a[j] = \textit{s}[\textit{right}+1] - \textit{s}[\textit{left}]
$$

例如 $a$ 的子数组 $[2,1,2]$ 的和就可以用 $s[4]-s[1]=6-1=5$ 算出来。

**注**：$s[0]=0$ 表示一个空数组的元素和。为什么要额外定义它？想一想，如果要计算的子数组恰好是一个前缀（从 $a[0]$ 到 $a[\textit{right}]$），你要用 $s[\textit{right}+1]$ 减去谁呢？通过定义 $s[0]=0$，任意子数组（包括前缀）都可以表示为两个前缀和的差。

## 思路

为方便描述，把 $\textit{nums}$ 简称为 $a$。

子数组 $a[i..j]$ 的元素和为 

$$
s[j+1]-s[i]
$$

枚举 $j$，我们需要找到最小的 $s[i]$，满足 $|a[i]-a[j]|=k$，即 $a[i] = a[j]-k$ 或 $a[i]=a[j]+k$。

定义哈希表 $\textit{minS}$，键为 $a[i]$，值为相同 $a[i]$ 下的 $s[i]$ 的最小值。

子数组最后一个数为 $a[j]$ 时，子数组的最大元素和为

$$
\begin{aligned}
& s[j+1]-\textit{minS}[a[i]]\\
=\ &s[j+1]-\min(\textit{minS}[a[j]-k],\textit{minS}[a[j]+k])
\end{aligned}
$$

```py [sol-Python3]
class Solution:
    def maximumSubarraySum(self, nums: List[int], k: int) -> int:
        min_s = defaultdict(lambda: inf)
        s = 0
        ans = -inf
        for x in nums:
            ans = max(ans, s + x - min(min_s[x - k], min_s[x + k]))
            min_s[x] = min(min_s[x], s)
            s += x
        return ans if ans > -inf else 0
```

```java [sol-Java]
class Solution {
    public long maximumSubarraySum(int[] nums, int k) {
        Map<Integer, Long> minS = new HashMap<>();
        long sum = 0;
        long ans = Long.MIN_VALUE;
        for (int x : nums) {
            Long s = minS.get(x - k);
            if (s != null) {
                ans = Math.max(ans, sum + x - s);
            }

            s = minS.get(x + k);
            if (s != null) {
                ans = Math.max(ans, sum + x - s);
            }

            minS.merge(x, sum, Math::min); // minS[x] = Math.min(minS[x], sum)
            sum += x;
        }
        return ans == Long.MIN_VALUE ? 0 : ans;
    }
}
```

```java [sol-Java 写法二]
class Solution {
    public long maximumSubarraySum(int[] nums, int k) {
        Map<Integer, Long> minS = new HashMap<>();
        long sum = 0;
        long ans = Long.MIN_VALUE;
        for (int x : nums) {
            long s1 = minS.getOrDefault(x - k, Long.MAX_VALUE / 2);
            long s2 = minS.getOrDefault(x + k, Long.MAX_VALUE / 2);
            ans = Math.max(ans, sum + x - Math.min(s1, s2));
            minS.merge(x, sum, Math::min);
            sum += x;
        }
        return ans > Long.MIN_VALUE / 4 ? ans : 0;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumSubarraySum(vector<int>& nums, int k) {
        unordered_map<int, long long> min_s;
        long long ans = LLONG_MIN, sum = 0;
        for (int x : nums) {
            auto it = min_s.find(x + k);
            if (it != min_s.end()) {
                ans = max(ans, sum + x - it->second);
            }

            it = min_s.find(x - k);
            if (it != min_s.end()) {
                ans = max(ans, sum + x - it->second);
            }

            it = min_s.find(x);
            if (it == min_s.end() || sum < it->second) {
                min_s[x] = sum;
            }

            sum += x;
        }
        return ans == LLONG_MIN ? 0 : ans;
    }
};
```

```go [sol-Go]
func maximumSubarraySum(nums []int, k int) int64 {
	minS := map[int]int{}
	sum := 0
	ans := math.MinInt
	for _, x := range nums {
		s, ok := minS[x+k]
		if ok {
			ans = max(ans, sum+x-s)
		}

		s, ok = minS[x-k]
		if ok {
			ans = max(ans, sum+x-s)
		}

		s, ok = minS[x]
		if !ok || sum < s {
			minS[x] = sum
		}

		sum += x
	}
	if ans == math.MinInt {
		return 0
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面数据结构题单的「**§1.2 前缀和与哈希表**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
