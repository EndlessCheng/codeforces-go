要计算合法方案数（每个小朋友分到的糖果都不超过 $\textit{limit}$），可以先计算所有方案数（没有 $\textit{limit}$ 限制），再减去不合法的方案数（至少一个小朋友分到的糖果超过 $\textit{limit}$）。

#### 所有方案数

相当于把 $n$ 个无区别的小球放入 $3$ 个有区别的盒子，允许空盒的方案数。

**隔板法**：假设 $n$ 个球和 $2$ 个隔板放到 $n+2$ 个位置，第一个隔板前的球放入第一个盒子，第一个隔板和第二个隔板之间的球放入第二个盒子，第二个隔板后的球放入第三个盒子。那么从 $n+2$ 个位置中选 $2$ 个位置放隔板，有 $C(n+2, 2)$ 种放法。

注意隔板可以放在最左边或最右边，也可以连续放，对应着空盒的情况。例如第一个隔板放在最左边，意味着第一个盒子是空的；又例如第一个隔板和第二个隔板相邻，意味着第二个盒子是空的。

#### 至少一个小朋友分到的糖果超过 limit

设三个小朋友分别叫 $A,B,C$。

只关注 $A$。如果 $A$ 分到的糖果超过 $\textit{limit}$，那么先分给他 $\textit{limit}+1$ 颗糖果，问题变成剩下 $n-(\textit{limit}+1)$ 颗糖果分给三个小朋友的方案数，即 $C(n-(\textit{limit}+1)+2, 2)$。注意 $B$ 和 $C$ 分到的糖果是否超过 $\textit{limit}$ 我们是不关注的。

只关注 $B$ 的情况和只关注 $C$ 的情况同上，均为 $C(n-(\textit{limit}+1)+2, 2)$。

直接加起来，就是 $3\cdot C(n-(\textit{limit}+1)+2, 2)$，但这样就重复统计了「至少两个小朋友分到的糖果超过 $\textit{limit}$」的情况，要减去。

> 注：三个小朋友分到的糖果均超过 $\textit{limit}$ 的情况，已经包含在**至少**两个小朋友分到的糖果超过 $\textit{limit}$ 的情况中了。

#### 至少两个小朋友分到的糖果超过 limit

只关注 $A$ 和 $B$。如果他们俩分到的糖果超过 $\textit{limit}$，那么先分给他俩 $2\cdot (\textit{limit}+1)$ 颗糖果，问题变成剩下 $n-2\cdot (\textit{limit}+1)$ 颗糖果分给三个小朋友的方案数，即 $C(n-2\cdot(\textit{limit}+1)+2, 2)$。注意 $C$ 分到的糖果是否超过 $\textit{limit}$ 我们是不关注的。

只关注 $A,C$ 的情况和只关注 $B,C$ 的情况同上，均为 $C(n-2\cdot(\textit{limit}+1)+2, 2)$。

直接加起来，就是 $3\cdot C(n-2\cdot(\textit{limit}+1)+2, 2)$，但这样就重复统计了「三个小朋友分到的糖果均超过 $\textit{limit}$」的情况，要减去。

#### 三个小朋友分到的糖果超过 limit

先分给三人一共 $3\cdot (\textit{limit}+1)$ 颗糖果，问题变成剩下 $n-3\cdot (\textit{limit}+1)$ 颗糖果分给三个小朋友的方案数，即 $C(n-3\cdot(\textit{limit}+1)+2, 2)$。

#### 总结

不合法的方案数为「至少一个」减去「至少两个」加上「三个」，这就是**容斥原理**。

最后用所有方案数减去不合法的方案数，整理得到答案：

$$
C(n+2, 2) - 3\cdot C(n-\textit{limit}+1, 2) + 3\cdot C(n-2\cdot\textit{limit}, 2) - C(n-3\cdot \textit{limit}-1, 2)
$$

请看 [视频讲解](https://www.bilibili.com/video/BV1Ww411T7JP/)，欢迎点赞关注~

```py [sol-Python3]
def c2(n: int) -> int:
    return n * (n - 1) // 2 if n > 1 else 0

class Solution:
    def distributeCandies(self, n: int, limit: int) -> int:
        return c2(n + 2) - 3 * c2(n - limit + 1) + 3 * c2(n - 2 * limit) - c2(n - 3 * limit - 1)
```

```java [sol-Java]
class Solution {
    public int distributeCandies(int n, int limit) {
        return c2(n + 2) - 3 * c2(n - limit + 1) + 3 * c2(n - 2 * limit) - c2(n - 3 * limit - 1);
    }

    private int c2(int n) {
        return n > 1 ? n * (n - 1) / 2 : 0;
    }
}
```

```cpp [sol-C++]
class Solution {
    int c2(int n) {
        return n > 1 ? n * (n - 1) / 2 : 0;
    }

public:
    int distributeCandies(int n, int limit) {
        return c2(n + 2) - 3 * c2(n - limit + 1) + 3 * c2(n - 2 * limit) - c2(n - 3 * limit - 1);
    }
};
```

```c [sol-C]
int c2(int n) {
    return n > 1 ? n * (n - 1) / 2 : 0;
}

int distributeCandies(int n, int limit) {
    return c2(n + 2) - 3 * c2(n - limit + 1) + 3 * c2(n - 2 * limit) - c2(n - 3 * limit - 1);
}
```

```go [sol-Go]
func c2(n int) int {
	if n < 2 {
		return 0
	}
	return n * (n - 1) / 2
}

func distributeCandies(n int, limit int) int {
	return c2(n+2) - 3*c2(n-limit+1) + 3*c2(n-2*limit) - c2(n-3*limit-1)
}
```

```js [sol-JavaScript]
function c2(n) {
    return n > 1 ? n * (n - 1) / 2 : 0;
}
    
var distributeCandies = function(n, limit) {
    return c2(n + 2) - 3 * c2(n - limit + 1) + 3 * c2(n - 2 * limit) - c2(n - 3 * limit - 1);
};
```

```rust [sol-Rust]
impl Solution {
    pub fn distribute_candies(n: i32, limit: i32) -> i32 {
        let c2 = |n| if n > 1 { n * (n - 1) / 2 } else { 0 };
        c2(n + 2) - 3 * c2(n - limit + 1) + 3 * c2(n - 2 * limit) - c2(n - 3 * limit - 1)
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

见 [数学题单](https://leetcode.cn/circle/discuss/IYT3ss/) 中的「容斥原理」。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
