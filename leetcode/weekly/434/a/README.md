设 $\textit{nums}$ 的元素和为 $S$，左子数组元素和为 $L$，那么右子数组的元素和为 $S-L$。

题目要求 $L - (S-L) = 2L - S$ 是偶数。由于 $2L$ 一定是偶数，所以问题变成判断 $S$ 是否为偶数。

注意这和 $i$ 无关。换句话说，如果 $S$ 是奇数，那么答案是 $0$；如果 $S$ 是偶数，那么所有分区方案都符合要求，答案是 $n-1$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV15sFNewEia/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countPartitions(self, nums: List[int]) -> int:
        return 0 if sum(nums) % 2 else len(nums) - 1
```

```java [sol-Java]
class Solution {
    public int countPartitions(int[] nums) {
        int s = Arrays.stream(nums).sum();
        return s % 2 > 0 ? 0 : nums.length - 1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countPartitions(vector<int>& nums) {
        int s = reduce(nums.begin(), nums.end());
        return s % 2 ? 0 : nums.size() - 1;
    }
};
```

```go [sol-Go]
func countPartitions(nums []int) int {
	s := 0
	for _, x := range nums {
		s += x
	}
	if s%2 == 0 {
		return len(nums) - 1
	}
	return 0
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

更多相似题目，见下面贪心题单中的「**§5.2 脑筋急转弯**」。

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
10. 【本题相关】[贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
