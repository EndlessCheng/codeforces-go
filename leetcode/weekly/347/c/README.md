## 思路

注意到，反转操作只会改变 $s[i-1]$ 与 $s[i]$ 是否相等，不会改变其他任何相邻字符是否相等（相等的都反转还是相等，不相等的都反转还是不相等），所以每对相邻字符是**互相独立**的，我们只需要分别计算这些不相等的相邻字符的最小成本，累加即为答案。

## 算法

遍历 $s$，如果 $s[i-1]\ne s[i]$，那么必须反转，不然这两个字符无法相等。

- 要么反转前缀 $s[0]$ 到 $s[i-1]$，成本为 $i$。
- 要么反转后缀 $s[i]$ 到 $s[n-1]$，成本为 $n-i$。

二者取最小值，即 $\min(i,n-i)$，加入答案。

```py [sol-Python3]
class Solution:
    def minimumCost(self, s: str) -> int:
        n = len(s)
        ans = 0
        for i in range(1, n):
            if s[i - 1] != s[i]:
                ans += min(i, n - i)
        return ans
```

```py [sol-Python3 一行]
class Solution:
    def minimumCost(self, s: str) -> int:
        return sum(min(i, len(s) - i) for i, (x, y) in enumerate(pairwise(s), 1) if x != y)
```

```java [sol-Java]
class Solution {
    public long minimumCost(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        long ans = 0;
        for (int i = 1; i < n; i++) {
            if (s[i - 1] != s[i]) {
                ans += Math.min(i, n - i);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minimumCost(string s) {
        int n = s.size();
        long long ans = 0;
        for (int i = 1; i < n; i++) {
            if (s[i - 1] != s[i]) {
                ans += min(i, n - i);
            }
        }
        return ans;
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))

long long minimumCost(char* s) {
    int n = strlen(s);
    long long ans = 0;
    for (int i = 1; i < n; i++) {
        if (s[i - 1] != s[i]) {
            ans += MIN(i, n - i);
        }
    }
    return ans;
}
```

```go [sol-Go]
func minimumCost(s string) (ans int64) {
    n := len(s)
    for i := 1; i < n; i++ {
        if s[i-1] != s[i] {
            ans += int64(min(i, n-i))
        }
    }
    return
}
```

```js [sol-JavaScript]
var minimumCost = function(s) {
    const n = s.length;
    let ans = 0;
    for (let i = 1; i < n; i++) {
        if (s[i - 1] !== s[i]) {
            ans += Math.min(i, n - i);
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn minimum_cost(s: String) -> i64 {
        let s = s.as_bytes();
        let n = s.len();
        let mut ans = 0;
        for i in 1..n {
            if s[i - 1] != s[i] {
                ans += i.min(n - i) as i64;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

更多相似题目，见下面贪心与思维题单中的「**§5.2 脑筋急转弯**」。

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
10. 【本题相关】[贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
