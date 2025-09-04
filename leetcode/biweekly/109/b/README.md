示例 1 的 $s = \texttt{lEetcOde}$，其中元音字母为 $\texttt{EeOe}$，排序后为 $\texttt{EOee}$。（大写字母排前面是因为大写字母的 ASCII 值更小）

原来的 $s = \texttt{l\underline{Ee}tc\underline{O}d\underline{e}}$ 视作 $\texttt{l}\_\_\texttt{tc}\_\texttt{d}\_$，包含 $4$ 个空位。

在空位依次填入 $\texttt{EOee}$，得到答案 $\texttt{l\underline{EO}tc\underline{e}d\underline{e}}$。

## 写法一

```py [sol-Python3]
class Solution:
    def sortVowels(self, s: str) -> str:
        vowels = sorted(ch for ch in s if ch in "AEIOUaeiou")
        t = list(s)  # str 无法修改，转成 list
        j = 0
        for i, ch in enumerate(t):
            if ch in "AEIOUaeiou":
                t[i] = vowels[j]  # 填空
                j += 1
        return ''.join(t)
```

```java [sol-Java]
class Solution {
    public String sortVowels(String S) {
        StringBuilder vowels = new StringBuilder();
        char[] s = S.toCharArray();
        for (char ch : s) {
            char c = Character.toLowerCase(ch);
            if (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u') {
                vowels.append(ch);
            }
        }

        char[] sortedVowels = vowels.toString().toCharArray();
        Arrays.sort(sortedVowels);

        int j = 0;
        for (int i = 0; i < s.length; i++) {
            char c = Character.toLowerCase(s[i]);
            if (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u') {
                s[i] = sortedVowels[j++];
            }
        }
        return new String(s);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string sortVowels(string s) {
        string vowels;
        for (char ch : s) {
            char c = tolower(ch);
            if (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u') {
                vowels += ch;
            }
        }

        ranges::sort(vowels);

        int j = 0;
        for (char& ch : s) {
            char c = tolower(ch);
            if (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u') {
                ch = vowels[j++];
            }
        }
        return s;
    }
};
```

```c [sol-C]
#define VOWEL_MASK 0x208222

int cmp(const void* a, const void* b) {
    return *(char*)a - *(char*)b;
}

char* sortVowels(char* s) {
    int n = strlen(s);
    char* vowels = malloc(n * sizeof(char));
    int k = 0;
    for (int i = 0; i < n; i++) {
        char c = tolower(s[i]);
        if (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u') {
            vowels[k++] = s[i];
        }
    }

    qsort(vowels, k, sizeof(char), cmp);

    k = 0;
    for (int i = 0; i < n; i++) {
        char c = tolower(s[i]);
        if (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u') {
            s[i] = vowels[k++];
        }
    }

    free(vowels);
    return s;
}
```

```go [sol-Go]
func sortVowels(s string) string {
	vowels := []byte{}
	for _, ch := range s {
		c := unicode.ToLower(ch)
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			vowels = append(vowels, byte(ch))
		}
	}

	slices.Sort(vowels)

	t := []byte(s)
	j := 0
	for i, ch := range t {
		c := unicode.ToLower(rune(ch))
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			t[i] = vowels[j]
			j++
		}
	}
	return string(t)
}
```

```js [sol-JavaScript]
var sortVowels = function(s) {
    const vowels = [];
    for (const ch of s) {
        if ("AEIOUaeiou".includes(ch)) {
            vowels.push(ch);
        }
    }

    vowels.sort();

    let j = 0;
    const t = s.split('');
    for (let i = 0; i < t.length; i++) {
        if ("AEIOUaeiou".includes(t[i])) {
            t[i] = vowels[j++];
        }
    }
    return t.join('');
};
```

```rust [sol-Rust]
impl Solution {
    pub fn sort_vowels(s: String) -> String {
        let mut vowels = s.bytes()
            .filter(|&ch| "AEIOUaeiou".contains(ch as char))
            .collect::<Vec<_>>();

        vowels.sort_unstable();

        let mut s = s.into_bytes();
        let mut j = 0;
        for ch in s.iter_mut() {
            if "AEIOUaeiou".contains(*ch as char) {
                *ch = vowels[j];
                j += 1;
            }
        }
        unsafe { String::from_utf8_unchecked(s) }
    }
}
```

## 写法二（优化）

查看 ASCII 表可知，$\texttt{A}$ 到 $\texttt{Z}$ 的 ASCII 值的二进制低 $5$ 位是 $1$ 到 $26$，$\texttt{a}$ 到 $\texttt{z}$ 的 ASCII 值的二进制低 $5$ 位也是 $1$ 到 $26$。

所以可以用 `ch & 31` 把字母 $\textit{ch}$ 转成 $1$ 到 $26$，无论 $\textit{ch}$ 是大写还是小写，规则是统一的。

由于 $\texttt{aeiou}$ 分别是第 $1,5,7,15,21$ 个字母，根据 [从集合论到位运算](https://leetcode.cn/circle/discuss/CaOJ45/)，我们可以把元音集合

$$
\{\texttt{a},\texttt{e},\texttt{i},\texttt{o},\texttt{u}\}
$$

视作数字

$$
2^1 + 2^5 + 2^7 + 2^{15} + 2^{21} = 2130466
$$

即十六进制的 $\texttt{0x208222}$。

可以用位运算快速判断字母是否在元音集合中。

```py [sol-Python3]
class Solution:
    def sortVowels(self, s: str) -> str:
        VOWEL_MASK = 0x208222
        is_vowel = lambda ch: VOWEL_MASK >> (ord(ch) & 31) & 1

        vowels = sorted(filter(is_vowel, s))
        t = list(s)  # str 无法修改，转成 list
        j = 0
        for i, ch in enumerate(t):
            if is_vowel(ch):
                t[i] = vowels[j]  # 填空
                j += 1
        return ''.join(t)
```

```java [sol-Java]
class Solution {
    public String sortVowels(String S) {
        final int VOWEL_MASK = 0x208222;

        char[] s = S.toCharArray();
        byte[] vowels = new byte[s.length]; // 比 StringBuilder 快
        int k = 0;
        for (char ch : s) {
            if ((VOWEL_MASK >> (ch & 31) & 1) > 0) {
                vowels[k++] = (byte) ch;
            }
        }

        Arrays.sort(vowels, 0, k);

        k = 0;
        for (int i = 0; i < s.length; i++) {
            if ((VOWEL_MASK >> (s[i] & 31) & 1) > 0) {
                s[i] = (char) vowels[k++];
            }
        }
        return new String(s);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string sortVowels(string s) {
        const int VOWEL_MASK = 0x208222;
        string vowels;
        for (char ch : s) {
            if (VOWEL_MASK >> (ch & 31) & 1) { // ch 是元音
                vowels += ch;
            }
        }

        ranges::sort(vowels);

        int j = 0;
        for (char& ch : s) {
            if (VOWEL_MASK >> (ch & 31) & 1) { // ch 是元音
                ch = vowels[j++];
            }
        }
        return s;
    }
};
```

```c [sol-C]
#define VOWEL_MASK 0x208222

int cmp(const void* a, const void* b) {
    return *(char*)a - *(char*)b;
}

char* sortVowels(char* s) {
    int n = strlen(s);
    char* vowels = malloc(n * sizeof(char));
    int k = 0;
    for (int i = 0; i < n; i++) {
        if (VOWEL_MASK >> (s[i] & 31) & 1) {
            vowels[k++] = s[i];
        }
    }

    qsort(vowels, k, sizeof(char), cmp);

    k = 0;
    for (int i = 0; i < n; i++) {
        if (VOWEL_MASK >> (s[i] & 31) & 1) {
            s[i] = vowels[k++];
        }
    }

    free(vowels);
    return s;
}
```

```go [sol-Go]
func sortVowels(s string) string {
	const vowelMask = 0x208222
	vowels := []byte{}
	for _, ch := range s {
		if vowelMask>>(ch&31)&1 > 0 { // ch 是元音
			vowels = append(vowels, byte(ch))
		}
	}

	slices.Sort(vowels)

	t := []byte(s)
	j := 0
	for i, ch := range t {
		if vowelMask>>(ch&31)&1 > 0 { // ch 是元音
			t[i] = vowels[j]
			j++
		}
	}
	return string(t)
}
```

```js [sol-JavaScript]
var sortVowels = function(s) {
    const VOWEL_MASK = 0x208222;
    const vowels = [];
    for (const ch of s) {
        if (VOWEL_MASK >> (ch.charCodeAt(0) & 31) & 1) {
            vowels.push(ch);
        }
    }

    vowels.sort();

    const t = s.split('');
    let j = 0;
    for (let i = 0; i < t.length; i++) {
        if (VOWEL_MASK >> (t[i].charCodeAt(0) & 31) & 1) {
            t[i] = vowels[j++];
        }
    }
    return t.join('');
};
```

```rust [sol-Rust]
impl Solution {
    pub fn sort_vowels(s: String) -> String {
        const VOWEL_MASK: u32 = 0x208222;
        let mut vowels = s.bytes()
            .filter(|&ch| VOWEL_MASK >> (ch & 31) & 1 > 0)
            .collect::<Vec<_>>();

        vowels.sort_unstable();

        let mut s = s.into_bytes();
        let mut j = 0;
        for ch in s.iter_mut() {
            if VOWEL_MASK >> (*ch & 31) & 1 > 0 {
                *ch = vowels[j];
                j += 1;
            }
        }
        unsafe { String::from_utf8_unchecked(s) }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 写法三：计数排序

```py [sol-Python3]
class Solution:
    def sortVowels(self, s: str) -> str:
        VOWELS = "AEIOUaeiou"
        cnt = Counter(ch for ch in s if ch in VOWELS)

        it = iter(VOWELS)
        cur = next(it)

        t = list(s)  # str 无法修改，转成 list
        for i, ch in enumerate(t):
            if ch in VOWELS:
                if cnt[cur] == 0:
                    # 找下一个出现次数大于 0 的元音字母
                    cur = next(c for c in it if cnt[c])
                t[i] = cur
                cnt[cur] -= 1
        return ''.join(t)
```

```java [sol-Java]
class Solution {
    public String sortVowels(String S) {
        final int VOWEL_MASK = 0x208222;

        char[] s = S.toCharArray();
        int[] cnt = new int['u' + 1];
        for (char ch : s) {
            if ((VOWEL_MASK >> (ch & 31) & 1) > 0) {
                cnt[ch]++;
            }
        }

        int j = 'A';
        for (int i = 0; i < s.length; i++) {
            if ((VOWEL_MASK >> (s[i] & 31) & 1) == 0) {
                continue;
            }
            // 找下一个出现次数大于 0 的元音字母
            while (cnt[j] == 0) {
                j = j == 'Z' ? 'a' : j + 1;
            }
            s[i] = (char) j;
            cnt[j]--;
        }
        return new String(s);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string sortVowels(string s) {
        const int VOWEL_MASK = 0x208222;
        int cnt['u' + 1]{};
        for (char ch : s) {
            if (VOWEL_MASK >> (ch & 31) & 1) {
                cnt[ch]++;
            }
        }

        char j = 'A';
        for (char& ch : s) {
            if ((VOWEL_MASK >> (ch & 31) & 1) == 0) {
                continue;
            }
            // 找下一个出现次数大于 0 的元音字母
            while (cnt[j] == 0) {
                j = j == 'Z' ? 'a' : j + 1;
            }
            ch = j;
            cnt[j]--;
        }
        return s;
    }
};
```

```c [sol-C]
#define VOWEL_MASK 0x208222

char* sortVowels(char* s) {
    int cnt['z' + 1] = {};
    for (int i = 0; s[i]; i++) {
        if (VOWEL_MASK >> (s[i] & 31) & 1) {
            cnt[s[i]]++;
        }
    }

    char j = 'A';
    for (int i = 0; s[i]; i++) {
        if ((VOWEL_MASK >> (s[i] & 31) & 1) == 0) {
            continue;
        }
        // 找下一个出现次数大于 0 的元音字母
        while (cnt[j] == 0) {
            j = j == 'Z' ? 'a' : j + 1;
        }
        s[i] = j;
        cnt[j]--;
    }
    return s;
}
```

```go [sol-Go]
func sortVowels(s string) string {
	const vowelMask = 0x208222
	cnt := ['u' + 1]int{}
	for _, ch := range s {
		if vowelMask>>(ch&31)&1 > 0 {
			cnt[ch]++
		}
	}

	t := []byte(s)
	j := byte('A')
	for i, ch := range t {
		if vowelMask>>(ch&31)&1 == 0 {
			continue
		}
		// 找下一个出现次数大于 0 的元音字母
		for cnt[j] == 0 {
			if j == 'Z' {
				j = 'a'
			} else {
				j++
			}
		}
		t[i] = j
		cnt[j]--
	}
	return string(t)
}
```

```js [sol-JavaScript]
var sortVowels = function(s) {
    const VOWEL_MASK = 0x208222;
    const cnt = Array('u'.charCodeAt(0) + 1).fill(0);
    for (const ch of s) {
        const c = ch.charCodeAt(0);
        if (VOWEL_MASK >> (c & 31) & 1) {
            cnt[c]++;
        }
    }

    const t = s.split('');
    const ordZ = 'Z'.charCodeAt(0);
    let j = 'A'.charCodeAt(0);
    for (let i = 0; i < t.length; i++) {
        if ((VOWEL_MASK >> (t[i].charCodeAt(0) & 31) & 1) === 0) {
            continue;
        }
        // 找下一个出现次数大于 0 的元音字母
        while (cnt[j] === 0) {
            j = j == ordZ ? 'a'.charCodeAt(0) : j + 1;
        }
        t[i] = String.fromCharCode(j);
        cnt[j]--;
    }
    return t.join('');
};
```

```rust [sol-Rust]
impl Solution {
    pub fn sort_vowels(s: String) -> String {
        const VOWEL_MASK: u32 = 0x208222;
        let mut cnt = [0; 'z' as usize + 1];
        for ch in s.bytes() {
            if (VOWEL_MASK >> (ch & 31)) & 1 > 0 {
                cnt[ch as usize] += 1;
            }
        }

        let mut s = s.into_bytes();
        let mut j = 0;
        for ch in s.iter_mut() {
            if VOWEL_MASK >> (*ch & 31) & 1 == 0 {
                continue;
            }
            // 找下一个出现次数大于 0 的元音字母
            while cnt[j as usize] == 0 {
                if j == b'Z' {
                    j = b'a';
                } else {
                    j += 1;
                }
            }
            *ch = j;
            cnt[j as usize] -= 1;
        }
        unsafe { String::from_utf8_unchecked(s) }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=10$ 或 $52$ 或 $128$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

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
