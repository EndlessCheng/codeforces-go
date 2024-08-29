## 提示 1

对于一个栈（数组），我们只能移除其前缀。

## 提示 2

对每个栈求其前缀和 $\textit{sum}$，$\textit{sum}$ 的第 $j$ 个元素视作一个体积为 $j$，价值为 $\textit{sum}[j]$ 的物品。

问题转化成求从 $n$ 个物品组里面取物品体积和为 $k$ 的物品，且每组至多取一个物品时的物品价值最大和，即分组背包模型。

## 思路

定义 $f[i][j]$ 表示从前 $i$ 个组取体积之和为 $j$ 的物品时，物品价值之和的最大值。

枚举第 $i$ 个组的所有物品，设当前物品体积为 $w$，价值为 $v$，用 $f[i-1][j-w]+v$ 更新 $f[i][j]$ 的最大值。

答案为 $f[n][k]$。

代码实现时，可以仿照 0-1 背包的写法，将第一维压缩掉。

```py [sol-Python3]
class Solution:
    def maxValueOfCoins(self, piles: List[List[int]], k: int) -> int:
        f = [0] * (k + 1)
        sum_n = 0
        for pile in piles:
            n = len(pile)
            for i in range(1, n):
                pile[i] += pile[i - 1]  # pile 前缀和
            sum_n = min(sum_n + n, k)  # 优化：j 从前 i 个栈的大小之和开始枚举（不超过 k）
            for j in range(sum_n, 0, -1):
                f[j] = max(f[j], max(f[j - w - 1] + pile[w] for w in range(min(n, j))))  # w 从 0 开始，物品体积为 w+1
        return f[k]
```

```java [sol-Java]
class Solution {
    public int maxValueOfCoins(List<List<Integer>> piles, int k) {
        int[] f = new int[k + 1];
        int sumN = 0;
        for (List<Integer> pile : piles) {
            int n = pile.size();
            for (int i = 1; i < n; i++) {
                pile.set(i, pile.get(i) + pile.get(i - 1)); // pile 前缀和
            }
            sumN = Math.min(sumN + n, k); // 优化：j 从前 i 个栈的大小之和开始枚举（不超过 k）
            for (int j = sumN; j > 0; j--) {
                for (var w = 0; w < Math.min(n, j); w++) {
                    f[j] = Math.max(f[j], f[j - w - 1] + pile.get(w)); // w 从 0 开始，物品体积为 w+1
                }
            }
        }
        return f[k];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxValueOfCoins(vector<vector<int>>& piles, int k) {
        vector<int> f(k + 1);
        int sum_n = 0;
        for (auto& pile : piles) {
            int n = pile.size();
            for (int i = 1; i < n; i++) {
                pile[i] += pile[i - 1]; // pile 前缀和
            }
            sum_n = min(sum_n + n, k); // 优化：j 从前 i 个栈的大小之和开始枚举（不超过 k）
            for (int j = sum_n; j; j--) {
                for (int w = 0; w < min(n, j); w++) {
                    f[j] = max(f[j], f[j - w - 1] + pile[w]); // w 从 0 开始，物品体积为 w+1
                }
            }
        }
        return f[k];
    }
};
```

```go [sol-Go]
func maxValueOfCoins(piles [][]int, k int) int {
	f := make([]int, k+1)
	sumN := 0
	for _, pile := range piles {
		n := len(pile)
		for i := 1; i < n; i++ {
			pile[i] += pile[i-1] // pile 前缀和
		}
		sumN = min(sumN+n, k) // 优化：j 从前 i 个栈的大小之和开始枚举（不超过 k）
		for j := sumN; j > 0; j-- {
			for w, v := range pile[:min(n, j)] {
				f[j] = max(f[j], f[j-w-1]+v) // w 从 0 开始，物品体积为 w+1
			}
		}
	}
	return f[k]
}
```

#### 复杂度分析

时间复杂度：$\mathcal{O}(ks)$。将外层循环与最内层循环合并，即每个栈的大小之和，记作 $\textit{s}$，算上中间这层的循环，时间复杂度为 $\mathcal{O}(ks)$。
空间复杂度：$\mathcal{O}(k)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
