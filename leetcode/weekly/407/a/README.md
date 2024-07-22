**请先阅读**：[从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

从集合的角度理解，$k$ 必须是 $n$ 的**子集**。如果不是，返回 $-1$。怎么用位运算判断，见上面的文章链接。

如果 $k$ 是 $n$ 的子集，答案为从 $n$ 中去掉 $k$ 后的集合大小，即 $n\oplus k$ 的二进制中的 $1$ 的个数。

> 注：也可以计算 $n-k$ 的二进制中的 $1$ 的个数。

具体请看 [视频讲解](https://www.bilibili.com/video/BV16Z421N7P2/)，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def minChanges(self, n: int, k: int) -> int:
        return -1 if (n & k) != k else (n ^ k).bit_count()
```

```java [sol-Java]
class Solution {
    public int minChanges(int n, int k) {
        return (n & k) != k ? -1 : Integer.bitCount(n ^ k);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minChanges(int n, int k) {
        return (n & k) != k ? -1 : __builtin_popcount(n ^ k);
    }
};
```

```go [sol-Go]
func minChanges(n, k int) int {
	if n&k != k {
		return -1
	}
	return bits.OnesCount(uint(n ^ k))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

更多相似题目，见下面的位运算题单。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
