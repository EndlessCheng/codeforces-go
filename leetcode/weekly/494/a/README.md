如果 $\textit{nums}_1$ 全是奇数或者全是偶数，那么只用第一种操作 `nums2[i] = nums1[i]` 即可满足要求。

否则，$\textit{nums}_1$ 奇数偶数都有。由于偶数减去奇数等于奇数，那么随便选一个奇数 $x$，把每个偶数都减去 $x$（第二种操作），即可让所有偶数都变成奇数。其余每个奇数用第一种操作。

所以，一定可以满足题目要求，返回 $\texttt{true}$ 即可。

[本题视频讲解](https://www.bilibili.com/video/BV1vfAuzyEp8/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def uniformArray(self, _) -> bool:
        return True
```

```java [sol-Java]
class Solution {
    public boolean uniformArray(int[] nums1) {
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool uniformArray(vector<int>&) {
        return true;
    }
};
```

```go [sol-Go]
func uniformArray([]int) bool {
	return true
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

- [2396. 严格回文的数字](https://leetcode.cn/problems/strictly-palindromic-number/)
- [877. 石子游戏](https://leetcode.cn/problems/stone-game/)

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)
