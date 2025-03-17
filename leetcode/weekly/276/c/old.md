## 解法一：倒序 DP（填表法）

填表法适用于大多数 DP：通过当前状态所依赖的状态，来计算当前状态。

设有 $n$ 个问题，定义 $f[i]$ 表示解决区间 $[i,n-1]$ 内的问题可以获得的最高分数。

倒序遍历问题列表，对于第 $i$ 个问题，我们有两种决策：跳过或解决。

- **跳过**，则有 $f[i]=f[i+1]$。
- **解决**，则需要跳过后续 $\textit{brainpower}[i]$ 个问题。记 $j=i+\textit{brainpower}[i]+1$，则有

$$
f[i] =
\begin{cases}
\textit{point}[i]+f[j],&j<n\\
\textit{point}[i],&j\ge n
\end{cases}
$$

这两种决策取最大值。

最后答案为 $f[0]$。

```py [sol-Python3]
class Solution:
    def mostPoints(self, questions: List[List[int]]) -> int:
        n = len(questions)
        f = [0] * (n + 1)
        for i in range(n - 1, -1, -1):
            point, brainpower = questions[i]
            j = i + brainpower + 1
            f[i] = max(f[i + 1], point + (f[j] if j < n else 0))
        return f[0]
```

```java [sol-Java]
class Solution {
    public long mostPoints(int[][] questions) {
        int n = questions.length;
        long[] f = new long[n + 1];
        for (int i = n - 1; i >= 0; i--) {
            int[] q = questions[i];
            int j = i + q[1] + 1;
            f[i] = Math.max(f[i + 1], q[0] + (j < n ? f[j] : 0));
        }
        return f[0];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long mostPoints(vector<vector<int>>& questions) {
        int n = questions.size();
        vector<long long> f(n + 1);
        for (int i = n - 1; i >= 0; i--) {
            auto& q = questions[i];
            int j = i + q[1] + 1;
            f[i] = max(f[i + 1], q[0] + (j < n ? f[j] : 0));
        }
        return f[0];
    }
};
```

```go [sol-Go]
func mostPoints(questions [][]int) int64 {
	n := len(questions)
	f := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		q := questions[i]
		if j := i + q[1] + 1; j < n {
			f[i] = max(f[i+1], q[0]+f[j])
		} else {
			f[i] = max(f[i+1], q[0])
		}
	}
	return int64(f[0])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 解法二：正序 DP（刷表法）

另一种做法是刷表法：用当前状态，去更新当前状态所影响的状态。

定义 $f[i]$ 表示在可以解决问题 $i$ 时，解决区间 $[0,i)$ 内的问题可以获得的最高分数。

对于问题 $i$，若跳过，则可以更新 $f[i+1]=\max(f[i+1],f[i])$。

若不跳过，记 $j=i+\textit{brainpower}[i]+1$，则可以更新 $f[j]=\max(f[j],f[i]+\textit{point}[i])$。

对于 $j\ge n$ 的情况，为了简化代码逻辑，我们可以将其更新到 $f[n]$ 中。

最后答案为 $f[n]$。

```py [sol-Python3]
class Solution:
    def mostPoints(self, questions: List[List[int]]) -> int:
        n = len(questions)
        f = [0] * (n + 1)
        for i, (point, brainpower) in enumerate(questions):
            f[i + 1] = max(f[i + 1], f[i])
            j = min(i + brainpower + 1, n)
            f[j] = max(f[j], f[i] + point)
        return f[n]
```

```java [sol-Java]
class Solution {
    public long mostPoints(int[][] questions) {
        int n = questions.length;
        long[] f = new long[n + 1];
        for (int i = 0; i < n; i++) {
            f[i + 1] = Math.max(f[i + 1], f[i]);
            int[] q = questions[i];
            int j = Math.min(i + q[1] + 1, n);
            f[j] = Math.max(f[j], f[i] + q[0]);
        }
        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long mostPoints(vector<vector<int>>& questions) {
        int n = questions.size();
        vector<long long> f(n + 1);
        for (int i = 0; i < n; i++) {
            f[i + 1] = max(f[i + 1], f[i]);
            auto& q = questions[i];
            int j = min(i + q[1] + 1, n);
            f[j] = max(f[j], f[i] + q[0]);
        }
        return f[n];
    }
};
```

```go [sol-Go]
func mostPoints(questions [][]int) int64 {
	n := len(questions)
	f := make([]int, n+1)
	for i, q := range questions {
		f[i+1] = max(f[i+1], f[i])
		j := i + q[1] + 1
		if j > n {
			j = n
		}
		f[j] = max(f[j], f[i]+q[0])
	}
	return int64(f[n])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

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
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
