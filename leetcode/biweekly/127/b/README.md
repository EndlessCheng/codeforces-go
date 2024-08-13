**题意**：把 $0$ 视作 $-1$，找一个最短前缀，其元素和大于剩余元素和。

设 $\textit{possible}$ 的元素和为 $s$（把 $0$ 视作 $-1$）。

枚举 $x=\textit{possible}[i]$，同时计算前缀和 $\textit{pre}$，那么剩余元素和为 $s - \textit{pre}$

如果

$$
\textit{pre} > s - \textit{pre}
$$

即

$$
\textit{pre}\cdot 2 > s
$$

就返回 $i+1$，即前缀长度。

代码实现时，计算 $\textit{pre}$ 可以把 $1$ 视作 $2$，把 $0$ 视作 $-2$，这样无需计算乘 $2$。

附：[视频讲解](https://www.bilibili.com/video/BV19t421g7Pd/)，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def minimumLevels(self, possible: List[int]) -> int:
        # s = cnt1 - cnt0 = cnt1 - (n - cnt1) = cnt1 * 2 - n
        s = sum(possible) * 2 - len(possible)
        pre = 0
        for i, x in enumerate(possible[:-1]):
            pre += 2 if x else -2
            if pre > s:
                return i + 1
        return -1
```

```java [sol-Java]
class Solution {
    public int minimumLevels(int[] possible) {
        int n = possible.length;
        // s = cnt1 - cnt0 = cnt1 - (n - cnt1) = cnt1 * 2 - n
        int s = 0;
        for (int x : possible) {
            s += x;
        }
        s = s * 2 - n;
        int pre = 0;
        for (int i = 0; i < n - 1; i++) {
            pre += possible[i] == 1 ? 2 : -2;
            if (pre > s) {
                return i + 1;
            }
        }
        return -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumLevels(vector<int>& possible) {
        int n = possible.size();
        // s = cnt1 - cnt0 = cnt1 - (n - cnt1) = cnt1 * 2 - n
        int s = reduce(possible.begin(), possible.end()) * 2 - n;
        int pre = 0;
        for (int i = 0; i < n - 1; i++) {
            pre += possible[i] ? 2 : -2;
            if (pre > s) {
                return i + 1;
            }
        }
        return -1;
    }
};
```

```c [sol-C]
int minimumLevels(int* possible, int possibleSize) {
    // s = cnt1 - cnt0 = cnt1 - (n - cnt1) = cnt1 * 2 - n
    int s = 0;
    for (int i = 0; i < possibleSize; i++) {
        s += possible[i];
    }
    s = s * 2 - possibleSize;
    int pre = 0;
    for (int i = 0; i < possibleSize - 1; i++) {
        pre += possible[i] ? 2 : -2;
        if (pre > s) {
            return i + 1;
        }
    }
    return -1;
}
```

```go [sol-Go]
func minimumLevels(possible []int) int {
    n := len(possible)
    // s = cnt1 - cnt0 = cnt1 - (n - cnt1) = cnt1 * 2 - n
    s := 0
    for _, x := range possible {
        s += x
    }
    s = s*2 - n
    pre := 0
    for i, x := range possible[:n-1] {
        pre += x*4 - 2
        if pre > s {
            return i + 1
        }
    }
    return -1
}
```

```js [sol-JavaScript]
var minimumLevels = function(possible) {
    const n = possible.length;
    // s = cnt1 - cnt0 = cnt1 - (n - cnt1) = cnt1 * 2 - n
    const s = _.sum(possible) * 2 - n;
    let pre = 0;
    for (let i = 0; i < n - 1; i++) {
        pre += possible[i] ? 2 : -2;
        if (pre > s) {
            return i + 1;
        }
    }
    return -1;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn minimum_levels(possible: Vec<i32>) -> i32 {
        // s = cnt1 - cnt0 = cnt1 - (n - cnt1) = cnt1 * 2 - n
        let s = possible.iter().sum::<i32>() * 2 - possible.len() as i32;
        let mut pre = 0;
        for i in 0..possible.len() - 1 {
            pre += if possible[i] == 1 { 2 } else { -2 };
            if pre > s {
                return (i + 1) as _;
            }
        }
        -1
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{possible}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。Python 忽略切片空间。

更多相似题目，见下面的 DP 题单中的「**前后缀分解**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
