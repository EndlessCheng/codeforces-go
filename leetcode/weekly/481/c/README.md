首先判断无解的情况。

把 $\textit{nums}$ 和 $\textit{forbidden}$ 排成两行，得到一个 $2\times n$ 的矩阵。

如果某个数 $x$ 在 $\textit{nums}$ 和 $\textit{forbidden}$ 中总出现次数超过 $n$，那么根据鸽巢原理（抽屉原理），必然有一列都是 $x$，此时无解。比如示例 3 有 $3$ 个 $7$，$3 > n$。

否则一定有解，以示例 2 为例说明。

把两个数组排成两行：

$$
\begin{bmatrix}
4 & 6 & 6 & 5   \\
4 & 6 & 5 & 5   \\
\end{bmatrix}
$$

有三列元素相同。把这样的列叫做「坏列」。

贪心地，优先交换元素不同的坏列，可以一次性消除两个坏列。比如交换 $4$ 和 $6$，得到

$$
\begin{bmatrix}
6 & 4 & 6 & 5   \\
4 & 6 & 5 & 5   \\
\end{bmatrix}
$$

剩下 $[5,5]$ 这个坏列，找两个数都不等于 $5$ 的列交换（比如第一列），得到

$$
\begin{bmatrix}
5 & 4 & 6 & 6   \\
4 & 6 & 5 & 5   \\
\end{bmatrix}
$$

注意两个数都不等于 $5$ 的列是一定存在的。**反证法**：如果每一列都至少有一个 $5$，再算上最后一列有两个 $5$，可知 $5$ 的总数大于 $n$，矛盾。

如果有很多个坏列，每次操作应该选择哪两个坏列消除更好呢？比如

$$
\begin{bmatrix}
1 & 1 & 1 & 2 & 2 & 3  \\
1 & 1 & 1 & 2 & 2 & 3  \\
\end{bmatrix}
$$

如果先交换 $2$ 和 $3$，得到

$$
\begin{bmatrix}
1 & 1 & 1 & 2 & 3 & 2  \\
1 & 1 & 1 & 2 & 2 & 3  \\
\end{bmatrix}
$$

再交换 $1$ 和 $2$，得到 

$$
\begin{bmatrix}
1 & 1 & 2 & 1 & 3 & 2  \\
1 & 1 & 1 & 2 & 2 & 3  \\
\end{bmatrix}
$$

剩余的两个坏列只能找好列交换，一共操作 $4$ 次。

然而，如果交换 $1$ 和 $2$，$1$ 和 $2$，$1$ 和 $3$，操作 $3$ 次，就能得到

$$
\begin{bmatrix}
2 & 2 & 3 & 1 & 1 & 1  \\
1 & 1 & 1 & 2 & 2 & 3  \\
\end{bmatrix}
$$

一般地，把坏列第一行的元素记在数组 $a$ 中（上例 $a=[1,1,1,2,2,3]$），问题变成：

- 给定数组 $a$，每次操作，删除 $a$ 中的**至多**两个**不同**元素。删除所有元素，最少要操作多少次？

**答**：最少操作 $\max\left(\left\lceil\dfrac{k}{2}\right\rceil,\textit{mx}\right)$ 次，其中 $k$ 是 $a$ 的长度，$\textit{mx}$ 是 $a$ 中出现次数最多的元素的出现次数。

[证明+具体操作方案](https://zhuanlan.zhihu.com/p/1945782212176909162)。

[本题视频讲解](https://www.bilibili.com/video/BV1HsqmBwEy3/?t=46m19s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minSwaps(self, nums: List[int], forbidden: List[int]) -> int:
        n = len(nums)
        total = Counter(nums) + Counter(forbidden)
        if any(c > n for c in total.values()):
            return -1

        same = Counter(x for x, y in zip(nums, forbidden) if x == y)
        k = same.total()
        mx = max(same.values(), default=0)
        return max((k + 1) // 2, mx)
```

```java [sol-Java]
class Solution {
    public int minSwaps(int[] nums, int[] forbidden) {
        int n = nums.length;
        Map<Integer, Integer> total = new HashMap<>();
        for (int x : nums) {
            total.merge(x, 1, Integer::sum);
        }

        Map<Integer, Integer> cnt = new HashMap<>();
        int k = 0;
        int mx = 0;
        for (int i = 0; i < n; i++) {
            int x = forbidden[i];
            int c = total.merge(x, 1, Integer::sum);
            if (c > n) {
                return -1;
            }
            if (x == nums[i]) {
                k++;
                c = cnt.merge(x, 1, Integer::sum);
                mx = Math.max(mx, c);
            }
        }

        return Math.max((k + 1) / 2, mx);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minSwaps(vector<int>& nums, vector<int>& forbidden) {
        int n = nums.size();
        unordered_map<int, int> total;
        for (int x : nums) {
            total[x]++;
        }

        unordered_map<int, int> cnt;
        int k = 0, mx = 0;
        for (int i = 0; i < forbidden.size(); i++) {
            int x = forbidden[i];
            if (++total[x] > n) {
                return -1;
            }
            if (x == nums[i]) {
                k++;
                mx = max(mx, ++cnt[x]);
            }
        }

        return max((k + 1) / 2, mx);
    }
};
```

```go [sol-Go]
func minSwaps(nums, forbidden []int) int {
	n := len(nums)
	total := map[int]int{}
	for _, x := range nums {
		total[x]++
	}

	cnt := map[int]int{}
	k, mx := 0, 0
	for i, x := range forbidden {
		total[x]++
		if total[x] > n {
			return -1
		}
		if x == nums[i] {
			k++
			cnt[x]++
			mx = max(mx, cnt[x])
		}
	}

	return max((k+1)/2, mx)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面贪心题单的「**§1.8 相邻不同**」。

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
