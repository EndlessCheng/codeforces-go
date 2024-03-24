请看 [视频讲解](https://www.bilibili.com/video/BV1wr421h7xY/) 第二题。

## 方法一：贪心 + 枚举

设当前数组最大值为 $m$，对它做加一操作更好（因为复制最大值最优）。

- 如果先复制 $m$，再加一，那么元素和增加了 $m+1$。
- 如果先加一，再复制 $m+1$，那么元素和增加了 $m+2$。

所以，先加一再复制更优。

所以，加一操作都应当在复制操作之前。

我们可以枚举加一操作执行了 $\textit{add}= 0,1,2,\cdots, k-1$ 次。

设 $m=1+\textit{add}$，我们还需要复制

$$
\left\lceil\dfrac{k}{m}\right\rceil-1 = \left\lfloor\dfrac{k-1}{m}\right\rfloor
$$

次，才能让元素和至少为 $k$。上式可以分类讨论 $k$ 是 $m$ 的倍数，和 $k$ 不是 $m$ 的倍数两种情况证明。

所以答案为

$$
\min\limits_{m=1}^{k} m-1 + \left\lfloor\dfrac{k-1}{m}\right\rfloor
$$

```py [sol-Python3]
class Solution:
    def minOperations(self, k: int) -> int:
        return min(m - 1 + (k - 1) // m for m in range(1, k + 1))
```

```java [sol-Java]
class Solution {
    public int minOperations(int k) {
        int ans = Integer.MAX_VALUE;
        for (int m = 1; m <= k; m++) {
            ans = Math.min(ans, m - 1 + (k - 1) / m);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(int k) {
        int ans = INT_MAX;
        for (int m = 1; m <= k; m++) {
            ans = min(ans, m - 1 + (k - 1) / m);
        }
        return ans;
    }
};
```

```go [sol-Go]
func minOperations(k int) int {
	ans := math.MaxInt
	for m := 1; m <= k; m++ {
		ans = min(ans, m-1+(k-1)/m)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(k)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：数学

由基本不等式，或者对勾函数性质可知，设 $\textit{rt} = \left\lfloor\sqrt{k-1}\right\rfloor$，那么当 $m$ 取 $\textit{rt}$ 或者 $\textit{rt}+1$ 时我们可以得到最小值。

为防止 $m=0$，可以和 $1$ 取最大值（或者特判）。

> 注：在本题数据范围下，开平方结果的整数部分是正确的，无需调整。

```py [sol-Python3]
class Solution:
    def minOperations(self, k: int) -> int:
        rt = max(isqrt(k - 1), 1)
        return min(rt - 1 + (k - 1) // rt, rt + (k - 1) // (rt + 1))
```

```java [sol-Java]
class Solution {
    public int minOperations(int k) {
        int rt = Math.max((int) Math.sqrt(k - 1), 1);
        return Math.min(rt - 1 + (k - 1) / rt, rt + (k - 1) / (rt + 1));
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(int k) {
        int rt = max((int) sqrt(k - 1), 1);
        return min(rt - 1 + (k - 1) / rt, rt + (k - 1) / (rt + 1));
    }
};
```

```go [sol-Go]
func minOperations(k int) int {
	rt := max(int(math.Sqrt(float64(k-1))), 1)
	return min(rt-1+(k-1)/rt, rt+(k-1)/(rt+1))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。开平方有专门的 CPU 指令，可以视作 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。


## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。
