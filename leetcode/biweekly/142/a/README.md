举例说明：

- 如果所有相邻字母都互不相同，那么 Alice 不可能犯错，所以方案数只有 $1$ 种。
- 如果有 $1$ 对相邻字母相同，那么 Alice 可以在这里犯错一次，例如 $\texttt{abb}$，一开始想要输入的可能是 $\texttt{abb}$，也可能是 $\texttt{ab}$，其中 $\texttt{b}$ 多按了一次，所以方案数有 $2$ 种。
- 如果有 $2$ 对相邻字母相同，那么一开始想要输入的字符串会再多一种：
   - 例如 $\texttt{abbb}$，一开始想要输入的可能是 $\texttt{abbb}$，也可能是 $\texttt{abb}$（$\texttt{b}$ 多按了一次），也可能是 $\texttt{ab}$（$\texttt{b}$ 多按了两次），所以方案数有 $3$ 种。
   - 例如 $\texttt{aabb}$，一开始想要输入的可能是 $\texttt{aabb}$，也可能是 $\texttt{abb}$（$\texttt{a}$ 多按了一次），也可能是 $\texttt{aab}$（$\texttt{b}$ 多按了一次），所以方案数有 $3$ 种。**注意**：一开始想要输入的不可能是 $\texttt{ab}$，因为题目说 Alice 至多犯错一次，也就是重复输入**一个**字母**多次**。
- 依此类推，每有一对相邻相同字母，Alice 就会多一种犯错的方案。所以方案数等于相邻相同字母对的个数加一，其中加一是不犯错的情况。

```py [sol-Python3]
class Solution:
    def possibleStringCount(self, word: str) -> int:
        ans = 1
        for x, y in pairwise(word):
            if x == y:
                ans += 1
        return ans
```

```py [sol-Python3 一行]
class Solution:
    def possibleStringCount(self, word: str) -> int:
        return 1 + sum(x == y for x, y in pairwise(word))
```

```java [sol-Java]
class Solution {
    public int possibleStringCount(String word) {
        int ans = 1;
        for (int i = 1; i < word.length(); i++) {
            if (word.charAt(i - 1) == word.charAt(i)) {
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int possibleStringCount(string word) {
        int ans = 1;
        for (int i = 1; i < word.length(); i++) {
            if (word[i - 1] == word[i]) {
                ans++;
            }
        }
        return ans;
    }
};
```

```c [sol-C]
int possibleStringCount(char* word) {
    int ans = 1;
    for (int i = 1; word[i]; i++) {
        if (word[i - 1] == word[i]) {
            ans++;
        }
    }
    return ans;
}
```

```go [sol-Go]
func possibleStringCount(word string) int {
	ans := 1
	for i := 1; i < len(word); i++ {
		if word[i-1] == word[i] {
			ans++
		}
	}
	return ans
}
```

```js [sol-JavaScript]
var possibleStringCount = function(word) {
    let ans = 1;
    for (let i = 1; i < word.length; i++) {
        if (word[i - 1] === word[i]) {
            ans++;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn possible_string_count(word: String) -> i32 {
        let s = word.into_bytes();
        let mut ans = 1;
        for i in 1..s.len() {
            if s[i - 1] == s[i] {
                ans += 1;
            }
        }
        ans
    }
}
```

```rust [sol-Rust 一行]
impl Solution {
    pub fn possible_string_count(word: String) -> i32 {
        1 + word.into_bytes().windows(2).filter(|w| w[0] == w[1]).count() as i32
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{word}$ 的长度。
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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
