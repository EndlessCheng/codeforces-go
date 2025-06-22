**前置知识**：[完全背包【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)

本题是一个反向构造题：给你完全背包的 DP 数组 $\textit{numWays}$，已知 $\textit{numWays}$ 是由数组 $a$ 算出来的（算方案数），请你还原数组 $a$。

想一想，$\textit{numWays}$ 中最小非零元素，意味着什么？

看示例 1 中，$\textit{numWays} = [0,1,0,2,0,3,0,4,0,5]$。

其中最小非零元素为 $\textit{numWays}[2]=1$（注意下标从 $1$ 开始），说明 $a$ 中没有比 $2$ 小的数（否则最小非零元素的下标比 $2$ 小），且元素和为 $2$ 的方案数有一种，说明 $a$ 中有 $2$，这个 $2$ 可以单独一个数，贡献一个和为 $2$ 的方案。

用这个 $2$ 去计算一个新的完全背包数组 $f$（下标从 $0$ 开始），我们得到 $f = [1,0,1,0,1,0,1,0,1,0,1]$，其中 $f[i]$ 对应和为 $i$ 的方案数。特别地，$f[0]=1$ 对应和为 $0$ 的方案数。

继续向后遍历 $\textit{numWays}$：

- 如果 $\textit{numWays}[i] = f[i]+1$，这意味着 $i$ 可以单独一个数，贡献一个和为 $i$ 的方案，所以 $a$ 中一定有 $i$。比如 $i=4$ 符合要求。
- 如果 $\textit{numWays}[i] = f[i]$，说明 $a$ 中没有 $i$。
- 其他情况：矛盾，$\textit{numWays}$ 不合法。

用这个 $4$ 去更新 $f$，得到 $f =[1,0,1,0,2,0,2,0,3,0,3]$。

继续向后遍历 $\textit{numWays}$，发现 $\textit{numWays}[6] = f[6]+1$，这意味着 $6$ 可以单独一个数，贡献一个和为 $4$ 的方案，所以 $a$ 中一定有 $6$。

用这个 $6$ 去更新 $f$，得到 $f =[1,0,1,0,2,0,3,0,4,0,5]$。

继续向后遍历 $\textit{numWays}$，没有 $\textit{numWays}[i] = f[i]+1$。

最终 $a=[2,4,6]$。

**细节**：为避免方案数溢出，可以把 $f[i]$ 与一个 $\ge \max(\textit{numWays})+2$ 的数取最小值。也可以在发现 $f[i]\ge \max(\textit{numWays})+2$ 时返回空数组。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def findCoins(self, numWays: List[int]) -> List[int]:
        mx = max(numWays)
        n = len(numWays)
        f = [1] + [0] * n
        ans = []
        for i, ways in enumerate(numWays, 1):
            if ways == f[i]:
                continue
            if ways - 1 != f[i]:
                return []
            ans.append(i)
            # 现在得到了一个大小为 i 的物品，用 i 计算完全背包（空间优化写法）
            for j in range(i, n + 1):
                f[j] += f[j - i]
                if f[j] > mx + 1:  # 不合法
                    return []
        return ans
```

```java [sol-Java]
class Solution {
    public List<Integer> findCoins(int[] numWays) {
        int n = numWays.length;
        int[] f = new int[n + 1];
        f[0] = 1;
        List<Integer> ans = new ArrayList<>();
        for (int i = 1; i <= n; i++) {
            int ways = numWays[i - 1];
            if (ways == f[i]) {
                continue;
            }
            if (ways - 1 != f[i]) {
                return List.of();
            }
            ans.add(i);
            // 现在得到了一个大小为 i 的物品，用 i 计算完全背包
            for (int j = i; j <= n; j++) {
                f[j] = Math.min(f[j] + f[j - i], Integer.MAX_VALUE / 2); // 防止溢出
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> findCoins(vector<int>& numWays) {
        int n = numWays.size();
        vector<int> f(n + 1);
        f[0] = 1;
        vector<int> ans;
        for (int i = 1; i <= n; i++) {
            int ways = numWays[i - 1];
            if (ways == f[i]) {
                continue;
            }
            if (ways - 1 != f[i]) {
                return {};
            }
            ans.push_back(i);
            // 现在得到了一个大小为 i 的物品，用 i 计算完全背包
            for (int j = i; j <= n; j++) {
                f[j] = min(f[j] + f[j - i], INT_MAX / 2); // 防止溢出
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func findCoins(numWays []int) (ans []int) {
	n := len(numWays)
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		ways := numWays[i-1]
		if ways == f[i] {
			continue
		}
		if ways-1 != f[i] {
			return nil
		}
		ans = append(ans, i)
		// 现在得到了一个大小为 i 的物品，用 i 计算完全背包（空间优化写法）
		for j := i; j <= n; j++ {
			f[j] = min(f[j]+f[j-i], math.MaxInt/2) // 防止溢出
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{numWays}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

1. 动态规划题单的「**§3.2 完全背包**」。
2. 贪心与思维题单的「**六、构造题**」。

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
