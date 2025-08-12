如果 $\textit{nums}$ 已经是有序的，按照题目要求，返回 $0$。下文讨论 $\textit{nums}$ 不是有序的情况。

先思考一个问题：本题是否一定有解？

这是肯定的，$k=0$ 的时候一定可以排成有序数组，构造方式如下：

- 找到等于 $0$ 的那个数 $\textit{nums}[i]$。
- 如果 $\textit{nums}[0]\ne 0$，先把 $\textit{nums}[i]=0$ 与 $\textit{nums}[0]$ 交换。因为 $0$ 和任何数的 AND 都等于 $0$，所以交换满足题目要求。下同。
- 我们可以用**三次交换**实现任意一对元素的交换：
  - 用 $0,a,b$ 表示这三个数（注意是元素值，不是下标）。目标是得到 $0,b,a$。
  - 交换 $0$ 和 $a$，得到 $a,0,b$。
  - 交换 $0$ 和 $b$，得到 $a,b,0$。
  - 交换 $0$ 和 $a$，得到 $0,b,a$。
  - 注意三次交换后，$0$ 的位置是**不变**的。
- 因此，当 $k=0$ 的时候，一定可以排序。

由于 $\textit{nums}$ 是 $0$ 到 $n-1$ 的排列，所以排序后，对于每个 $i$ 都满足 $\textit{nums}[i] = i$。

如果某些元素已经满足 $\textit{nums}[i] = i$ 呢？$k$ 能不能再大点？

首先确定答案的**上界**。

对于不满足 $\textit{nums}[i] = i$ 的元素，一定要交换，这些元素一定会出现在式子 $\textit{nums}[i]\ \texttt{AND}\ \textit{nums}[j] = k$ 中。根据 [从集合论到位运算](https://leetcode.cn/circle/discuss/CaOJ45/)，$k$ 是 $\textit{nums}[i]$ 和 $\textit{nums}[j]$ 的子集，所以 $k$ 是每个不满足 $\textit{nums}[i] = i$ 的 $\textit{nums}[i]$ 的子集。$k$ 最大是这些集合的**交集**。所以答案至多为所有不满足 $\textit{nums}[i] = i$ 的 $\textit{nums}[i]$ 的 AND，下文就用 $k$ 表示这个上界。

答案能等于上界 $k$ 吗？可以，方法和上面的 $0$ 是一样的：

- 如果 $\textit{nums}[k]\ne k$，那么找到等于 $k$ 的那个数，和 $\textit{nums}[k]$ 交换。注意：由于 $\textit{nums}[k]\ne k$，所以 $\textit{nums}[k]$ 参与上面的交集运算，所以 $k$ 是 $\textit{nums}[k]$ 的子集，所以 $\textit{nums}[k] \ \texttt{AND}\ k = k$，这个交换满足题目要求。
- 然后用三次交换实现任意一对元素的交换，做法同上。由于 $k$ 是这些需要交换的元素的子集，所以这些元素与 $k$ 的 AND 都等于 $k$，满足题目要求。

**注**：由于参与 AND 运算的数越多，结果越小，所以 $0\le k\le \max(\textit{nums}) = n-1$，所以 $k$ 一定在 $\textit{nums}$ 中。 

```py [sol-Python3]
class Solution:
    def sortPermutation(self, nums: List[int]) -> int:
        if nums[0]:  # 小优化：此时 0 参与 AND，结果一定是 0
            return 0
        ans = -1  # 二进制全为 1
        for i, x in enumerate(nums):
            if i != x:
                ans &= x
        return max(ans, 0)
```

```java [sol-Java]
class Solution {
    public int sortPermutation(int[] nums) {
        int ans = -1; // 二进制全为 1
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            if (i != x) {
                ans &= x;
            }
        }
        return Math.max(ans, 0);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int sortPermutation(vector<int>& nums) {
        int ans = -1; // 二进制全为 1
        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];
            if (i != x) {
                ans &= x;
            }
        }
        return max(ans, 0);
    }
};
```

```go [sol-Go]
func sortPermutation(nums []int) int {
	ans := -1 // 二进制全为 1
	for i, x := range nums {
		if i != x {
			ans &= x
		}
	}
	return max(ans, 0)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

有多少个不同的 $k$，可以让数组排成有序的？（保证数组不是有序的）

## 专题训练

见下面贪心与思维题单的「**§5.2 脑筋急转弯**」。

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
