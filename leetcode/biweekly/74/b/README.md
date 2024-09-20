首先计算插入之前的 $\textit{pattern}$ 子序列的个数，然后计算因为插入字母额外增加的 $\textit{pattern}$ 子序列的个数。

设 $x=\textit{pattern}[0],\ y=\textit{pattern}[1]$。

遍历 $\textit{text}$ 统计答案：遇到 $y$ 时，如果左边出现了 $3$ 个 $x$，那么就意味着我们找到了 $3$ 个 $\textit{pattern}$ 子序列，把 $3$ 加入答案。一般地，在遍历 $\textit{text}$ 的同时，维护 $x$ 的出现次数 $\textit{cntX}$。遇到 $y$ 时，把 $\textit{cntX}$ 加入答案。

然后考虑插入字母。

根据题意，$x$ 插入的位置越靠左，$\textit{pattern}$ 子序列的个数越多；$y$ 插入的位置越靠右，$\textit{pattern}$ 子序列的个数越多。那么 $x$ 应插在 $\textit{text}$ 最左侧，$y$ 应插在 $\textit{text}$ 最右侧。

分类讨论：

- 把 $x$ 插在 $\textit{text}$ 最左侧：答案额外增加 $\textit{cntY}$，其中 $\textit{cntY}$ 是 $y$ 在 $\textit{text}$ 中的出现次数。
- 把 $y$ 插在 $\textit{text}$ 最右侧：答案额外增加 $\textit{cntX}$，其中 $\textit{cntX}$ 是 $x$ 在 $\textit{text}$ 中的出现次数。

⚠**注意**：代码没有特判 $x=y$ 的情况，要先更新答案，再更新 $\textit{cntX}$，这可以保证更新答案时 $\textit{cntX}$ 表示的是当前字母**左边**的 $\textit{x}$ 的出现次数，$\textit{cntX}$ 尚未计入当前字母。

```py [sol-Python3]
class Solution:
    def maximumSubsequenceCount(self, text: str, pattern: str) -> int:
        x, y = pattern
        ans = cnt_x = cnt_y = 0
        for c in text:
            if c == y:
                ans += cnt_x
                cnt_y += 1
            if c == x:
                cnt_x += 1
        return ans + max(cnt_x, cnt_y)
```

```java [sol-Java]
class Solution {
    public long maximumSubsequenceCount(String text, String pattern) {
        char x = pattern.charAt(0);
        char y = pattern.charAt(1);
        long ans = 0;
        int cntX = 0;
        int cntY = 0;
        for (char c : text.toCharArray()) {
            if (c == y) {
                ans += cntX;
                cntY++;
            }
            if (c == x) {
                cntX++;
            }
        }
        return ans + Math.max(cntX, cntY);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumSubsequenceCount(string text, string pattern) {
        char x = pattern[0], y = pattern[1];
        long long ans = 0;
        int cnt_x = 0, cnt_y = 0;
        for (char c : text) {
            if (c == y) {
                ans += cnt_x;
                cnt_y++;
            }
            if (c == x) {
                cnt_x++;
            }
        }
        return ans + max(cnt_x, cnt_y);
    }
};
```

```c [sol-C]
#define MAX(a, b) ((b) > (a) ? (b) : (a))

long long maximumSubsequenceCount(char* text, char* pattern) {
    char x = pattern[0], y = pattern[1];
    long long ans = 0;
    int cnt_x = 0, cnt_y = 0;
    for (int i = 0; text[i]; i++) {
        if (text[i] == y) {
            ans += cnt_x;
            cnt_y++;
        }
        if (text[i] == x) {
            cnt_x++;
        }
    }
    return ans + MAX(cnt_x, cnt_y);
}
```

```go [sol-Go]
func maximumSubsequenceCount(text, pattern string) (ans int64) {
    x, y := pattern[0], pattern[1]
    cntX, cntY := 0, 0
    for i := range text {
        c := text[i]
        if c == y {
            ans += int64(cntX)
            cntY++
        }
        if c == x {
            cntX++
        }
    }
    return ans + int64(max(cntX, cntY))
}
```

```js [sol-JavaScript]
var maximumSubsequenceCount = function(text, pattern) {
    const [x, y] = pattern;
    let ans = 0, cntX = 0, cntY = 0;
    for (const c of text) {
        if (c === y) {
            ans += cntX;
            cntY++;
        }
        if (c === x) {
            cntX++;
        }
    }
    return ans + Math.max(cntX, cntY);
};
```

```rust [sol-Rust]
impl Solution {
    pub fn maximum_subsequence_count(text: String, pattern: String) -> i64 {
        let pattern = pattern.as_bytes();
        let x = pattern[0];
        let y = pattern[1];
        let mut ans = 0i64;
        let mut cnt_x = 0;
        let mut cnt_y = 0;
        for c in text.bytes() {
            if c == y {
                ans += cnt_x as i64;
                cnt_y += 1;
            }
            if c == x {
                cnt_x += 1;
            }
        }
        ans + cnt_x.max(cnt_y) as i64
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{text}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

如果 $\textit{pattern}$ 的长度是 $3$ 呢？是 $m$ 呢？

欢迎在评论区分享你的思路/代码。

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
