怎么计算子数组的 OR？

首先，我们有如下 $\mathcal{O}(n^2)$ 的暴力算法：

从左到右正向遍历 $\textit{nums}$，对于 $x=\textit{nums}[i]$，从 $i-1$ 开始倒着遍历 $\textit{nums}[j]$，更新 $\textit{nums}[j]=\textit{nums}[j]\ \vert\ x$。

- $i=1$ 时，我们会把 $\textit{nums}[0]$ 到 $\textit{nums}[1]$ 的 OR 记录在 $\textit{nums}[0]$ 中。 
- $i=2$ 时，我们会把 $\textit{nums}[1]$ 到 $\textit{nums}[2]$ 的 OR 记录在 $\textit{nums}[1]$ 中，$\textit{nums}[0]$ 到 $\textit{nums}[2]$ 的 OR 记录在 $\textit{nums}[0]$ 中。
- $i=3$ 时，我们会把 $\textit{nums}[2]$ 到 $\textit{nums}[3]$ 的 OR 记录在 $\textit{nums}[2]$ 中；$\textit{nums}[1]$ 到 $\textit{nums}[3]$ 的 OR 记录在 $\textit{nums}[1]$ 中；$\textit{nums}[0]$ 到 $\textit{nums}[3]$ 的 OR 记录在 $\textit{nums}[0]$ 中。
- 按照该算法，可以计算出所有子数组的 OR。注意单个元素也算子数组。

下面来优化该算法。

前置知识：[从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

把二进制数看成集合，两个数的 OR 就是两个集合的**并集**。

对于两个二进制数 $a$ 和 $b$，如果 $a\ \vert\ b = a$，从集合的角度上看，$b$ 对应的集合是 $a$ 对应的集合的子集。

据此我们可以提出如下优化：

仍然是从左到右正向遍历 $\textit{nums}$，对于 $x=\textit{nums}[i]$，从 $i-1$ 开始倒着遍历 $\textit{nums}[j]$：
- 如果 $\textit{nums}[j]\ \vert\ x\ne\textit{nums}[j]$，说明 $\textit{nums}[j]$ 可以变大（求并集后，集合元素只会增多不会减少），更新 $\textit{nums}[j]=\textit{nums}[j]\ \vert\ x$。
- 否则 $\textit{nums}[j]\ \vert\ x=\textit{nums}[j]$，从集合的角度看，此时 $x$ 不仅是 $\textit{nums}[j]$ 的子集，同时也是 $\textit{nums}[k]\ (k<j)$ 的子集（因为前面的循环保证了每个集合都是其左侧相邻集合的子集），在 $B\subseteq A$ 的前提下，$A\cup B=A$，所以后续的循环都不会改变元素值，退出内层循环。
- 在循环中，用 $|\textit{nums}[j]-k|$ 更新答案的最小值。
- 注意单个元素也可以组成子数组，也要用 $|\textit{nums}[i]-k|$ 更新答案的最小值。

具体例子可以看 [视频讲解](https://www.bilibili.com/video/BV1Qx4y1E7zj/) 第四题（计算的是子数组 AND）。

```py [sol-Python3]
class Solution:
    def minimumDifference(self, nums: List[int], k: int) -> int:
        ans = inf
        for i, x in enumerate(nums):
            ans = min(ans, abs(x - k))
            j = i - 1
            while j >= 0 and nums[j] | x != nums[j]:
                nums[j] |= x
                ans = min(ans, abs(nums[j] - k))
                j -= 1
        return ans
```

```java [sol-Java]
class Solution {
    public int minimumDifference(int[] nums, int k) {
        int ans = Integer.MAX_VALUE;
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            ans = Math.min(ans, Math.abs(x - k));
            for (int j = i - 1; j >= 0 && (nums[j] | x) != nums[j]; j--) {
                nums[j] |= x;
                ans = Math.min(ans, Math.abs(nums[j] - k));
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumDifference(vector<int>& nums, int k) {
        int ans = INT_MAX;
        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];
            ans = min(ans, abs(x - k));
            for (int j = i - 1; j >= 0 && (nums[j] | x) != nums[j]; j--) {
                nums[j] |= x;
                ans = min(ans, abs(nums[j] - k));
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumDifference(nums []int, k int) int {
	ans := math.MaxInt
	for i, x := range nums {
		ans = min(ans, abs(x-k))
		for j := i - 1; j >= 0 && nums[j]|x != nums[j]; j-- {
			nums[j] |= x
			ans = min(ans, abs(nums[j]-k))
		}
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。由于 $2^{29}-1<10^9<2^{30}-1$，二进制数对应集合的大小不会超过 $29$，因此在 OR 运算下，每个数字至多可以增大 $29$ 次。总体上看，二重循环的总循环次数等于每个数字可以增大的次数之和，即 $O(n\log U)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

1. 把 OR 换成 AND 怎么做？[1521. 找到最接近目标值的函数值](https://leetcode.cn/problems/find-a-value-of-a-mysterious-function-closest-to-target/)
2. 把 OR 换成 GCD 怎么做？
3. 把 OR 换成 LCM 怎么做？

欢迎在评论区发表你的思路/代码。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
