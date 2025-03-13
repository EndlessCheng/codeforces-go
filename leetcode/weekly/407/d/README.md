**前置知识**：[差分数组原理讲解](https://leetcode.cn/problems/car-pooling/solution/suan-fa-xiao-ke-tang-chai-fen-shu-zu-fu-9d4ra/)，推荐和[【图解】从一维差分到二维差分](https://leetcode.cn/problems/stamping-the-grid/solution/wu-nao-zuo-fa-er-wei-qian-zhui-he-er-wei-zwiu/) 一起看。

计算两个数组的差值 $a[i] = \textit{target}[i] - \textit{nums}[i]$。把 $\textit{nums}$ 变成 $\textit{target}$，等价于把一个全为 $0$ 的数组变成 $a$，或者反过来，把 $a$ 变成一个全为 $0$ 的数组。

示例 1 相当于把 $a=[1,1,1,2]$ 变成一个全为 $0$ 的数组。我们可以先把每个数都减一，然后把最后一个数减一。

设 $d$ 为 $a$ 的差分数组，其中

$$
d[i] =
\begin{cases} 
a[i], & i=0     \\
a[i]-a[i-1], & i\ge 1     \\
\end{cases}
$$

由于全 $0$ 数组的差分数组也全为 $0$，所以把 $a$ 变成一个全为 $0$ 的数组，等价于把 $d$ 变成一个全为 $0$ 的数组。

根据前置知识，「子数组内的每个元素的值增加或减少 $1$」这个操作等价于修改差分数组**两个位置**上的数，一个加一，另一个减一。特别地，如果修改的是 $a$ 的后缀，那么操作等价于把差分数组中的一个数单独加一或者单独减一。

示例 1 的 $a=[1,1,1,2]$，差分数组 $d=[1,0,0,1]$，需要执行两次单独的减一操作。

示例 2 的 $a=[1,-2,2]$，差分数组 $d=[1,-3,4]$，这个要怎么操作？

- 贪心地想，由于每次操作可以执行一次加一和一次减一，那么选一个负数和一个正数操作是最优的。
- 执行 $3$ 次操作后可以把 $-3$ 变成 $0$，$4$ 变成 $1$。此时 $d=[1,0,1]$，和示例 1 相同，执行两次单独的减一操作。
- 一共执行 $3+2=5$ 次操作。

一般地，设 $d$ 中的正数之和为 $\textit{posSum}$，负数之和的**绝对值**为 $\textit{negSum}$。我们可以先执行 $\min(\textit{posSum},\textit{negSum})$ 次操作，让 $d$ 中只剩下正数或者只剩下负数；然后再执行 $|\textit{posSum}-\textit{negSum}|$ 次操作，让剩下的数都变成 $0$。总的操作次数为

$$
\min(\textit{posSum},\textit{negSum}) + |\textit{posSum}-\textit{negSum}|
$$

上式可以继续简化：

- 如果 $\textit{posSum} \ge \textit{negSum}$，那么上式为 $\textit{negSum} + (\textit{posSum}-\textit{negSum}) = \textit{posSum}$。
- 如果 $\textit{posSum} < \textit{negSum}$，那么上式为 $\textit{posSum} + (\textit{negSum}-\textit{posSum}) = \textit{negSum}$。

所以最终答案为

$$
\max(\textit{posSum},\textit{negSum})
$$

```py [sol-Python3]
class Solution:
    def minimumOperations(self, nums: List[int], target: List[int]) -> int:
        pos_sum = neg_sum = 0
        d = target[0] - nums[0]
        if d > 0:
            pos_sum = d
        else:
            neg_sum = -d
        for (a1, a2), (b1, b2) in pairwise(zip(nums, target)):
            d = (b2 - a2) - (b1 - a1)
            if d > 0:
                pos_sum += d
            else:
                neg_sum -= d
        return max(pos_sum, neg_sum)
```

```py [sol-Python3 写法二]
class Solution:
    def minimumOperations(self, nums: List[int], target: List[int]) -> int:
        pos_sum = neg_sum = 0
        for (a1, a2), (b1, b2) in pairwise(zip([0] + nums, [0] + target)):
            d = (b2 - a2) - (b1 - a1)
            if d > 0:
                pos_sum += d
            else:
                neg_sum -= d
        return max(pos_sum, neg_sum)
```

```java [sol-Java]
class Solution {
    public long minimumOperations(int[] nums, int[] target) {
        long posSum = 0;
        long negSum = 0;
        for (int i = 0; i < nums.length; i++) {
            int d = (target[i] - nums[i]) - (i > 0 ? target[i - 1] - nums[i - 1] : 0);
            if (d > 0) {
                posSum += d;
            } else {
                negSum -= d;
            }
        }
        return Math.max(posSum, negSum);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minimumOperations(vector<int>& nums, vector<int>& target) {
        long long pos_sum = 0, neg_sum = 0;
        for (int i = 0; i < nums.size(); i++) {
            int d = (target[i] - nums[i]) - (i ? target[i - 1] - nums[i - 1] : 0);
            if (d > 0) {
                pos_sum += d;
            } else {
                neg_sum -= d;
            }
        }
        return max(pos_sum, neg_sum);
    }
};
```

```go [sol-Go]
func minimumOperations(nums, target []int) int64 {
	posSum, negSum := 0, 0
	d := target[0] - nums[0]
	if d > 0 {
		posSum = d
	} else {
		negSum = -d
	}
	for i := 1; i < len(nums); i++ {
		d := (target[i] - nums[i]) - (target[i-1] - nums[i-1])
		if d > 0 {
			posSum += d
		} else {
			negSum -= d
		}
	}
	return int64(max(posSum, negSum))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

更多相似题目，见下面数据结构题单中的「**§2.1 一维差分**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
