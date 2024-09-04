从左到右遍历 $s$，同时维护连续递增长度 $\textit{cnt}$：

- 如果 $s[i-1]+1 = s[i]$，把 $\textit{cnt}$ 加一，然后用 $\textit{cnt}$ 更新答案的最大值。
- 如果 $s[i-1]+1\ne s[i]$，把 $\textit{cnt}$ 重置为 $1$。

$\textit{cnt}$ 和答案的初始值均为 $1$。

```py [sol-Python3]
class Solution:
    def longestContinuousSubstring(self, s: str) -> int:
        ans = cnt = 1
        for x, y in pairwise(map(ord, s)):
            cnt = cnt + 1 if x + 1 == y else 1
            ans = max(ans, cnt)
        return ans
```

```java [sol-Java]
class Solution {
    public int longestContinuousSubstring(String S) {
        char[] s = S.toCharArray();
        int ans = 1;
        int cnt = 1;
        for (int i = 1; i < s.length; i++) {
            if (s[i - 1] + 1 == s[i]) {
                ans = Math.max(ans, ++cnt);
            } else {
                cnt = 1;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestContinuousSubstring(string s) {
        int ans = 1, cnt = 1;
        for (int i = 1; i < s.length(); i++) {
            if (s[i - 1] + 1 == s[i]) {
                ans = max(ans, ++cnt);
            } else {
                cnt = 1;
            }
        }
        return ans;
    }
};
```

```c [sol-C]
#define MAX(a, b) ((b) > (a) ? (b) : (a))

int longestContinuousSubstring(char* s) {
    int ans = 1, cnt = 1;
    for (int i = 1; s[i]; i++) {
        if (s[i - 1] + 1 == s[i]) {
            cnt++;
            ans = MAX(ans, cnt);
        } else {
            cnt = 1;
        }
    }
    return ans;
}
```

```go [sol-Go]
func longestContinuousSubstring(s string) int {
    ans, cnt := 1, 1
    for i := 1; i < len(s); i++ {
        if s[i-1]+1 == s[i] {
            cnt++
            ans = max(ans, cnt)
        } else {
            cnt = 1
        }
    }
    return ans
}
```

```js [sol-JavaScript]
var longestContinuousSubstring = function(s) {
    let ans = 1, cnt = 1;
    for (let i = 1; i < s.length; i++) {
        if (s.charCodeAt(i - 1) + 1 === s.charCodeAt(i)) {
            ans = Math.max(ans, ++cnt);
        } else {
            cnt = 1;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn longest_continuous_substring(s: String) -> i32 {
        let mut ans = 1;
        let mut cnt = 1;
        let s = s.as_bytes();
        for i in 1..s.len() {
            if s[i - 1] + 1 == s[i] {
                cnt += 1;
                ans = ans.max(cnt);
            } else {
                cnt = 1;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
