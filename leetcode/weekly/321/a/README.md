## 方法一

$1$ 到 $x$ 的元素和为 $\dfrac{x(x+1)}{2}$，$x$ 到 $n$ 的元素和为 $1$ 到 $n$ 的元素和减去 $1$ 到 $x-1$ 的元素和，即 $\dfrac{n(n+1)-x(x-1)}{2}$。

两式相等，简化后即

$$
x = \sqrt{\dfrac{n(n+1)}{2}}
$$

如果 $x$ 不是整数则返回 $-1$。

[本题视频讲解](https://www.bilibili.com/video/BV1sD4y1e7pr/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def pivotInteger(self, n: int) -> int:
        m = n * (n + 1) // 2
        x = isqrt(m)
        return x if x * x == m else -1
```

```java [sol-Java]
class Solution {
    public int pivotInteger(int n) {
        int m = n * (n + 1) / 2;
        int x = (int) Math.sqrt(m);
        return x * x == m ? x : -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int pivotInteger(int n) {
        int m = n * (n + 1) / 2;
        int x = sqrt(m);
        return x * x == m ? x : -1;
    }
};
```

```go [sol-Go]
func pivotInteger(n int) int {
	m := n * (n + 1) / 2
	x := int(math.Sqrt(float64(m)))
	if x*x == m {
		return x
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。计算平方根有专门的 CPU 指令，可以视作是 $\mathcal{O}(1)$ 时间。Python 的 `math.isqrt` 用的牛顿迭代法，这里也视作 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二

注意到 $\dfrac{n(n+1)}{2}$ 同时也是完全平方数的情况应该是比较少的（见 [OEIS A001108](https://oeis.org/A001108)）。在本题数据范围下，$n$ 只有

$$
1,8,49,288
$$

这四个，对应的答案（见 [OEIS A001109](https://oeis.org/A001109)）为

$$
1,6,35,204
$$

> 上述数据用程序枚举 $[1,1000]$ 内的 $n$，调用上面的代码，即可得到。

```py [sol-Python3]
ANS = {1: 1, 8: 6, 49: 35, 288: 204}

class Solution:
    def pivotInteger(self, n: int) -> int:
        return ANS.get(n, -1)
```

```java [sol-Java]
class Solution {
    private static final Map<Integer, Integer> m = Map.of(1, 1, 8, 6, 49, 35, 288, 204);

    public int pivotInteger(int n) {
        return m.getOrDefault(n, -1);
    }
}
```

```cpp [sol-C++]
class Solution {
    const unordered_map<int, int> m{{1, 1}, {8, 6}, {49, 35}, {288, 204}};
public:
    int pivotInteger(int n) {
        auto it = m.find(n);
        return it != m.end() ? it->second : -1;
    }
};
```

```go [sol-Go]
var m = map[int]int{1: 1, 8: 6, 49: 35, 288: 204}

func pivotInteger(n int) int {
	if ans, ok := m[n]; ok {
		return ans
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
