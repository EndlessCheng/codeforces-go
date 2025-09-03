在遍历 $s$ 的过程中，用哈希表（或者数组）统计每个字母的出现次数 $\textit{cnt}$。

- 如果字母是元音，更新 $\textit{maxVowelCnt}$ 的最大值。
- 如果字母是辅音，更新 $\textit{maxConsonantCnt}$ 的最大值。

最终答案是 $\textit{maxVowelCnt} + \textit{maxConsonantCnt}$。

## 写法一

```py [sol-Python3]
class Solution:
    def maxFreqSum(self, s: str) -> int:
        cnt = [0] * 26
        max_vowel_cnt = max_consonant_cnt = 0
        for ch in s:
            idx = ord(ch) - ord('a')
            cnt[idx] += 1
            if ch in "aeiou":
                max_vowel_cnt = max(max_vowel_cnt, cnt[idx])
            else:
                max_consonant_cnt = max(max_consonant_cnt, cnt[idx])
        return max_vowel_cnt + max_consonant_cnt
```

```py [sol-Python3 写法二]
class Solution:
    def maxFreqSum(self, s: str) -> int:
        cnt = Counter(s)

        max_vowel_cnt = 0
        for ch in "aeiou":
            max_vowel_cnt = max(max_vowel_cnt, cnt[ch])
            del cnt[ch]  # 这样下面计算的一定是辅音出现次数的最大值

        max_consonant_cnt = max(cnt.values(), default=0)
        return max_vowel_cnt + max_consonant_cnt
```

```java [sol-Java]
class Solution {
    public int maxFreqSum(String s) {
        int[] cnt = new int[26];
        int maxVowelCnt = 0;
        int maxConsonantCnt = 0;
        for (char ch : s.toCharArray()) {
            cnt[ch - 'a']++;
            if (ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u') {
                maxVowelCnt = Math.max(maxVowelCnt, cnt[ch - 'a']);
            } else {
                maxConsonantCnt = Math.max(maxConsonantCnt, cnt[ch - 'a']);
            }
        }
        return maxVowelCnt + maxConsonantCnt;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxFreqSum(string s) {
        int cnt[26]{};
        int max_vowel_cnt = 0;
        int max_consonant_cnt = 0;
        for (char ch : s) {
            cnt[ch - 'a']++;
            if (ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u') {
                max_vowel_cnt = max(max_vowel_cnt, cnt[ch - 'a']);
            } else {
                max_consonant_cnt = max(max_consonant_cnt, cnt[ch - 'a']);
            }
        }
        return max_vowel_cnt + max_consonant_cnt;
    }
};
```

```c [sol-C]
#define MAX(a, b) ((b) > (a) ? (b) : (a))

int maxFreqSum(char* s) {
    int cnt[26] = {};
    int max_vowel_cnt = 0;
    int max_consonant_cnt = 0;
    for (int i = 0; s[i]; i++) {
        char ch = s[i];
        cnt[ch - 'a']++;
        if (ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u') {
            max_vowel_cnt = MAX(max_vowel_cnt, cnt[ch - 'a']);
        } else {
            max_consonant_cnt = MAX(max_consonant_cnt, cnt[ch - 'a']);
        }
    }
    return max_vowel_cnt + max_consonant_cnt;
}
```

```go [sol-Go]
func maxFreqSum(s string) int {
	cnt := [26]int{}
	maxVowelCnt := 0
	maxConsonantCnt := 0
	for _, ch := range s {
		cnt[ch-'a']++
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			maxVowelCnt = max(maxVowelCnt, cnt[ch-'a'])
		} else {
			maxConsonantCnt = max(maxConsonantCnt, cnt[ch-'a'])
		}
	}
	return maxVowelCnt + maxConsonantCnt
}
```

```js [sol-JavaScript]
var maxFreqSum = function(s) {
    const cnt = Array(26).fill(0);
    let maxVowelCnt = 0;
    let maxConsonantCnt = 0;
    for (const ch of s) {
        const idx = ch.charCodeAt(0) - 'a'.charCodeAt(0);
        cnt[idx]++;
        if (ch === 'a' || ch === 'e' || ch === 'i' || ch === 'o' || ch === 'u') {
            maxVowelCnt = Math.max(maxVowelCnt, cnt[idx]);
        } else {
            maxConsonantCnt = Math.max(maxConsonantCnt, cnt[idx]);
        }
    }
    return maxVowelCnt + maxConsonantCnt;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn max_freq_sum(s: String) -> i32 {
        let mut cnt = [0; 26];
        let mut max_vowel_cnt = 0;
        let mut max_consonant_cnt = 0;
        for ch in s.bytes() {
            let idx = (ch as u8 - b'a') as usize;
            cnt[idx] += 1;
            if ch == b'a' || ch == b'e' || ch == b'i' || ch == b'o' || ch == b'u' {
                max_vowel_cnt = max_vowel_cnt.max(cnt[idx]);
            } else {
                max_consonant_cnt = max_consonant_cnt.max(cnt[idx]);
            }
        }
        max_vowel_cnt + max_consonant_cnt
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n + |\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

## 写法二

根据 [从集合论到位运算](https://leetcode.cn/circle/discuss/CaOJ45/)，我们可以把元音集合

$$
\{\texttt{a},\texttt{e},\texttt{i},\texttt{o},\texttt{u}\}
$$

视作数字

$$
2^0 + 2^4 + 2^8 + 2^{14} + 2^{20} = 1065233
$$

即十六进制的 $\texttt{0x104111}$。

可以用位运算快速判断字母是否在元音集合中。

```py [sol-Python3]
class Solution:
    def maxFreqSum(self, s: str) -> int:
        VOWEL_MASK = 0x104111
        cnt = [0] * 26
        max_cnt = [0] * 2
        for ch in s:
            ch = ord(ch) - ord('a')
            bit = VOWEL_MASK >> ch & 1
            cnt[ch] += 1
            max_cnt[bit] = max(max_cnt[bit], cnt[ch])
        return sum(max_cnt)
```

```java [sol-Java]
class Solution {
    public int maxFreqSum(String s) {
        final int VOWEL_MASK = 0x104111;
        int[] cnt = new int[26];
        int[] maxCnt = new int[2];
        for (char ch : s.toCharArray()) {
            ch -= 'a';
            int bit = VOWEL_MASK >> ch & 1;
            cnt[ch]++;
            maxCnt[bit] = Math.max(maxCnt[bit], cnt[ch]);
        }
        return maxCnt[0] + maxCnt[1];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxFreqSum(string s) {
        const int VOWEL_MASK = 0x104111;
        int cnt[26]{};
        int max_cnt[2]{};
        for (char ch : s) {
            ch -= 'a';
            int bit = VOWEL_MASK >> ch & 1;
            cnt[ch]++;
            max_cnt[bit] = max(max_cnt[bit], cnt[ch]);
        }
        return max_cnt[0] + max_cnt[1];
    }
};
```

```c [sol-C]
#define VOWEL_MASK 0x104111
#define MAX(a, b) ((b) > (a) ? (b) : (a))

int maxFreqSum(char* s) {
    int cnt[26] = {};
    int max_cnt[2] = {};
    for (int i = 0; s[i]; i++) {
        int ch = s[i] - 'a';
        int bit = VOWEL_MASK >> ch & 1;
        cnt[ch]++;
        max_cnt[bit] = MAX(max_cnt[bit], cnt[ch]);
    }
    return max_cnt[0] + max_cnt[1];
}
```

```go [sol-Go]
func maxFreqSum(s string) int {
	const vowelMask = 0x104111
	cnt := [26]int{}
	maxCnt := [2]int{}
	for _, ch := range s {
		ch -= 'a'
		bit := vowelMask >> ch & 1
		cnt[ch]++
		maxCnt[bit] = max(maxCnt[bit], cnt[ch])
	}
	return maxCnt[0] + maxCnt[1]
}
```

```js [sol-JavaScript]
var maxFreqSum = function(s) {
    const VOWEL_MASK = 0x104111;
    const cnt = Array(26).fill(0);
    const maxCnt = [0, 0];
    for (const ch of s) {
        const idx = ch.charCodeAt(0) - 'a'.charCodeAt(0);
        const bit = VOWEL_MASK >> idx & 1;
        cnt[idx]++;
        maxCnt[bit] = Math.max(maxCnt[bit], cnt[idx]);
    }
    return maxCnt[0] + maxCnt[1];
};
```

```rust [sol-Rust]
impl Solution {
    pub fn max_freq_sum(s: String) -> i32 {
        const VOWEL_MASK: usize = 0x104111;
        let mut cnt = [0; 26];
        let mut max_cnt = [0; 2];
        for ch in s.bytes() {
            let idx = (ch as u8 - b'a') as usize;
            let bit = VOWEL_MASK >> idx & 1;
            cnt[idx] += 1;
            max_cnt[bit] = max_cnt[bit].max(cnt[idx]);
        }
        max_cnt[0] + max_cnt[1]
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
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
