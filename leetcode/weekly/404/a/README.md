## 方法一：枚举

用一个长为 $2$ 的 $\textit{cnt}$ 数组，分别记录偶数行和奇数行分别放了多少个球。

从 $i=1$ 开始枚举，如果 $i$ 是奇数则把 $\textit{cnt}[1]$ 增加 $i$，如果 $i$ 是偶数则把 $\textit{cnt}[0]$ 增加 $i$。

增加后，如果

- $\textit{cnt}[0] > \textit{red}$ 或者 $\textit{cnt}[1] > \textit{blue}$
- $\textit{cnt}[0] > \textit{blue}$ 或者 $\textit{cnt}[1] > \textit{red}$

这两个条件都成立，说明无法把第 $i$ 行填满，返回 $i-1$。

```py [sol-Python3]
class Solution:
    def maxHeightOfTriangle(self, red: int, blue: int) -> int:
        cnt = [0, 0]
        for i in count(1):
            cnt[i % 2] += i
            if (cnt[0] > red or cnt[1] > blue) and (cnt[0] > blue or cnt[1] > red):
                return i - 1
```

```java [sol-Java]
class Solution {
    public int maxHeightOfTriangle(int red, int blue) {
        int[] cnt = new int[2];
        for (int i = 1; ; i++) {
            cnt[i % 2] += i;
            if ((cnt[0] > red || cnt[1] > blue) && (cnt[0] > blue || cnt[1] > red)) {
                return i - 1;
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxHeightOfTriangle(int red, int blue) {
        int cnt[2]{};
        for (int i = 1; ; i++) {
            cnt[i % 2] += i;
            if ((cnt[0] > red || cnt[1] > blue) && (cnt[0] > blue || cnt[1] > red)) {
                return i - 1;
            }
        }
    }
};
```

```go [sol-Go]
func maxHeightOfTriangle(red, blue int) int {
	cnt := [2]int{}
	for i := 1; ; i++ {
		cnt[i%2] += i
		if (cnt[0] > red || cnt[1] > blue) && (cnt[0] > blue || cnt[1] > red) {
			return i - 1
		}
	}
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\min(\sqrt{\textit{red}}, \sqrt{\textit{blue}}))$。理由见方法二的分析。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：数学公式

### 核心思路

奇数行放红球，偶数行放蓝球；或者奇数行放蓝球，偶数行放红球。

计算最多能放多少排。两种情况取最大值。

### 奇数行

设奇数行有 $k$ 行，那么需要

$$
1+3+5+\cdots + (2k-1) = k^2
$$

个球。（等差数列求和公式）

假设我们有 $n$ 个球，那么有

$$
n\ge k^2
$$

解得

$$
k \le \left\lfloor\sqrt n\right\rfloor
$$

### 偶数行

设偶数行有 $k$ 行，那么需要

$$
2+4+6+\cdots + 2k = k^2 + k
$$

个球。（等差数列求和公式）

假设我们有 $n$ 个球，那么有

$$
n\ge k^2 + k
$$

解得

$$
k \le \left\lfloor\dfrac{\sqrt{4n+1}-1}{2}\right\rfloor = \left\lfloor\dfrac{\lfloor\sqrt{4n+1}\rfloor-1}{2}\right\rfloor
$$

### 答案

设有 $\textit{odd}$ 个奇数行，$\textit{even}$ 个偶数行，那么总行数为

$$
\begin{cases}
2\cdot \textit{even} + 1, & odd > even      \\
2\cdot \textit{odd}, & \text{otherwise}     \\
\end{cases}
$$

具体请看 [视频讲解](https://www.bilibili.com/video/BV16w4m1e7y3/)，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def maxHeightOfTriangle(self, red: int, blue: int) -> int:
        def f(n: int, m: int) -> int:
            odd = isqrt(n)
            even = (isqrt(m * 4 + 1) - 1) // 2
            return even * 2 + 1 if odd > even else odd * 2
        return max(f(red, blue), f(blue, red))
```

```java [sol-Java]
class Solution {
    public int maxHeightOfTriangle(int red, int blue) {
        return Math.max(f(red, blue), f(blue, red));
    }

    private int f(int n, int m) {
        int odd = (int) Math.sqrt(n);
        int even = ((int) Math.sqrt(m * 4 + 1) - 1) / 2;
        return odd > even ? even * 2 + 1 : odd * 2;
    }
}
```

```cpp [sol-C++]
class Solution {
    int f(int n, int m) {
        int odd = sqrt(n);
        int even = ((int) sqrt(m * 4 + 1) - 1) / 2;
        return odd > even ? even * 2 + 1 : odd * 2;
    }

public:
    int maxHeightOfTriangle(int red, int blue) {
        return max(f(red, blue), f(blue, red));
    }
};
```

```go [sol-Go]
func f(n, m int) int {
	odd := int(math.Sqrt(float64(n)))
	even := (int(math.Sqrt(float64(m*4+1))) - 1) / 2
	if odd > even {
		return even*2 + 1
	}
	return odd * 2
}

func maxHeightOfTriangle(red, blue int) int {
	return max(f(red, blue), f(blue, red))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。CPU 有专门的计算平方根的指令，可以视作 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

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
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
