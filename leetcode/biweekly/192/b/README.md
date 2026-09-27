设 $\textit{source}$ 的总和为 $S$。

对于 $\textit{source}$ 中的两个数 $x$ 和 $y$，操作后变成了 $x+y - \textit{delta}$ 和 $\textit{delta}$，两数之和仍然为 $x+y$，$\textit{source}$ 的总和仍然为 $S$。所以操作不改变 $\textit{source}$ 的总和。

所以把 $\textit{source}$ 变成 $\textit{target}$ 的**必要条件**是两个数组的总和相等。

这个条件也是**充分**的，给出一个具体的操作方法即可：

如果 $\textit{source}[0] = \textit{target}[0]$，那么去掉下标 $0$，问题变成两个数组都有 $n-1$ 个元素的子问题。

否则，选择 $i=1,j=0,\textit{delta}=\textit{target}[0]$ 操作，可以把 $\textit{source}[0]$ 变成 $\textit{target}[0]$。问题变成两个数组都有 $n-1$ 个元素的子问题。

反复操作，直到两个数组都只剩下一个数。

如果两个数组的总和相等，那么最后剩下的数也相等。所以前述条件也是充分的。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def canTransform(self, source: list[int], target: list[int]) -> bool:
        return sum(source) == sum(target)
```

```java [sol-Java]
class Solution {
    public boolean canTransform(int[] source, int[] target) {
        long diff = 0;
        for (int i = 0; i < source.length; i++) {
            diff += source[i] - target[i];
        }
        return diff == 0;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool canTransform(vector<int>& source, vector<int>& target) {
        long long sum_s = reduce(source.begin(), source.end(), 0LL);
        long long sum_t = reduce(target.begin(), target.end(), 0LL);
        return sum_s == sum_t;
    }
};
```

```go [sol-Go]
func canTransform(source, target []int) bool {
	diff := 0
	for i, x := range source {
		diff += x - target[i]
	}
	return diff == 0
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{source}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面思维题单的「**§5.2 脑筋急转弯**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
