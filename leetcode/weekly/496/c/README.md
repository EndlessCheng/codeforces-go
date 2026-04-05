对于要增大的数 $\textit{nums}[i]$，必须增大到 $\max(\textit{nums}[i-1], \textit{nums}[i+1]) + 1$，操作次数为

$$
\max(\textit{nums}[i-1], \textit{nums}[i+1]) - \textit{nums}[i] + 1
$$

如果 $i$ 已经是特殊下标，则上式 $\le 0$，所以还要把上式与 $0$ 取最大值。

分类讨论：

- 如果 $n$ 是奇数，那么要修改的下标是唯一的，即所有奇数下标。枚举奇数下标，累加上式，即为答案。
- 如果 $n$ 是偶数，那么枚举修改 $[1,i]$ 中的奇数下标，以及 $[i+3,n-2]$ 中的偶数下标。

可以预处理 $\textit{suf}[i]$ 表示 $[i+3,n-2]$ 中的偶数下标的总操作次数。这可以倒着递推算出来（类似后缀和）。

答案可以是 $[2,n-2]$ 中的偶数下标的总操作次数，所以初始化答案为 $\textit{suf}[2]$。

然后一边遍历 $\textit{nums}$，一边计算 $[1,i]$ 中的奇数下标的总操作次数 $\textit{pre}$，用 $\textit{pre} + \textit{suf}[i+3]$ 更新答案的最小值。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

## 优化前

```py [sol-Python3]
class Solution:
    def minIncrease(self, nums: List[int]) -> int:
        n = len(nums)
        suf = [0] * (n + 1)
        for i in range(n - 2, 0, -2):
            suf[i] = suf[i + 2] + max(max(nums[i - 1], nums[i + 1]) - nums[i] + 1, 0)

        if n % 2 > 0:
            # 修改所有奇数下标
            return suf[1]

        ans = suf[2]  # 修改 [2,n-2] 中的所有偶数下标
        pre = 0
        # 枚举修改 [1,i] 中的奇数下标，以及 [i+3,n-2] 中的偶数下标
        for i in range(1, n - 1, 2):
            pre += max(max(nums[i - 1], nums[i + 1]) - nums[i] + 1, 0)
            ans = min(ans, pre + suf[i + 3])

        return ans
```

```java [sol-Java]
class Solution {
    public long minIncrease(int[] nums) {
        int n = nums.length;
        long[] suf = new long[n + 1];
        for (int i = n - 2; i > 0; i -= 2) {
            suf[i] = suf[i + 2] + Math.max(Math.max(nums[i - 1], nums[i + 1]) - nums[i] + 1, 0);
        }

        if (n % 2 > 0) {
            // 修改所有奇数下标
            return suf[1];
        }

        long ans = suf[2]; // 修改 [2,n-2] 中的所有偶数下标
        long pre = 0;
        // 枚举修改 [1,i] 中的奇数下标，以及 [i+3,n-2] 中的偶数下标
        for (int i = 1; i < n - 1; i += 2) {
            pre += Math.max(Math.max(nums[i - 1], nums[i + 1]) - nums[i] + 1, 0);
            ans = Math.min(ans, pre + suf[i + 3]);
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minIncrease(vector<int>& nums) {
        int n = nums.size();
        vector<long long> suf(n + 1);
        for (int i = n - 2; i > 0; i -= 2) {
            suf[i] = suf[i + 2] + max(max(nums[i - 1], nums[i + 1]) - nums[i] + 1, 0);
        }

        if (n % 2 > 0) {
            // 修改所有奇数下标
            return suf[1];
        }

        long long ans = suf[2]; // 修改 [2,n-2] 中的所有偶数下标
        long long pre = 0;
        // 枚举修改 [1,i] 中的奇数下标，以及 [i+3,n-2] 中的偶数下标
        for (int i = 1; i < n - 1; i += 2) {
            pre += max(max(nums[i - 1], nums[i + 1]) - nums[i] + 1, 0);
            ans = min(ans, pre + suf[i + 3]);
        }

        return ans;
    }
};
```

```go [sol-Go]
func minIncrease(nums []int) int64 {
	n := len(nums)
	suf := make([]int, n+1)
	for i := n - 2; i > 0; i -= 2 {
		suf[i] = suf[i+2] + max(max(nums[i-1], nums[i+1])-nums[i]+1, 0)
	}

	if n%2 > 0 {
		// 修改所有奇数下标
		return int64(suf[1])
	}

	ans := suf[2] // 修改 [2,n-2] 中的所有偶数下标
	pre := 0
	// 枚举修改 [1,i] 中的奇数下标，以及 [i+3,n-2] 中的偶数下标
	for i := 1; i < n-1; i += 2 {
		pre += max(max(nums[i-1], nums[i+1])-nums[i]+1, 0)
		ans = min(ans, pre+suf[i+3])
	}

	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 空间优化

```py [sol-Python3]
class Solution:
    def minIncrease(self, nums: List[int]) -> int:
        n = len(nums)
        suf = 0
        for i in range(n - 2, 0, -2):
            suf += max(max(nums[i - 1], nums[i + 1]) - nums[i] + 1, 0)

        if n % 2 > 0:
            # 修改所有奇数下标
            return suf

        ans = suf  # 修改 [2,n-2] 中的所有偶数下标
        pre = 0
        # 枚举修改 [1,i] 中的奇数下标，以及 [i+3,n-2] 中的偶数下标
        for i in range(1, n - 1, 2):
            pre += max(max(nums[i - 1], nums[i + 1]) - nums[i] + 1, 0)
            suf -= max(max(nums[i], nums[i + 2]) - nums[i + 1] + 1, 0)  # 撤销 i+1，撤销后 suf 对应 [i+3,n-2]
            ans = min(ans, pre + suf)

        return ans
```

```java [sol-Java]
class Solution {
    public long minIncrease(int[] nums) {
        int n = nums.length;
        long suf = 0;
        for (int i = n - 2; i > 0; i -= 2) {
            suf += Math.max(Math.max(nums[i - 1], nums[i + 1]) - nums[i] + 1, 0);
        }

        if (n % 2 > 0) {
            // 修改所有奇数下标
            return suf;
        }

        long ans = suf; // 修改 [2,n-2] 中的所有偶数下标
        long pre = 0;
        // 枚举修改 [1,i] 中的奇数下标，以及 [i+3,n-2] 中的偶数下标
        for (int i = 1; i < n - 1; i += 2) {
            pre += Math.max(Math.max(nums[i - 1], nums[i + 1]) - nums[i] + 1, 0);
            suf -= Math.max(Math.max(nums[i], nums[i + 2]) - nums[i + 1] + 1, 0); // 撤销 i+1，撤销后 suf 对应 [i+3,n-2]
            ans = Math.min(ans, pre + suf);
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minIncrease(vector<int>& nums) {
        int n = nums.size();
        long long suf = 0;
        for (int i = n - 2; i > 0; i -= 2) {
            suf += max(max(nums[i - 1], nums[i + 1]) - nums[i] + 1, 0);
        }

        if (n % 2 > 0) {
            // 修改所有奇数下标
            return suf;
        }

        long long ans = suf; // 修改 [2,n-2] 中的所有偶数下标
        long long pre = 0;
        // 枚举修改 [1,i] 中的奇数下标，以及 [i+3,n-2] 中的偶数下标
        for (int i = 1; i < n - 1; i += 2) {
            pre += max(max(nums[i - 1], nums[i + 1]) - nums[i] + 1, 0);
            suf -= max(max(nums[i], nums[i + 2]) - nums[i + 1] + 1, 0); // 撤销 i+1，撤销后 suf 对应 [i+3,n-2]
            ans = min(ans, pre + suf);
        }

        return ans;
    }
};
```

```go [sol-Go]
func minIncrease(nums []int) int64 {
	n := len(nums)
	suf := 0
	for i := n - 2; i > 0; i -= 2 {
		suf += max(max(nums[i-1], nums[i+1])-nums[i]+1, 0)
	}

	if n%2 > 0 {
		// 修改所有奇数下标
		return int64(suf)
	}

	ans := suf // 修改 [2,n-2] 中的所有偶数下标
	pre := 0
	// 枚举修改 [1,i] 中的奇数下标，以及 [i+3,n-2] 中的偶数下标
	for i := 1; i < n-1; i += 2 {
		pre += max(max(nums[i-1], nums[i+1])-nums[i]+1, 0)
		suf -= max(max(nums[i], nums[i+2])-nums[i+1]+1, 0) // 撤销 i+1，撤销后 suf 对应 [i+3,n-2]
		ans = min(ans, pre+suf)
	}

	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面动态规划题单的「**专题：前后缀分解**」。

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
