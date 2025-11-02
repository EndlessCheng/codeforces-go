请先完成更简单的 [3720. 大于目标字符串的最小字典序排列](https://leetcode.cn/problems/lexicographically-smallest-permutation-greater-than-target/)。

对于本题，首先 $s$ 不能有超过一个字母出现奇数次，否则无法形成回文串。

我们可以先把 $\textit{target}$ 左半翻转到右半，看看能否比 $\textit{target}$ 大。如果可以且 $s$ 字母足够，直接返回 $\textit{target}$ 左半翻转到右半的结果。

否则和 3720 题一样，倒序贪心。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

## 优化前

```py [sol-Python3]
class Solution:
    def lexPalindromicPermutation(self, s: str, target: str) -> str:
        left = Counter(s)

        def valid() -> bool:
            return all(c >= 0 for c in left.values())

        mid_ch = ''
        for ch, c in left.items():
            if c % 2 == 0:
                continue
            # s 不能有超过一个字母出现奇数次
            if mid_ch:
                return ""
            # 记录填在正中间的字母
            mid_ch = ch
            left[ch] -= 1

        n = len(s)
        ans = list(target)
        # 先假设答案左半与 t 的左半（不含正中间）相同
        for i, b in enumerate(target[:n // 2]):
            left[b] -= 2
            ans[-1 - i] = b  # 把 target 左半翻转到右半
        # 正中间只能填那个出现奇数次的字母
        if mid_ch:
            ans[n // 2] = mid_ch

        # 把 target 左半翻转到右半，能否比 target 大？
        if valid() and (t := ''.join(ans)) > target:
            return t

        for i in range(n // 2 - 1, -1, -1):
            b = target[i]
            left[b] += 2  # 撤销消耗
            if not valid():  # [0,i-1] 无法做到全部一样
                continue

            # 把 target[i] 和 target[n-1-i] 都增大到 j
            for j in range(ord(b) - ord('a') + 1, 26):
                ch = ascii_lowercase[j]
                if left[ch] == 0:
                    continue

                # 找到答案（下面的循环在整个算法中只会跑一次）
                left[ch] -= 2
                ans[i] = ans[-1 - i] = ch
                right = ans[-1 - i:]
                del ans[i + 1:]

                # 中间的空位可以随便填
                t = []
                for ch in ascii_lowercase:
                    t.extend(ch * (left[ch] // 2))

                ans += t
                if mid_ch:
                    ans.append(mid_ch)
                ans += t[::-1]
                ans += right

                return ''.join(ans)
        return ""
```

```java [sol-Java]
class Solution {
    public String lexPalindromicPermutation(String s, String target) {
        int[] left = new int[26];
        for (char b : s.toCharArray()) {
            left[b - 'a']++;
        }
        if (!valid(left)) {
            return "";
        }

        char midCh = 0;
        for (int i = 0; i < 26; i++) {
            int c = left[i];
            if (c % 2 == 0) {
                continue;
            }
            // s 不能有超过一个字母出现奇数次
            if (midCh > 0) {
                return "";
            }
            // 记录填在正中间的字母
            midCh = (char) ('a' + i);
            left[i]--;
        }

        int n = s.length();
        StringBuilder ans = new StringBuilder(target);
        // 先假设答案左半与 target 的左半（不含正中间）相同
        for (int i = 0; i < n / 2; i++) {
            char b = target.charAt(i);
            left[b - 'a'] -= 2;
            ans.setCharAt(n - 1 - i, b); // 把 target 左半翻转到右半
        }
        // 正中间只能填那个出现奇数次的字母
        if (midCh > 0) {
            ans.setCharAt(n / 2, midCh);
        }

        if (valid(left)) {
            // 把 target 左半翻转到右半，能否比 target 大？
            String t = ans.toString();
            if (t.compareTo(target) > 0) {
                return t;
            }
        }

        for (int i = n / 2 - 1; i >= 0; i--) {
            int b = target.charAt(i) - 'a';
            left[b] += 2; // 撤销消耗
            if (!valid(left)) { // [0,i-1] 无法做到全部一样
                continue;
            }

            // 把 target[i] 和 target[n-1-i] 都增大到 j
            for (int j = b + 1; j < 26; j++) {
                if (left[j] == 0) {
                    continue;
                }

                // 找到答案（下面的循环在整个算法中只会跑一次）
                left[j] -= 2;
                ans.setCharAt(i, (char) ('a' + j));
                ans.setCharAt(n - 1 - i, (char) ('a' + j));

                String right = ans.substring(n - 1 - i);
                ans.setLength(i + 1);

                // 中间的空位可以随便填
                StringBuilder t = new StringBuilder();
                for (int k = 0; k < 26; k++) {
                    t.repeat('a' + k, left[k] / 2);
                }

                ans.append(t);
                if (midCh > 0) {
                    ans.append(midCh);
                }
                ans.append(t.reverse());
                ans.append(right);

                return ans.toString();
            }
            // 增大失败，继续枚举
        }
        return "";
    }

    private boolean valid(int[] left) {
        for (int c : left) {
            if (c < 0) {
                return false;
            }
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string lexPalindromicPermutation(string s, string target) {
        int left[26]{};
        for (char b : s) {
            left[b - 'a']++;
        }
        auto valid = [&]() -> bool {
            for (int c : left) {
                if (c < 0) {
                    return false;
                }
            }
            return true;
        };

        char mid_ch = 0;
        for (int i = 0; i < 26; i++) {
            int c = left[i];
            if (c % 2 == 0) {
                continue;
            }
            // s 不能有超过一个字母出现奇数次
            if (mid_ch) {
                return "";
            }
            // 记录填在正中间的字母
            mid_ch = 'a' + i;
            left[i]--;
        }

        int n = s.size();
        string ans = target;
        // 先假设答案左半与 t 的左半（不含正中间）相同
        for (int i = 0; i < n / 2; i++) {
            char b = target[i];
            left[b - 'a'] -= 2;
            ans[n - 1 - i] = b; // 把 target 左半翻转到右半
        }
        // 正中间只能填那个出现奇数次的字母
        if (mid_ch > 0) {
            ans[n / 2] = mid_ch;
        }

        // 把 target 左半翻转到右半，能否比 target 大？
        if (valid() && ans > target) {
            return ans;
        }

        for (int i = n / 2 - 1; i >= 0; i--) {
            int b = target[i] - 'a';
            left[b] += 2; // 撤销消耗
            if (!valid()) { // [0,i-1] 无法做到全部一样
                continue;
            }

            // 把 target[i] 和 target[n-1-i] 都增大到 j
            for (int j = b + 1; j < 26; j++) {
                if (left[j] == 0) {
                    continue;
                }

                // 找到答案（下面的循环在整个算法中只会跑一次）
                left[j] -= 2;
                ans[i] = ans[n - 1 - i] = 'a' + j;
                string right = ans.substr(n - 1 - i);
                ans.resize(i + 1);

                // 中间的空位可以随便填
                string t;
                for (int k = 0; k < 26; k++) {
                    t += string(left[k] / 2, 'a' + k);
                }

                ans += t;
                if (mid_ch) {
                    ans += mid_ch;
                }
                ranges::reverse(t);
                ans += t;
                ans += right;

                return ans;
            }
            // 增大失败，继续枚举
        }
        return "";
    }
};
```

```go [sol-Go]
func lexPalindromicPermutation(s, target string) string {
	left := make([]int, 26)
	for _, b := range s {
		left[b-'a']++
	}
	valid := func() bool {
		for _, c := range left {
			if c < 0 {
				return false
			}
		}
		return true
	}

	midCh := byte(0)
	for i, c := range left {
		if c%2 == 0 {
			continue
		}
		// s 不能有超过一个字母出现奇数次
		if midCh > 0 {
			return ""
		}
		// 记录填在正中间的字母
		midCh = 'a' + byte(i)
		left[i]--
	}

	n := len(s)
	ans := []byte(target)
	// 先假设答案左半与 t 的左半（不含正中间）相同
	for i, b := range target[:n/2] {
		left[b-'a'] -= 2
		ans[n-1-i] = byte(b) // 把 target 左半翻转到右半
	}
	// 正中间只能填那个出现奇数次的字母
	if midCh > 0 {
		ans[n/2] = midCh
	}

	if valid() {
		// 把 target 左半翻转到右半，能否比 target 大？
		t := string(ans)
		if t > target {
			return t
		}
	}

	for i := n/2 - 1; i >= 0; i-- {
		b := target[i] - 'a'
		left[b] += 2 // 撤销消耗
		if !valid() { // [0,i-1] 无法做到全部一样
			continue
		}

		// 把 target[i] 和 target[n-1-i] 都增大到 j
		for j := b + 1; j < 26; j++ {
			if left[j] == 0 {
				continue
			}

			// 找到答案（下面的循环在整个算法中只会跑一次）
			left[j] -= 2
			ans[i] = 'a' + j
			ans[n-1-i] = ans[i]

			// 中间的空位可以随便填
			t := make([]byte, 0, n-(i+1)*2)
			for k, c := range left {
				ch := string('a' + byte(k))
				t = append(t, strings.Repeat(ch, c/2)...)
			}

			// 把 t、midCh、Reverse(t) 依次填在 ans[i] 的右边
			a := append(ans[:i+1], t...)
			if midCh > 0 {
				a = append(a, midCh)
			}
			slices.Reverse(t)
			a = append(a, t...)

			return string(ans)
		}
		// 增大失败，继续枚举
	}
	return ""
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|)$，其中 $n$ 是 $\textit{nums}$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。返回值不计入。

## 优化

我们可以减少不必要的循环，快速判断能否增大 $\textit{target}[i]$：

1. 维护 $\textit{left}$ 中的负数个数 $\textit{neg}$。
2. 维护 $\textit{left}$ 中的正数个数对应的字母最大值 $\textit{mx}$。

如果 $\textit{neg} < 0$ 且 $\textit{target}[i] \ge \textit{mx}$，那么无法增大 $\textit{target}[i]$。

```py [sol-Python3]
class Solution:
    def lexPalindromicPermutation(self, s: str, target: str) -> str:
        left = Counter(s)

        def valid() -> bool:
            return all(c >= 0 for c in left.values())

        mid_ch = ''
        for ch, c in left.items():
            if c % 2 == 0:
                continue
            # s 不能有超过一个字母出现奇数次
            if mid_ch:
                return ""
            # 记录填在正中间的字母
            mid_ch = ch
            left[ch] -= 1

        n = len(s)
        ans = list(target)
        # 先假设答案左半与 t 的左半（不含正中间）相同
        for i, b in enumerate(target[:n // 2]):
            left[b] -= 2
            ans[-1 - i] = b  # 把 target 左半翻转到右半
        # 正中间只能填那个出现奇数次的字母
        if mid_ch:
            ans[n // 2] = mid_ch

        neg, left_max = 0, ''
        for c, cnt in left.items():
            if cnt < 0:
                neg += 1  # 统计 left 中的负数个数
            elif cnt > 0:
                left_max = max(left_max, c)  # 剩余可用字母的最大值

        # 把 target 左半翻转到右半，能否比 target 大？
        if neg == 0 and (t := ''.join(ans)) > target:
            return t

        for i in range(n // 2 - 1, -1, -1):
            b = target[i]
            left[b] += 2  # 撤销消耗

            if left[b] == 0:
                neg -= 1
            elif left[b] == 2:
                left_max = max(left_max, b)

            # left 有负数 or 没有大于 target[i] 的字母
            if neg > 0 or left_max <= b:
                continue

            # 找到答案（下面的循环在整个算法中只会跑一次）
            j = ord(b) - ord('a') + 1
            while left[ascii_lowercase[j]] == 0:
                j += 1

            # 把 target[i] 和 target[n-1-i] 增大到 ch
            ch = ascii_lowercase[j]
            left[ch] -= 2
            ans[i] = ans[-1 - i] = ch
            right = ans[-1 - i:]
            del ans[i + 1:]

            # 中间的空位可以随便填
            t = []
            for ch in ascii_lowercase:
                t.extend(ch * (left[ch] // 2))

            ans += t
            if mid_ch:
                ans.append(mid_ch)
            ans += t[::-1]
            ans += right

            return ''.join(ans)
        return ""
```

```java [sol-Java]
class Solution {
    public String lexPalindromicPermutation(String s, String target) {
        int[] left = new int[26];
        for (char b : s.toCharArray()) {
            left[b - 'a']++;
        }

        char midCh = 0;
        for (int i = 0; i < 26; i++) {
            int c = left[i];
            if (c % 2 == 0) {
                continue;
            }
            // s 不能有超过一个字母出现奇数次
            if (midCh > 0) {
                return "";
            }
            // 记录填在正中间的字母
            midCh = (char) ('a' + i);
            left[i]--;
        }

        int n = s.length();
        StringBuilder ans = new StringBuilder(target);
        // 先假设答案左半与 target 的左半（不含正中间）相同
        for (int i = 0; i < n / 2; i++) {
            char b = target.charAt(i);
            left[b - 'a'] -= 2;
            ans.setCharAt(n - 1 - i, b); // 把 target 左半翻转到右半
        }
        // 正中间只能填那个出现奇数次的字母
        if (midCh > 0) {
            ans.setCharAt(n / 2, midCh);
        }

        int neg = 0;
        int leftMax = 0;
        for (int i = 0; i < 26; i++) {
            if (left[i] < 0) {
                neg++; // 统计 left 中的负数个数
            } else if (left[i] > 0) {
                leftMax = Math.max(leftMax, i); // 剩余可用字母的最大值
            }
        }

        if (neg == 0) {
            // 把 target 左半翻转到右半，能否比 target 大？
            String t = ans.toString();
            if (t.compareTo(target) > 0) {
                return t;
            }
        }

        for (int i = n / 2 - 1; i >= 0; i--) {
            int b = target.charAt(i) - 'a';
            left[b] += 2; // 撤销消耗

            if (left[b] == 0) {
                neg--;
            } else if (left[b] == 2) {
                leftMax = Math.max(leftMax, b);
            }

            // left 有负数 or 没有大于 target[i] 的字母
            if (neg > 0 || leftMax <= b) {
                continue;
            }

            // 找到答案（下面的循环在整个算法中只会跑一次）
            int j = b + 1;
            while (left[j] == 0) {
                j++;
            }

            // 把 target[i] 和 target[n-1-i] 增大到 j
            left[j] -= 2;
            ans.setCharAt(i, (char) ('a' + j));
            ans.setCharAt(n - 1 - i, (char) ('a' + j));

            String right = ans.substring(n - 1 - i);
            ans.setLength(i + 1);

            // 中间的空位可以随便填
            StringBuilder t = new StringBuilder();
            for (int k = 0; k < 26; k++) {
                t.repeat('a' + k, left[k] / 2);
            }

            ans.append(t);
            if (midCh > 0) {
                ans.append(midCh);
            }
            ans.append(t.reverse());
            ans.append(right);

            return ans.toString();
        }
        return "";
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string lexPalindromicPermutation(string s, string target) {
        int left[26]{};
        for (char b : s) {
            left[b - 'a']++;
        }

        char mid_ch = 0;
        for (int i = 0; i < 26; i++) {
            int c = left[i];
            if (c % 2 == 0) {
                continue;
            }
            // s 不能有超过一个字母出现奇数次
            if (mid_ch) {
                return "";
            }
            // 记录填在正中间的字母
            mid_ch = 'a' + i;
            left[i]--;
        }

        int n = s.size();
        string ans = target;
        // 先假设答案左半与 t 的左半（不含正中间）相同
        for (int i = 0; i < n / 2; i++) {
            char b = target[i];
            left[b - 'a'] -= 2;
            ans[n - 1 - i] = b; // 把 target 左半翻转到右半
        }
        // 正中间只能填那个出现奇数次的字母
        if (mid_ch > 0) {
            ans[n / 2] = mid_ch;
        }

        int neg = 0, left_max = 0;
        for (int i = 0; i < 26; i++) {
            if (left[i] < 0) {
                neg++; // 统计 left 中的负数个数
            } else if (left[i] > 0) {
                left_max = max(left_max, i); // 剩余可用字母的最大值
            }
        }

        // 把 target 左半翻转到右半，能否比 target 大？
        if (neg == 0 && ans > target) {
            return ans;
        }

        for (int i = n / 2 - 1; i >= 0; i--) {
            int b = target[i] - 'a';
            left[b] += 2; // 撤销消耗

            if (left[b] == 0) {
                neg--;
            } else if (left[b] == 2) {
                left_max = max(left_max, b);
            }

            // left 有负数 or 没有大于 target[i] 的字母
            if (neg > 0 || left_max <= b) {
                continue;
            }

            // 找到答案（下面的循环在整个算法中只会跑一次）
            int j = b + 1;
            while (left[j] == 0) {
                j++;
            }

            // 把 target[i] 和 target[n-1-i] 增大到 j
            left[j] -= 2;
            ans[i] = ans[n - 1 - i] = 'a' + j;
            string right = ans.substr(n - 1 - i);
            ans.resize(i + 1);

            // 中间的空位可以随便填
            string t;
            for (int k = 0; k < 26; k++) {
                t += string(left[k] / 2, 'a' + k);
            }

            ans += t;
            if (mid_ch) {
                ans += mid_ch;
            }
            ranges::reverse(t);
            ans += t;
            ans += right;

            return ans;
        }
        return "";
    }
};
```

```go [sol-Go]
func lexPalindromicPermutation(s, target string) string {
	left := make([]int, 26)
	for _, b := range s {
		left[b-'a']++
	}

	midCh := byte(0)
	for i, c := range left {
		if c%2 == 0 {
			continue
		}
		// s 不能有超过一个字母出现奇数次
		if midCh > 0 {
			return ""
		}
		// 记录填在正中间的字母
		midCh = 'a' + byte(i)
		left[i]--
	}

	n := len(s)
	ans := []byte(target)
	// 先假设答案左半与 t 的左半（不含正中间）相同
	for i, b := range target[:n/2] {
		left[b-'a'] -= 2
		ans[n-1-i] = byte(b) // 把 target 左半翻转到右半
	}
	// 正中间只能填那个出现奇数次的字母
	if midCh > 0 {
		ans[n/2] = midCh
	}

	neg, leftMax := 0, byte(0)
	for i, cnt := range left {
		if cnt < 0 {
			neg++ // 统计 left 中的负数个数
		} else if cnt > 0 {
			leftMax = max(leftMax, byte(i)) // 剩余可用字母的最大值
		}
	}

	if neg == 0 {
		// 把 target 左半翻转到右半，能否比 target 大？
		t := string(ans)
		if t > target {
			return t
		}
	}

	for i := n/2 - 1; i >= 0; i-- {
		b := target[i] - 'a'
		left[b] += 2 // 撤销消耗

		if left[b] == 0 {
			neg--
		} else if left[b] == 2 {
			leftMax = max(leftMax, b)
		}

		// left 有负数 or 没有大于 target[i] 的字母
		if neg > 0 || leftMax <= b {
			continue
		}

		// 找到答案（下面的循环在整个算法中只会跑一次）
		j := b + 1
		for left[j] == 0 {
			j++
		}

		// 把 target[i] 和 target[n-1-i] 增大到 j
		left[j] -= 2
		ans[i] = 'a' + j
		ans[n-1-i] = ans[i]

		// 中间的空位可以随便填
		t := make([]byte, 0, n-(i+1)*2)
		for k, c := range left {
			ch := string('a' + byte(k))
			t = append(t, strings.Repeat(ch, c/2)...)
		}

		// 把 t、midCh、Reverse(t) 依次填在 ans[i] 的右边
		a := append(ans[:i+1], t...)
		if midCh > 0 {
			a = append(a, midCh)
		}
		slices.Reverse(t)
		a = append(a, t...)

		return string(ans)
	}
	return ""
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|)$，其中 $n$ 是 $\textit{nums}$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。返回值不计入。

## 专题训练

见下面贪心题单的「**§3.1 字典序最小/最大**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
