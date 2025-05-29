![lc3439.png](https://pic.leetcode.cn/1738458316-BLfZXa-lc3439.png)

看示例 1，把会议区间 $[1,2]$ 移动到 $[0,1]$ 或者 $[2,3]$，会产生空余时间段 $[1,3]$ 或者 $[0,2]$，相当于把两个**相邻**的长为 $1$ 空余时间段 $[0,1]$ 和 $[2,3]$ **合并**成一个更大的长为 $1+1=2$ 的空余时间段。

如果 $k=1$，那么我们可以合并 $2$ 个相邻的空余时间段。

如果 $k=2$，为了让答案尽量大，合并连续 $3$ 个空余时间段，相比其他方案是最优的。（注意题目要求会议之间的相对顺序必须保持不变）

一般地，最优做法是合并**连续** $k+1$ 个空余时间段。

现在问题变成：

- 给你 $n+1$ 个空余时间段，合并其中**连续** $k+1$ 个空余时间段，得到的最大长度是多少？

> **注 1**：最左边和最右边各有一个空余时间段，中间有 $n-1$ 个空余时间段夹在两个相邻会议之间，所以有 $n+1$ 个空余时间段。
> 
> **注 2**：空余时间段的长度可以是 $0$。

这可以用**定长滑动窗口**解决，窗口大小为 $k+1$。原理见[【套路】教你解决定长滑窗！适用于所有定长滑窗题目！](https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/solutions/2809359/tao-lu-jiao-ni-jie-jue-ding-chang-hua-ch-fzfo/)

[本题视频讲解](https://www.bilibili.com/video/BV1eUF6eaERQ/?t=38s)，欢迎点赞关注~

## 写法一

创建一个长为 $n+1$ 的数组，保存所有空余时间段的长度：

- 最左边的空余时间段：长为 $\textit{startTime}[0]$。
- 中间的空余时间段：长为 $\textit{startTime}[i] - \textit{endTime}[i-1]$。
- 最右边的空余时间段：长为 $\textit{eventTime} - \textit{endTime}[n-1]$。

示例 1 的数组为 $[1, 1, 0]$。

示例 2 的数组为 $[0, 1, 5, 0]$。

示例 3 的数组为 $[0, 0, 0, 0, 0, 0]$。

```py [sol-Python3]
class Solution:
    def maxFreeTime(self, eventTime: int, k: int, startTime: List[int], endTime: List[int]) -> int:
        n = len(startTime)
        free = [0] * (n + 1)
        free[0] = startTime[0]  # 最左边的空余时间段
        for i in range(1, n):
            free[i] = startTime[i] - endTime[i - 1]  # 中间的空余时间段
        free[n] = eventTime - endTime[-1]  # 最右边的空余时间段

        # 套定长滑窗模板（窗口长为 k+1）
        ans = s = 0
        for i, f in enumerate(free):
            s += f
            if i < k:
                continue
            ans = max(ans, s)
            s -= free[i - k]
        return ans
```

```py [sol-Python3 精简写法]
class Solution:
    def maxFreeTime(self, eventTime: int, k: int, startTime: List[int], endTime: List[int]) -> int:
        free = [startTime[0]] + [s - e for s, e in zip(startTime[1:], endTime)] + [eventTime - endTime[-1]]
        ans = s = 0
        for i, f in enumerate(free):
            s += f
            if i < k:
                continue
            ans = max(ans, s)
            s -= free[i - k]
        return ans
```

```java [sol-Java]
class Solution {
    public int maxFreeTime(int eventTime, int k, int[] startTime, int[] endTime) {
        int n = startTime.length;
        int[] free = new int[n + 1];
        free[0] = startTime[0]; // 最左边的空余时间段
        for (int i = 1; i < n; i++) {
            free[i] = startTime[i] - endTime[i - 1]; // 中间的空余时间段
        }
        free[n] = eventTime - endTime[n - 1]; // 最右边的空余时间段

        // 套定长滑窗模板（窗口长为 k+1）
        int ans = 0;
        int s = 0;
        for (int i = 0; i <= n; i++) {
            s += free[i];
            if (i < k) {
                continue;
            }
            ans = Math.max(ans, s);
            s -= free[i - k];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxFreeTime(int eventTime, int k, vector<int>& startTime, vector<int>& endTime) {
        int n = startTime.size();
        vector<int> free(n + 1);
        free[0] = startTime[0]; // 最左边的空余时间段
        for (int i = 1; i < n; i++) {
            free[i] = startTime[i] - endTime[i - 1]; // 中间的空余时间段
        }
        free[n] = eventTime - endTime[n - 1]; // 最右边的空余时间段

        // 套定长滑窗模板（窗口长为 k+1）
        int ans = 0, s = 0;
        for (int i = 0; i <= n; i++) {
            s += free[i];
            if (i < k) {
                continue;
            }
            ans = max(ans, s);
            s -= free[i - k];
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxFreeTime(eventTime, k int, startTime, endTime []int) (ans int) {
	n := len(startTime)
	free := make([]int, n+1)
	free[0] = startTime[0] // 最左边的空余时间段
	for i := 1; i < n; i++ {
		free[i] = startTime[i] - endTime[i-1] // 中间的空余时间段
	}
	free[n] = eventTime - endTime[n-1] // 最右边的空余时间段

	// 套定长滑窗模板（窗口长为 k+1）
	s := 0
	for i, f := range free {
		s += f
		if i < k {
			continue
		}
		ans = max(ans, s)
		s -= free[i-k]
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{startTime}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 写法二：空间优化

改用一个函数计算第 $i$ 个空余时间段的长度，从而省去 $\textit{free}$ 数组。

```py [sol-Python3]
class Solution:
    def maxFreeTime(self, eventTime: int, k: int, startTime: List[int], endTime: List[int]) -> int:
        def get(i: int) -> int:
            if i == 0:
                return startTime[0]  # 最左边的空余时间段
            if i == n:
                return eventTime - endTime[-1]  # 最右边的空余时间段
            return startTime[i] - endTime[i - 1]  # 中间的空余时间段

        n = len(startTime)
        ans = s = 0
        for i in range(n + 1):
            s += get(i)
            if i < k:
                continue
            ans = max(ans, s)
            s -= get(i - k)
        return ans
```

```java [sol-Java]
class Solution {
    public int maxFreeTime(int eventTime, int k, int[] startTime, int[] endTime) {
        int ans = 0;
        int s = 0;
        for (int i = 0; i <= startTime.length; i++) {
            s += get(i, eventTime, startTime, endTime);
            if (i < k) {
                continue;
            }
            ans = Math.max(ans, s);
            s -= get(i - k, eventTime, startTime, endTime);
        }
        return ans;
    }

    private int get(int i, int eventTime, int[] startTime, int[] endTime) {
        if (i == 0) {
            return startTime[0]; // 最左边的空余时间段
        }
        int n = startTime.length;
        if (i == n) {
            return eventTime - endTime[n - 1]; // 最右边的空余时间段
        }
        return startTime[i] - endTime[i - 1]; // 中间的空余时间段
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxFreeTime(int eventTime, int k, vector<int>& startTime, vector<int>& endTime) {
        int n = startTime.size();
        auto get = [&](int i) -> int {
            if (i == 0) {
                return startTime[0]; // 最左边的空余时间段
            }
            if (i == n) {
                return eventTime - endTime[n - 1]; // 最右边的空余时间段
            }
            return startTime[i] - endTime[i - 1]; // 中间的空余时间段
        };

        int s = 0, ans = 0;
        for (int i = 0; i <= n; i++) {
            s += get(i);
            if (i < k) {
                continue;
            }
            ans = max(ans, s);
            s -= get(i - k);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxFreeTime(eventTime, k int, startTime, endTime []int) (ans int) {
	n := len(startTime)
	get := func(i int) int {
		if i == 0 {
			return startTime[0] // 最左边的空余时间段
		}
		if i == n {
			return eventTime - endTime[n-1] // 最右边的空余时间段
		}
		return startTime[i] - endTime[i-1] // 中间的空余时间段
	}

	s := 0
	for i := range n + 1 {
		s += get(i)
		if i < k {
			continue
		}
		ans = max(ans, s)
		s -= get(i - k)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{startTime}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
