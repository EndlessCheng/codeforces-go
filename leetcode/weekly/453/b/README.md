题目说：用计算机 $j$ 解锁计算机 $i$ 的前提是 $j<i$ 且 $\textit{complexity}[j] < \textit{complexity}[i]$。

**观察**：

- 一开始就解锁的只有计算机 $0$。
- 第一轮，被 $0$ 解锁的计算机（记作集合 $A$），密码复杂度比 $\textit{complexity}[0]$ 大。
- 第二轮，被集合 $A$ 中的计算机解锁的计算机（记作集合 $B$），密码复杂度更大，所以也比 $\textit{complexity}[0]$ 大。
- 第三轮，被集合 $B$ 中的计算机解锁的计算机（记作集合 $C$），密码复杂度更大，所以也比 $\textit{complexity}[0]$ 大。
- 依此类推，所有被解锁的计算机的密码复杂度都要比 $\textit{complexity}[0]$ 大。

**定理**：当且仅当计算机 $0$ 右边的所有计算机的密码复杂度都比 $\textit{complexity}[0]$ 大，才能解锁所有计算机。

**充分性**：如果计算机 $0$ 右边的所有计算机的密码复杂度都比 $\textit{complexity}[0]$ 大，根据题意，仅用计算机 $0$ 便可解锁所有计算机。

**必要性**：如果可以解锁所有的计算机，那么计算机 $0$ 右边的所有计算机的密码复杂度都比 $\textit{complexity}[0]$ 大。考虑其逆否命题，即如果在计算机 $0$ 的右边存在计算机 $A$，满足 $\textit{complexity}[i] \le \textit{complexity}[0]$，那么不可能解锁计算机 $A$，更不可能解锁所有计算机。为了解锁计算机 $A$，我们需要在其左边找比 $\textit{complexity}[i]$ 更小的计算机。不断往左找，计算机的密码复杂度只会更小，直到找到一台被计算机 $0$ 解锁的计算机 $B$。$B$ 的密码复杂度必须比 $\textit{complexity}[0]$ 大，但为了能解锁计算机 $A$，$B$ 的密码复杂度又要 $< \textit{complexity}[i] \le \textit{complexity}[0]$，矛盾，所以不可能解锁计算机 $A$。

根据定理，如果计算机 $0$ 右边的所有计算机的密码复杂度都比 $\textit{complexity}[0]$ 大，那么我们可以按照任意顺序解锁这 $n-1$ 台计算机，方案数为 $n-1$ 个不同物品的全排列个数，即 

$$
(n-1)!
$$

注意取模。关于模运算的知识点，见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

```py [sol-Python3]
class Solution:
    def countPermutations(self, complexity: List[int]) -> int:
        MOD = 1_000_000_007
        ans = 1
        for i in range(1, len(complexity)):
            if complexity[i] <= complexity[0]:
                return 0
            ans = ans * i % MOD
        return ans
```

```java [sol-Java]
class Solution {
    public int countPermutations(int[] complexity) {
        final int MOD = 1_000_000_007;
        long ans = 1;
        for (int i = 1; i < complexity.length; i++) {
            if (complexity[i] <= complexity[0]) {
                return 0;
            }
            ans = ans * i % MOD;
        }
        return (int) ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countPermutations(vector<int>& complexity) {
        const int MOD = 1'000'000'007;
        long long ans = 1;
        for (int i = 1; i < complexity.size(); i++) {
            if (complexity[i] <= complexity[0]) {
                return 0;
            }
            ans = ans * i % MOD;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countPermutations(complexity []int) int {
	const mod = 1_000_000_007
	ans := 1
	for i := 1; i < len(complexity); i++ {
		if complexity[i] <= complexity[0] {
			return 0
		}
		ans = ans * i % mod
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{complexity}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

更多相似题目，见下面思维题单的「**§5.2 脑筋急转弯**」。

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
