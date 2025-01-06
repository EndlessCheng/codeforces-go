下文把 $\textit{word}_1$ 和 $\textit{word}_2$ 简称为 $s$ 和 $t$。

由于子串可以重排，只要子串可以涵盖（见 76 题）字符串 $t$，那么子串就可以通过重排，使得 $t$ 是子串的前缀。

所以本题是 [76. 最小覆盖子串](https://leetcode.cn/problems/minimum-window-substring/) 的求个数版本，做法都是**滑动窗口**，请看 [我的题解](https://leetcode.cn/problems/minimum-window-substring/solutions/2713911/liang-chong-fang-fa-cong-o52mn-dao-omnfu-3ezz/)。

滑动窗口的内层循环结束时，右端点**固定**在 $\textit{right}$，左端点在 $0,1,2,\ldots,\textit{left}-1$ 的所有子串都是合法的，这一共有 $\textit{left}$ 个，把 $\textit{left}$ 加入答案。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1WRtDejEjD/) 第三+四题，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def validSubstringCount(self, s: str, t: str) -> int:
        if len(s) < len(t):
            return 0

        # t 的字母出现次数与 s 的字母出现次数之差
        diff = defaultdict(int)  # 也可以用 Counter(t)，但是会慢很多
        for c in t:
            diff[c] += 1

        # 窗口内有 less 个字母的出现次数比 t 的少
        less = len(diff)

        ans = left = 0
        for c in s:
            diff[c] -= 1
            if diff[c] == 0:
                # c 移入窗口后，窗口内 c 的出现次数和 t 的一样
                less -= 1
            while less == 0:  # 窗口符合要求
                if diff[s[left]] == 0:
                    # s[left] 移出窗口之前，检查出现次数，
                    # 如果窗口内 s[left] 的出现次数和 t 的一样，
                    # 那么 s[left] 移出窗口后，窗口内 s[left] 的出现次数比 t 的少
                    less += 1
                diff[s[left]] += 1
                left += 1
            ans += left
        return ans
```

```java [sol-Java]
class Solution {
    public long validSubstringCount(String S, String T) {
        if (S.length() < T.length()) {
            return 0;
        }

        char[] s = S.toCharArray();
        char[] t = T.toCharArray();
        int[] diff = new int[26]; // t 的字母出现次数与 s 的字母出现次数之差
        for (char c : t) {
            diff[c - 'a']++;
        }

        // 统计窗口内有多少个字母的出现次数比 t 的少
        int less = 0;
        for (int d : diff) {
            if (d > 0) {
                less++;
            }
        }

        long ans = 0;
        int left = 0;
        for (char c : s) {
            diff[c - 'a']--;
            if (diff[c - 'a'] == 0) {
                // c 移入窗口后，窗口内 c 的出现次数和 t 的一样
                less--;
            }
            while (less == 0) { // 窗口符合要求
                char outChar = s[left++]; // 准备移出窗口的字母
                if (diff[outChar - 'a'] == 0) {
                    // outChar 移出窗口之前检查出现次数，
                    // 如果窗口内 outChar 的出现次数和 t 的一样，
                    // 那么 outChar 移出窗口后，窗口内 outChar 的出现次数比 t 的少
                    less++;
                }
                diff[outChar - 'a']++;
            }
            ans += left;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long validSubstringCount(string s, string t) {
        if (s.length() < t.length()) {
            return 0;
        }

        int diff[26]{}; // t 的字母出现次数与 s 的字母出现次数之差
        for (char c : t) {
            diff[c - 'a']++;
        }

        // 统计窗口内有多少个字母的出现次数比 t 的少
        int less = 0;
        for (int d : diff) {
            if (d > 0) {
                less++;
            }
        }

        long long ans = 0;
        int left = 0;
        for (char c : s) {
            diff[c - 'a']--;
            if (diff[c - 'a'] == 0) {
                // c 移入窗口后，窗口内 c 的出现次数和 t 的一样
                less--;
            }
            while (less == 0) { // 窗口符合要求
                char out_char = s[left++] - 'a'; // 准备移出窗口的字母
                if (diff[out_char] == 0) {
                    // out_char 移出窗口之前，检查出现次数，
                    // 如果窗口内 out_char 的出现次数和 t 的一样，
                    // 那么 out_char 移出窗口后，窗口内 out_char 的出现次数比 t 的少
                    less++;
                }
                diff[out_char]++;
            }
            ans += left;
        }
        return ans;
    }
};
```

```c [sol-C]
long long validSubstringCount(char* s, char* t) {
    int diff[26] = {}; // t 的字母出现次数与 s 的字母出现次数之差
    for (int i = 0; t[i]; i++) {
        diff[t[i] - 'a']++;
    }

    // 统计窗口内有多少个字母的出现次数比 t 的少
    int less = 0;
    for (int i = 0; i < 26; i++) {
        if (diff[i] > 0) {
            less++;
        }
    }

    long long ans = 0;
    int left = 0;
    for (int i = 0; s[i]; i++) {
        diff[s[i] - 'a']--;
        if (diff[s[i] - 'a'] == 0) {
            // s[i] 移入窗口后，窗口内 s[i] 的出现次数和 t 的一样
            less--;
        }
        while (less == 0) { // 窗口符合要求
            char out_char = s[left++] - 'a'; // 准备移出窗口的字母
            if (diff[out_char] == 0) {
                // out_char 移出窗口之前，检查出现次数，
                // 如果窗口内 out_char 的出现次数和 t 的一样，
                // 那么 out_char 移出窗口后，窗口内 out_char 的出现次数比 t 的少
                less++;
            }
            diff[out_char]++;
        }
        ans += left;
    }
    return ans;
}
```

```go [sol-Go]
func validSubstringCount(s, t string) (ans int64) {
	if len(s) < len(t) {
		return 0
	}

	diff := [26]int{} // t 的字母出现次数与 s 的字母出现次数之差
	for _, c := range t {
		diff[c-'a']++
	}

	// 统计窗口内有多少个字母的出现次数比 t 的少
	less := 0
	for _, d := range diff {
		if d > 0 {
			less++
		}
	}

	left := 0
	for _, c := range s {
		diff[c-'a']--
		if diff[c-'a'] == 0 {
			// c 移入窗口后，窗口内 c 的出现次数和 t 的一样
			less--
		}
		for less == 0 { // 窗口符合要求
			if diff[s[left]-'a'] == 0 {
                // s[left] 移出窗口之前，检查出现次数，
                // 如果窗口内 s[left] 的出现次数和 t 的一样，
                // 那么 s[left] 移出窗口后，窗口内 s[left] 的出现次数比 t 的少
				less++
			}
			diff[s[left]-'a']++
			left++
		}
		ans += int64(left)
	}
	return
}
```

```js [sol-JavaScript]
var validSubstringCount = function(s, t) {
    if (s.length < t.length) {
        return 0;
    }

    const diff = Array(26).fill(0); // t 的字母出现次数与 s 的字母出现次数之差
    for (const c of t) {
        diff[c.charCodeAt(0) - 'a'.charCodeAt(0)]++;
    }

    // 统计窗口内有多少个字母的出现次数比 t 的少
    let less = 0;
    for (const d of diff) {
        if (d > 0) {
            less++;
        }
    }

    let ans = 0;
    let left = 0;
    for (const ch of s) {
        const c = ch.charCodeAt(0) - 'a'.charCodeAt(0);
        diff[c]--;
        if (diff[c] === 0) {
            // c 移入窗口后，窗口内 c 的出现次数和 t 的一样
            less--;
        }
        while (less === 0) { // 窗口符合要求
            const outChar = s[left++].charCodeAt(0) - 'a'.charCodeAt(0);
            if (diff[outChar] === 0) {
                // outChar 移出窗口之前，检查出现次数，
                // 如果窗口内 outChar 的出现次数和 t 的一样，
                // 那么 outChar 移出窗口后，窗口内 outChar 的出现次数比 t 的少
                less++;
            }
            diff[outChar]++;
        }
        ans += left;
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn valid_substring_count(s: String, t: String) -> i64 {
        if s.len() < t.len() {
            return 0;
        }

        let mut diff = vec![0; 26]; // t 的字母出现次数与 s 的字母出现次数之差
        for c in t.bytes() {
            diff[(c - b'a') as usize] += 1;
        }

        // 统计窗口内有多少个字母的出现次数比 t 的少
        let mut less = diff.iter().filter(|&&d| d > 0).count() as i32;

        let mut ans = 0;
        let mut left = 0;
        let s = s.as_bytes();
        for c in s {
            let c = (c - b'a') as usize;
            diff[c] -= 1;
            if diff[c] == 0 {
                // c 移入窗口后，窗口内 c 的出现次数和 t 的一样
                less -= 1;
            }
            while less == 0 { // 窗口符合要求
                let out_char = (s[left] - b'a') as usize; // 准备移出窗口的字母
                if diff[out_char] == 0 {
                    // out_char 移出窗口之前，检查出现次数，
                    // 如果窗口内 out_char 的出现次数和 t 的一样，
                    // 那么 out_char 移出窗口后，窗口内 out_char 的出现次数比 t 的少
                    less += 1;
                }
                diff[out_char] += 1;
                left += 1;
            }
            ans += left;
        }
        ans as _
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|)$，其中 $n$ 是 $\textit{nums}$ 的长度，$|\Sigma|=26$ 是字符集合的大小。虽然写了个二重循环，但是内层循环中对 $\textit{left}$ 加一的**总**执行次数不会超过 $n$ 次，所以总的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

更多相似题目，见下面滑动窗口题单中的「**§2.3 求子数组个数**」，例如 [2962. 统计最大元素出现至少 K 次的子数组](https://leetcode.cn/problems/count-subarrays-where-max-element-appears-at-least-k-times/) 等。

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
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
