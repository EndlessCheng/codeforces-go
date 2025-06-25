设 $k$ 的二进制长度为 $m$。比如 $k=4=100_{(2)}$，二进制长度 $m=3$。

**核心思路**：

1. 任意长为 $m-1$ 的子序列，数值一定小于 $k$。比如 $00$，$01$，$10$，$11$ 都小于 $100$。
2. 前导零不影响数值。**子序列越靠右，我们能添加的前导零就越多，子序列就越长**。贪心地，选 $s$ 的**后缀**作为子序列。（为什么？理由见下面的问答）
3. 如果 $s$ 的长为 $m$ 的后缀 $\le k$，那么就选长为 $m$ 的后缀作为子序列；否则选长为 $m-1$ 的后缀（一定小于 $k$）作为子序列。在此基础上，加上长为 $n-m$ 的前缀中的 $0$ 的个数，即为答案。

在示例 1 中，$s=1001010$，$k=101_{(2)}$，$s$ 的长为 $m=3$ 的后缀 $010\le k$。如果在 $010$ 的前面添加 $1$，就超过 $k$ 了。此外，由于我们选的是 $s$ 的后缀，可以往 $010$ 前面添加尽量多的前导零，从而增加子序列的长度，但不会增大子序列的数值。最终得到 $00 + 010 = 00010$，长为 $5$。

又比如 $s=1001110$，$k=101_{(2)}$，长为 $m$ 的后缀 $110>k$，但长为 $m-1$ 的后缀 $10$ 一定小于 $k$。在这个后缀的前面添加两个前导零，最终得到 $00 + 10 = 0010$，长为 $4$。

**问**：在上例中，虽然 $110>k$，但我可以选其他长为 $m$ 的子序列呀？比如 $100$，$101$，长为 $m$ 且 $\le k$。为什么不考虑这样的子序列呢？

**答**：这是因为，要想得到比后缀 $110$ 更小的长为 $m$ 的子序列，比如 $100$，$101$，这样的子序列必然包含「本应添加到后缀 $10$ 前的前导零」，这些长为 $m$ 的子序列，相比长为 $m-1$ 的后缀，长度仅仅增加了 $1$，但我们只需要在长为 $m-1$ 的后缀的基础上，再添加一个前导零，也能得到长为 $m$ 的子序列 $010$，且这个子序列的第一个字符是「尽量靠右」的，可以添加更多的前导零。所以这些长为 $m$ 的子序列，并不会比「一个 $0$ + 长为 $m-1$ 的后缀」这种构造方式更优。

```py [sol-Python3]
class Solution:
    def longestSubsequence(self, s: str, k: int) -> int:
        n, m = len(s), k.bit_length()
        if n < m:  # int(s, 2) < k
            return n  # 全选
        ans = m if int(s[-m:], 2) <= k else m - 1  # 后缀长度
        return ans + s[:-m].count('0')  # 添加前导零
```

```java [sol-Java]
class Solution {
    public int longestSubsequence(String s, int k) {
        int n = s.length();
        int m = 32 - Integer.numberOfLeadingZeros(k); // k 的二进制长度
        if (n < m) {
            return n; // 全选
        }

        int sufVal = Integer.parseInt(s.substring(n - m), 2);
        int ans = sufVal <= k ? m : m - 1; // 后缀长度

        for (int i = 0; i < n - m; i++) {
            ans += '1' - s.charAt(i); // 添加前导零
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestSubsequence(string s, int k) {
        int n = s.length();
        int m = bit_width((uint32_t) k);
        if (n < m) {
            return n; // 全选
        }
        int suf_val = stoi(s.substr(n - m), nullptr, 2);
        int ans = suf_val <= k ? m : m - 1; // 后缀长度
        return ans + count(s.begin(), s.end() - m, '0'); // 添加前导零
    }
};
```

```c [sol-C]
int longestSubsequence(char* s, int k) {
    int n = strlen(s);
    int m = 32 - __builtin_clz(k); // k 的二进制长度
    if (n < m) {
        return n; // 全选
    }

    int suf_val = strtol(s + n - m, NULL, 2);
    int ans = suf_val <= k ? m : m - 1; // 后缀长度

    for (int i = 0; i < n - m; i++) {
        ans += '1' - s[i]; // 添加前导零
    }
    return ans;
}
```

```go [sol-Go]
func longestSubsequence(s string, k int) int {
	n, m := len(s), bits.Len(uint(k))
	if n < m {
		return n // 全选
	}
	ans := m // 后缀长度
	sufVal, _ := strconv.ParseInt(s[n-m:], 2, 0)
	if int(sufVal) > k {
		ans--
	}
	return ans + strings.Count(s[:n-m], "0") // 添加前导零
}
```

```js [sol-JavaScript]
var longestSubsequence = function(s, k) {
    const n = s.length;
    const m = 32 - Math.clz32(k); // k 的二进制长度
    if (n < m) {
        return n; // 全选
    }

    const sufVal = parseInt(s.slice(n - m), 2);
    let ans = sufVal <= k ? m : m - 1; // 后缀长度

    for (let i = 0; i < n - m; i++) {
        if (s[i] === '0') {
            ans++; // 添加前导零
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn longest_subsequence(s: String, k: i32) -> i32 {
        let n = s.len();
        let m = (32 - k.leading_zeros()) as usize; // k 的二进制长度
        if n < m {
            return n as _; // 全选
        }

        let suf_val = i32::from_str_radix(&s[n - m..], 2).unwrap();
        let ans = if suf_val <= k { m } else { m - 1 }; // 后缀长度

        (ans + s[..n - m].bytes().filter(|&c| c == b'0').count()) as _ // 添加前导零
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。这里是严格意义上的一次遍历，$s$ 的每个字符都会被恰好遍历一次。（C 语言要计算字符串长度，遍历两次）
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(m)$ 或 $\mathcal{O}(1)$，取决于实现。手动遍历计算可以做到 $\mathcal{O}(1)$ 空间。

## 思考题

把**子序列**改成**子串**，要怎么做？你能做到 $\mathcal{O}(n)$ 时间复杂度吗？

欢迎在评论区发表你的思路/代码。

## 专题训练

见下面贪心与思维题单的「**§5.2 脑筋急转弯**」。

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
