操作相当于：选择一个连续子数组，去掉子数组中的除了子数组最大值以外的其他元素。

**定理**：如果在元素 $x=\textit{nums}[i]$ 的左边，有比 $x$ 大的数 $y$，那么 $x$ 必须去掉。

**证明**：反证法，假设能保留 $x$，那么必须去掉在 $x$ 左边的所有大于 $x$ 的数（不然就不是非递减的了）。去掉 $y$，说明在 $x$ 的左边还有 $\ge y$ 的其他数 $z$（去删除 $y$）。但 $z > x$，数组仍然是非递减的。如果继续去掉 $z$，说明还有 $\ge z$ 的其他数。依此类推，我们无法彻底去掉在 $x$ 左边的**所有**大于 $x$ 的数，矛盾。故原命题成立。

去掉这些（左边有更大元素的）$x$ 后，剩余元素是非递减的，都保留最好。

由于第一个数 $x=\textit{nums}[0]$ 的左边没有比 $x$ 大的数，所以第一个数可以保留。我们从 $\textit{mx} = \textit{nums}[0]$ 开始，向右遍历：

- 如果 $\textit{nums}[i] < \textit{mx}$，必须删除 $\textit{nums}[i]$。
- 否则，保留 $\textit{nums}[i]$，并更新 $\textit{mx} = \textit{nums}[i]$。
- 这个过程中保留的元素个数，即为答案。

```py [sol-Python3]
class Solution:
    def maximumPossibleSize(self, nums: List[int]) -> int:
        ans = mx = 0
        for x in nums:
            if x >= mx:
                mx = x
                ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int maximumPossibleSize(int[] nums) {
        int ans = 0;
        int mx = 0;
        for (int x : nums) {
            if (x >= mx) {
                mx = x;
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumPossibleSize(vector<int>& nums) {
        int ans = 0, mx = 0;
        for (int x : nums) {
            if (x >= mx) {
                mx = x;
                ans++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumPossibleSize(nums []int) (ans int) {
	mx := 0
	for _, x := range nums {
		if x >= mx {
			mx = x
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
