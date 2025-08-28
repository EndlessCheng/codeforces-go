## 思路一：利用对称性化简

有效值为 $(\textit{nums}[i]|\textit{nums}[j])\&\textit{nums}[k]$。

分类讨论：

- 如果 $i\ne j$，那么 $(i,j,k)$ 和 $(j,i,k)$ 的有效值相等。由于两个相等的数的异或结果为 $0$，所以 $i\ne j$ 的情况对答案无贡献。
- 如果 $i = j\ne k$，那么 $(i,i,k)$ 和 $(k,k,i)$ 的有效值相等。由于两个相等的数的异或结果为 $0$，所以 $i = j\ne k$ 的情况对答案无贡献。
- 现在只剩下 $i = j = k$，此时 $(\textit{nums}[i]|\textit{nums}[j])\&\textit{nums}[k] = \textit{nums}[i]$。

所以答案为 $\textit{nums}$ 的异或和。

## 思路二：拆位法

思路一虽然简单，但并不通用。

比如说，把问题中的「有效值的异或和」改成「有效值的和」，要怎么做？

**拆位法**：由于或运算（$|$）和与运算（$\&$）是按位定义的，每一位互相独立，因此常把问题拆成按位统计。

每个比特位只有 $0$ 和 $1$，更容易计算。

有效值 $(a|b)\&c=0$ 时，对答案无贡献，所以只需统计有效值 $(a|b)\&c=1$ 的情况。

这意味着 $c$ 必须是 $1$，且 $a$ 和 $b$ 不能都是 $0$。

设有 $y$ 个 $1$，$x=n-y$ 个 $0$。

那么 $c$ 有 $y$ 个，$a|b$ 有 $n^2-x^2$ 个（任意选是 $n^2$，减去 $a$ 和 $b$ 都是 $0$ 的 $x^2$ 个）。

根据乘法原理，一共可以产生

$$
\textit{total} = (n^2-x^2)y = (n^2-(n-y)^2)y = (2ny-y^2)y = 2ny^2-y^3
$$

个 $1$。

> 如果把问题中的「有效值的异或和」改成「有效值的和」，可以枚举比特位 $i$，计算 $\textit{total}\cdot 2^i$ 的和。

$\textit{total}$ 个 $1$ 的异或结果为 $\textit{total}\bmod 2$，而 $2ny^2$ 是偶数，异或结果一定是 $0$，可以去掉。所以只需看 $y^3$ 的奇偶性，也就是 $y$ 的奇偶性。

如果 $y$ 是奇数，那么这个比特位的异或值就是 $1$，否则是 $0$。

这实际上就是看每个比特位的异或值是否为 $1$。

计算 $\textit{nums}$ 的异或和，就是答案。

[视频讲解](https://www.bilibili.com/video/BV1i24y1e7E7/) 第三题。

```py [sol-Python3]
class Solution:
    def xorBeauty(self, nums: List[int]) -> int:
        return reduce(xor, nums)
```

```java [sol-Java]
class Solution {
    public int xorBeauty(int[] nums) {
        int ans = 0;
        for (int x : nums) ans ^= x;
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int xorBeauty(vector<int>& nums) {
        int ans = 0;
        for (int x : nums) ans ^= x;
        return ans;
    }
};
```

```go [sol-Go]
func xorBeauty(nums []int) (ans int) {
	for _, x := range nums {
		ans ^= x
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
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
