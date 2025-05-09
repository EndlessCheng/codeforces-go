注意输入的字符串长度都是 $2$。

## 分类讨论

- 字符串的两个字母**不同**：假设有 $2$ 个 $\texttt{ab}$ 和 $3$ 个 $\texttt{ba}$，我们可以在回文串的左半边中放 $2$ 个 $\texttt{ab}$，右半边中的对称位置放 $2$ 个 $\texttt{ba}$。剩下 $1$ 个 $\texttt{ba}$ 没法用，舍去。
- 字符串的两个字母**相同**：
    - 偶数个：假设有 $4$ 个 $\texttt{aa}$，我们可以在回文串的左半边中放 $2$ 个 $\texttt{aa}$，右半边中的对称位置放 $2$ 个 $\texttt{aa}$。
    - 奇数个：假设有 $5$ 个 $\texttt{aa}$，我们可以在回文串的左半边中放 $2$ 个 $\texttt{aa}$，右半边中的对称位置放 $2$ 个 $\texttt{aa}$。多出的 $1$ 个 $\texttt{aa}$ 放在回文串的**正中间**。

## 算法

首先遍历 $\textit{words}$，统计每个字符串的出现次数，记在 $\textit{cnt}$ 中。为方便下面枚举，可以用一个 $26 \times 26$ 的二维数组统计。

然后计算总个数，也就是回文串由多少个长为 $2$ 的字符串组成。最后把总个数乘以 $2$，就是回文串的长度。

对于两个字母不同的情况，枚举 $0\le i < j < 26$。设 $c$ 为 $\texttt{ij}$ 和 $\texttt{ji}$ 的出现次数的较小者，我们可以在回文串的左半边中放 $c$ 个 $\texttt{ij}$，右半边中的对称位置放 $c$ 个 $\texttt{ji}$，也就是把总个数增加

$$
\min(\textit{cnt}[i][j], \textit{cnt}[j][i])\cdot 2
$$

对于两个字母相同的情况，枚举 $0\le i < 26$。设 $c=\textit{cnt}[i][i]$，分类讨论：

- 如果 $c$ 是偶数，那么把总个数增加 $c$。
- 如果 $c$ 是奇数，那么把总个数增加 $c-1$。
- 这两种情况可以统一为：把总个数增加 $c - c\bmod 2$。

最后，如果存在奇数 $\textit{cnt}[i][i]$，那么可以把一个 $\texttt{ii}$ 放在回文串的正中间，总个数额外加一。

```py [sol-Python3]
class Solution:
    def longestPalindrome(self, words: List[str]) -> int:
        cnt = [[0] * 26 for _ in range(26)]  # 更快的写法见【Python3 写法二】
        for w in words:
            cnt[ord(w[0]) - ord('a')][ord(w[1]) - ord('a')] += 1

        ans = odd = 0
        for i in range(26):
            c = cnt[i][i]
            ans += c - c % 2  # 保证结果是偶数，也可以写成 c & ~1
            odd |= c % 2  # 存在出现奇数次的 cnt[i][i]
            for j in range(i + 1, 26):
                ans += min(cnt[i][j], cnt[j][i]) * 2
        return (ans + odd) * 2  # 上面统计的是字符串个数，乘以 2 就是长度
```

```py [sol-Python3 写法二]
class Solution:
    def longestPalindrome(self, words: List[str]) -> int:
        cnt = Counter(words)

        ans = odd = 0
        for w, c in cnt.items():
            if w[0] == w[1]:
                ans += c - c % 2
                odd |= c % 2
            elif w[0] < w[1]:
                ans += min(c, cnt[w[::-1]]) * 2
        return (ans + odd) * 2
```

```java [sol-Java]
class Solution {
    public int longestPalindrome(String[] words) {
        int[][] cnt = new int[26][26];
        for (String w : words) {
            cnt[w.charAt(0) - 'a'][w.charAt(1) - 'a']++;
        }

        int ans = 0;
        int odd = 0;
        for (int i = 0; i < 26; i++) {
            int c = cnt[i][i];
            ans += c - c % 2; // 保证结果是偶数，也可以写成 c & ~1
            odd |= c % 2; // 存在出现奇数次的 cnt[i][i]
            for (int j = i + 1; j < 26; j++) {
                ans += Math.min(cnt[i][j], cnt[j][i]) * 2;
            }
        }
        return (ans + odd) * 2; // 上面统计的是字符串个数，乘以 2 就是长度
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestPalindrome(vector<string>& words) {
        int cnt[26][26]{};
        for (auto& w : words) {
            cnt[w[0] - 'a'][w[1] - 'a']++;
        }

        int ans = 0, odd = 0;
        for (int i = 0; i < 26; i++) {
            int c = cnt[i][i];
            ans += c - c % 2; // 保证结果是偶数，也可以写成 c & ~1
            odd |= c % 2; // 存在出现奇数次的 cnt[i][i]
            for (int j = i + 1; j < 26; j++) {
                ans += min(cnt[i][j], cnt[j][i]) * 2;
            }
        }
        return (ans + odd) * 2; // 上面统计的是字符串个数，乘以 2 就是长度
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))

int longestPalindrome(char** words, int wordsSize) {
    int cnt[26][26] = {};
    for (int i = 0; i < wordsSize; i++) {
        cnt[words[i][0] - 'a'][words[i][1] - 'a']++;
    }

    int ans = 0, odd = 0;
    for (int i = 0; i < 26; i++) {
        int c = cnt[i][i];
        ans += c - c % 2; // 保证结果是偶数，也可以写成 c & ~1
        odd |= c % 2; // 存在出现奇数次的 cnt[i][i]
        for (int j = i + 1; j < 26; j++) {
            ans += MIN(cnt[i][j], cnt[j][i]) * 2;
        }
    }
    return (ans + odd) * 2; // 上面统计的是字符串个数，乘以 2 就是长度
}
```

```go [sol-Go]
func longestPalindrome(words []string) (ans int) {
	cnt := [26][26]int{}
	for _, w := range words {
		cnt[w[0]-'a'][w[1]-'a']++
	}

	odd := 0 // 是否存在出现奇数次的 cnt[i][i]
	for i := range cnt {
		c := cnt[i][i]
		ans += c - c%2 // 保证结果是偶数，也可以写成 c &^ 1
		odd |= c % 2   // 存在出现奇数次的 cnt[i][i]
		for j := i + 1; j < 26; j++ {
			ans += min(cnt[i][j], cnt[j][i]) * 2
		}
	}
	return (ans + odd) * 2 // 上面统计的是字符串个数，乘以 2 就是长度 
}
```

```js [sol-JavaScript]
var longestPalindrome = function(words) {
    const cnt = Array.from({ length: 26 }, () => Array(26).fill(0));
    const ordA = 'a'.charCodeAt(0);
    for (const w of words) {
        cnt[w.charCodeAt(0) - ordA][w.charCodeAt(1) - ordA]++;
    }

    let ans = 0, odd = 0;
    for (let i = 0; i < 26; i++) {
        const c = cnt[i][i];
        ans += c - c % 2; // 保证结果是偶数，也可以写成 c & ~1
        odd |= c % 2; // 存在出现奇数次的 cnt[i][i]
        for (let j = i + 1; j < 26; j++) {
            ans += Math.min(cnt[i][j], cnt[j][i]) * 2;
        }
    }
    return (ans + odd) * 2; // 上面统计的是字符串个数，乘以 2 就是长度
};
```

```rust [sol-Rust]
impl Solution {
    pub fn longest_palindrome(words: Vec<String>) -> i32 {
        let mut cnt = [[0; 26]; 26];
        for w in words {
            let w = w.into_bytes();
            cnt[(w[0] - b'a') as usize][(w[1] - b'a') as usize] += 1;
        }

        let mut ans = 0;
        let mut odd = 0;
        for i in 0..26 {
            let c = cnt[i][i];
            ans += c - c % 2; // 保证结果是偶数，也可以写成 c & !1
            odd |= c % 2; // 存在出现奇数次的 cnt[i][i]
            for j in i + 1..26 {
                ans += cnt[i][j].min(cnt[j][i]) * 2;
            }
        }
        (ans + odd) * 2 // 上面统计的是字符串个数，乘以 2 就是长度
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|^2)$，其中 $n$ 是 $\textit{words}$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|^2)$。

更多相似题目，见下面贪心题单的「**§3.2 回文串贪心**」。

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
