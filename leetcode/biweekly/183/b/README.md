枚举 $x$ 和 $y$。

把 $\textit{nums}[i]$ 变成满足 $\textit{nums}[i]\bmod k = x$ 的数，也就是变成一个与 $x$ 关于模 $k$ [同余](https://leetcode.cn/circle/discuss/mDfnkW/) 的数。

操作次数等价于：

- 在一个长为 $k$ 的环上，计算 $a = \textit{nums}[i]\bmod k$ 到 $x$ 的最短距离。

例如 $k=12$，类似生活中的圆形挂钟（闹钟），把小时从 $17$ 点（即 $5$ 点位置）拨到 $3$ 点位置，需要往回拨 $2$ 圈。

设 $d = |a - x|$，分类讨论：

- 不经过 $0$，直接从 $a$ 走到 $x$，需要 $d$ 步。
- 经过 $0$，绕一圈从 $a$ 走到 $x$，需要 $k-d$ 步。

取二者的最小值 $\min(d, k-d)$，即为 $a$ 到 $x$ 的最短距离。

[本题视频讲解](https://www.bilibili.com/video/BV1iuG76VEXy/?t=4m4s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: list[int], k: int) -> int:
        ans = inf

        for x in range(k):
            for y in range(k):
                if y == x:
                    continue
                target = [x, y]
                s = 0
                for i, v in enumerate(nums):
                    d = abs(v % k - target[i % 2])
                    s += min(d, k - d)  # 直接走到 target[i%2]，或者反向绕一圈到 target[i%2]
                ans = min(ans, s)

        return ans
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums, int k) {
        int ans = Integer.MAX_VALUE;

        for (int x = 0; x < k; x++) {
            for (int y = 0; y < k; y++) {
                if (y == x) {
                    continue;
                }
                int[] target = new int[]{x, y};
                int sum = 0;
                for (int i = 0; i < nums.length; i++) {
                    int d = Math.abs(nums[i] % k - target[i % 2]);
                    sum += Math.min(d, k - d); // 直接走到 target[i%2]，或者反向绕一圈到 target[i%2]
                }
                ans = Math.min(ans, sum);
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(vector<int>& nums, int k) {
        int ans = INT_MAX;

        for (int x = 0; x < k; x++) {
            for (int y = 0; y < k; y++) {
                if (y == x) {
                    continue;
                }
                int target[2] = {x, y};
                int sum = 0;
                for (int i = 0; i < nums.size(); i++) {
                    int d = abs(nums[i] % k - target[i % 2]);
                    sum += min(d, k - d); // 直接走到 target[i%2]，或者反向绕一圈到 target[i%2]
                }
                ans = min(ans, sum);
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func minOperations(nums []int, k int) int {
	ans := math.MaxInt
	for x := range k {
		for y := range k {
			if y == x {
				continue
			}
			target := [2]int{x, y}
			sum := 0
			for i, v := range nums {
				d := abs(v%k - target[i%2])
				sum += min(d, k-d) // 直接走到 target[i%2]，或者反向绕一圈到 target[i%2]
			}
			ans = min(ans, sum)
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk^2)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

**注**：分别解决奇偶，本题是**环形邮局问题**，利用 [中位数贪心](https://zhuanlan.zhihu.com/p/1922938031687595039) 的结论，可以做到 $\mathcal{O}(n\log n)$ 时间复杂度（瓶颈在排序上）。

## 专题训练

见下面贪心题单的「**§4.5 中位数贪心**」。

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
