## 运动员的视角

对于任意一种匹配方案，总是可以把其中的运动员改成能力值更小的没有匹配的运动员，仍然满足题目要求。

所以可以只考虑 $\textit{players}$ 中的能力值最小的一部分人。

为方便计算，把 $\textit{players}$ 和 $\textit{trainers}$ 从小到大排序。参与匹配的运动员是 $\textit{players}$ 的前缀。

从能力值最小的运动员 $\textit{player}[0]$ 开始思考，他应当匹配训练能力值大于等于 $\textit{player}[0]$ 且最接近 $\textit{player}[0]$ 的训练师（如果选了一个训练能力值更大的，可能会导致能力值更大的运动员无法匹配）。找到训练师 $\textit{trainers}[j]$ 后，下一个与 $\textit{players}[1]$ 匹配的训练师在 $j$ 右边，所以我们可以用双指针做。

```py [sol-Python3]
class Solution:
    def matchPlayersAndTrainers(self, players: List[int], trainers: List[int]) -> int:
        players.sort()
        trainers.sort()
        j, m = 0, len(trainers)
        for i, p in enumerate(players):
            while j < m and trainers[j] < p:
                j += 1
            if j == m:  # 无法找到匹配的训练师
                return i
            j += 1  # 匹配一位训练师
        return len(players)  # 所有运动员都有匹配的训练师
```

```java [sol-Java]
class Solution {
    public int matchPlayersAndTrainers(int[] players, int[] trainers) {
        Arrays.sort(players);
        Arrays.sort(trainers);
        int n = players.length;
        int m = trainers.length;
        int j = 0;
        for (int i = 0; i < n; i++) {
            int p = players[i];
            while (j < m && trainers[j] < p) {
                j++;
            }
            if (j == m) { // 无法找到匹配的训练师
                return i;
            }
            j++; // 匹配一位训练师
        }
        return n; // 所有运动员都有匹配的训练师
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int matchPlayersAndTrainers(vector<int>& players, vector<int>& trainers) {
        ranges::sort(players);
        ranges::sort(trainers);
        int n = players.size(), m = trainers.size();
        int j = 0;
        for (int i = 0; i < n; i++) {
            while (j < m && trainers[j] < players[i]) {
                j++;
            }
            if (j == m) { // 无法找到匹配的训练师
                return i;
            }
            j++; // 匹配一位训练师
        }
        return n; // 所有运动员都有匹配的训练师
    }
};
```

```go [sol-Go]
func matchPlayersAndTrainers(players, trainers []int) int {
	slices.Sort(players)
	slices.Sort(trainers)
	j, m := 0, len(trainers)
	for i, p := range players {
		for j < m && trainers[j] < p {
			j++
		}
		if j == m { // 无法找到匹配的训练师
			return i
		}
		j++ // 匹配一位训练师
	}
	return len(players) // 所有运动员都有匹配的训练师
}
```

## 训练师的视角

也可以遍历训练师 $\textit{trainers}$，去找对应的运动员 $\textit{players}$：

- 初始化 $i=0$，$j=0$。
- 如果 $\textit{trainers}[i] < \textit{players}[j]$，不匹配，把 $i$ 加一，考虑下一个训练师能否匹配（训练能力值更大）。
- 如果 $\textit{trainers}[i] \ge  \textit{players}[j]$，匹配，把 $i$ 和 $j$ 都加一。

**注**：把 $\textit{players}$ 视作子序列，这个做法类似在 $\textit{trainers}$ 中寻找子序列 $\textit{players}$。读者可以对比下面的代码和 [392. 判断子序列](https://leetcode.cn/problems/is-subsequence/) 的代码。

```py [sol-Python3]
class Solution:
    def matchPlayersAndTrainers(self, players: List[int], trainers: List[int]) -> int:
        players.sort()
        trainers.sort()
        j, m = 0, len(players)
        for t in trainers:
            if j < m and players[j] <= t:
                j += 1
        return j
```

```java [sol-Java]
class Solution {
    public int matchPlayersAndTrainers(int[] players, int[] trainers) {
        Arrays.sort(players);
        Arrays.sort(trainers);
        int m = players.length;
        int j = 0;
        for (int t : trainers) {
            if (j < m && players[j] <= t) {
                j++;
            }
        }
        return j;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int matchPlayersAndTrainers(vector<int>& players, vector<int>& trainers) {
        ranges::sort(players);
        ranges::sort(trainers);
        int m = players.size();
        int j = 0;
        for (int t : trainers) {
            if (j < m && players[j] <= t) {
                j++;
            }
        }
        return j;
    }
};
```

```go [sol-Go]
func matchPlayersAndTrainers(players, trainers []int) int {
	slices.Sort(players)
	slices.Sort(trainers)
	j, m := 0, len(players)
	for _, t := range trainers {
		if j < m && players[j] <= t {
			j++
		}
	}
	return j
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + m\log m)$，其中 $n$ 为 $\textit{players}$ 的长度，$m$ 为 $\textit{trainers}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$，忽略排序时的栈开销。

## 总结

训练师的视角不需要写二重循环，更简单。

这类问题类似「子序列匹配」，我们要在 $\textit{trainers}$ 中寻找子序列 $\textit{players}$，通常的子序列匹配是 $\textit{trainers}[i]=\textit{players}[j]$，本题是 $\textit{trainers}[i]\ge \textit{players}[j]$。对于这类问题，我的经验是外层循环遍历相对来说更长的 $\textit{trainers}$：不匹配就继续循环，匹配就把 $j$ 加一，继续循环。如果外层循环遍历更短的子序列 $\textit{players}$，就需要写二重循环了。

## 专题训练

贪心题单的「**§1.3 双序列配对**」。

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

