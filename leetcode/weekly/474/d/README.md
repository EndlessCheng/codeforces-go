首先，$s$ 不能有超过一个字母出现奇数次，否则无法形成回文串。

回文串的特点是，一旦确定了左半的字母，就确定了右半的字母，所以只需考虑左半字母如何排列。也就是对于左半，计算 [3720. 大于目标字符串的最小字典序排列](https://leetcode.cn/problems/lexicographically-smallest-permutation-greater-than-target/)。

特殊情况：如果把 $\textit{target}$ 的左半翻转到右半，就比 $\textit{target}$ 大，那么 $\textit{target}$ 的左半是不需要变的（而 3720 题必须变大）。我们特判这种情况，其余代码逻辑同 3720 题。

[本题视频讲解](https://www.bilibili.com/video/BV1MgyfBoEuX/?t=47m49s)，欢迎点赞关注~

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
        # 先假设答案左半与 t 的左半（不含正中间）相同
        for i, b in enumerate(target[:n // 2]):
            left[b] -= 2

        if valid():
            # 特殊情况：把 target 左半翻转到右半，能否比 target 大？
            left_s = target[:n // 2]
            right_s = mid_ch + left_s[::-1]
            if right_s > target[n // 2:]:  # 由于左半是一样的，所以只需比右半
                return left_s + right_s

        for i in range(n // 2 - 1, -1, -1):
            b = target[i]
            left[b] += 2  # 撤销消耗
            if not valid():  # [0,i-1] 无法做到全部一样
                continue

            # 把 target[i] 增大到 j
            for j in range(ord(b) - ord('a') + 1, 26):
                ch = ascii_lowercase[j]
                if left[ch] == 0:
                    continue

                # 找到答案（下面的循环在整个算法中只会跑一次）
                left[ch] -= 2
                ans = list(target[:i + 1])
                ans[i] = ch

                # 中间可以随便填
                for ch in ascii_lowercase:
                    ans.extend(ch * (left[ch] // 2))

                # 镜像翻转
                right_s = ans[::-1]
                ans.append(mid_ch)
                ans.extend(right_s)

                return ''.join(ans)
            # 增大失败，继续枚举

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

        String midCh = "";
        for (int i = 0; i < 26; i++) {
            int c = left[i];
            if (c % 2 == 0) {
                continue;
            }
            // s 不能有超过一个字母出现奇数次
            if (!midCh.isEmpty()) {
                return "";
            }
            // 记录填在正中间的字母
            midCh = "" + (char) ('a' + i);
            left[i]--;
        }

        int n = s.length();
        // 先假设答案左半与 target 的左半（不含正中间）相同
        for (int i = 0; i < n / 2; i++) {
            left[target.charAt(i) - 'a'] -= 2;
        }

        if (valid(left)) {
            // 特殊情况：把 target 左半翻转到右半，能否比 target 大？
            String leftS = target.substring(0, n / 2);
            String rightS = midCh + new StringBuilder(leftS).reverse();
            if (rightS.compareTo(target.substring(n / 2)) > 0) { // 由于左半是一样的，所以只需比右半
                return leftS + rightS;
            }
        }

        for (int i = n / 2 - 1; i >= 0; i--) {
            int b = target.charAt(i) - 'a';
            left[b] += 2; // 撤销消耗
            if (!valid(left)) { // [0,i-1] 无法做到全部一样
                continue;
            }

            // 把 target[i] 增大到 j
            for (int j = b + 1; j < 26; j++) {
                if (left[j] == 0) {
                    continue;
                }

                // 找到答案（下面的循环在整个算法中只会跑一次）
                left[j] -= 2;
                StringBuilder ans = new StringBuilder(target.substring(0, i + 1));
                ans.setCharAt(i, (char) ('a' + j));
    
                // 中间可以随便填
                for (int k = 0; k < 26; k++) {
                    ans.repeat('a' + k, left[k] / 2);
                }
    
                // 镜像翻转
                StringBuilder rightS = new StringBuilder(ans).reverse();
                return ans.append(midCh).append(rightS).toString();
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

        string mid_ch;
        for (int i = 0; i < 26; i++) {
            int c = left[i];
            if (c % 2 == 0) {
                continue;
            }
            // s 不能有超过一个字母出现奇数次
            if (!mid_ch.empty()) {
                return "";
            }
            // 记录填在正中间的字母
            mid_ch = 'a' + i;
            left[i]--;
        }

        int n = s.size();
        // 先假设答案左半与 t 的左半（不含正中间）相同
        for (int i = 0; i < n / 2; i++) {
            left[target[i] - 'a'] -= 2;
        }

        if (valid()) {
            // 特殊情况：把 target 左半翻转到右半，能否比 target 大？
            string right_s = target.substr(0, n / 2);
            ranges::reverse(right_s);
            right_s = mid_ch + right_s;
            if (right_s > target.substr(n / 2)) { // 由于左半是一样的，所以只需比右半
                return target.substr(0, n / 2) + right_s;
            }
        }

        for (int i = n / 2 - 1; i >= 0; i--) {
            int b = target[i] - 'a';
            left[b] += 2; // 撤销消耗
            if (!valid()) { // [0,i-1] 无法做到全部一样
                continue;
            }

            // 把 target[i] 增大到 j
            for (int j = b + 1; j < 26; j++) {
                if (left[j] == 0) {
                    continue;
                }

                // 找到答案（下面的循环在整个算法中只会跑一次）
                left[j] -= 2;
                target.resize(i + 1);
                target[i] = 'a' + j;
    
                // 中间的空位可以随便填
                for (int k = 0; k < 26; k++) {
                    target += string(left[k] / 2, 'a' + k);
                }
    
                // 镜像翻转
                string right_s = target;
                ranges::reverse(right_s);
                target += mid_ch;
                target += right_s;
    
                return target;
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

	midCh := ""
	for i, c := range left {
		if c%2 == 0 {
			continue
		}
		// s 不能有超过一个字母出现奇数次
		if midCh != "" {
			return ""
		}
		// 记录填在正中间的字母
		midCh = string('a' + byte(i))
		left[i]--
	}

	n := len(s)
	// 先假设答案左半与 t 的左半（不含正中间）相同
	for _, b := range target[:n/2] {
		left[b-'a'] -= 2
	}

	if valid() {
		// 特殊情况：把 target 左半翻转到右半，能否比 target 大？
		leftS := target[:n/2]
		tmp := []byte(leftS)
		slices.Reverse(tmp)
		rightS := midCh + string(tmp)
		if rightS > target[n/2:] { // 由于左半是一样的，所以只需比右半
			return leftS + rightS
		}
	}

	for i := n/2 - 1; i >= 0; i-- {
		b := target[i] - 'a'
		left[b] += 2  // 撤销消耗
		if !valid() { // [0,i-1] 无法做到全部一样
			continue
		}

		// 把 target[i] 增大到 j
		for j := b + 1; j < 26; j++ {
			if left[j] == 0 {
				continue
			}

			// 找到答案（下面的循环在整个算法中只会跑一次）
			left[j] -= 2
			ans := []byte(target[:i+1])
			ans[i] = 'a' + j

			// 中间可以随便填
			for k, c := range left {
				ch := string('a' + byte(k))
				ans = append(ans, strings.Repeat(ch, c/2)...)
			}

			// 镜像翻转
			rightS := slices.Clone(ans)
			slices.Reverse(rightS)
			ans = append(ans, midCh...)
			ans = append(ans, rightS...)

			return string(ans)
		}
		// 增大失败，继续枚举
	}
	return ""
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(n+|\Sigma|)$。

## 优化

我们可以减少不必要的循环，快速判断能否增大 $\textit{target}[i]$：

1. 维护 $\textit{left}$ 中的负数个数 $\textit{neg}$。
2. 维护 $\textit{left}$ 中的正数个数对应的字母最大值 $\textit{leftMax}$。

如果 $\textit{neg} < 0$ 且 $\textit{leftMax} \le \textit{target}[i]$，那么无法增大 $\textit{target}[i]$。

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
        # 先假设答案左半与 t 的左半（不含正中间）相同
        for i, b in enumerate(target[:n // 2]):
            left[b] -= 2

        neg, left_max = 0, ''
        for c, cnt in left.items():
            if cnt < 0:
                neg += 1  # 统计 left 中的负数个数
            elif cnt > 0:
                left_max = max(left_max, c)  # 剩余可用字母的最大值

        if neg == 0:
            # 特殊情况：把 target 左半翻转到右半，能否比 target 大？
            left_s = target[:n // 2]
            right_s = mid_ch + left_s[::-1]
            if right_s > target[n // 2:]:  # 由于左半是一样的，所以只需比右半
                return left_s + right_s

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

            # 把 target[i] 增大到 ch
            ch = ascii_lowercase[j]
            left[ch] -= 2
            ans = list(target[:i + 1])
            ans[i] = ch

            # 中间可以随便填
            for ch in ascii_lowercase:
                ans.extend(ch * (left[ch] // 2))

            # 镜像翻转
            right_s = ans[::-1]
            ans.append(mid_ch)
            ans.extend(right_s)

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

        String midCh = "";
        for (int i = 0; i < 26; i++) {
            int c = left[i];
            if (c % 2 == 0) {
                continue;
            }
            // s 不能有超过一个字母出现奇数次
            if (!midCh.isEmpty()) {
                return "";
            }
            // 记录填在正中间的字母
            midCh = "" + (char) ('a' + i);
            left[i]--;
        }

        int n = s.length();
        // 先假设答案左半与 target 的左半（不含正中间）相同
        for (int i = 0; i < n / 2; i++) {
            left[target.charAt(i) - 'a'] -= 2;
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
            // 特殊情况：把 target 左半翻转到右半，能否比 target 大？
            String leftS = target.substring(0, n / 2);
            String rightS = midCh + new StringBuilder(leftS).reverse();
            if (rightS.compareTo(target.substring(n / 2)) > 0) { // 由于左半是一样的，所以只需比右半
                return leftS + rightS;
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

            // 把 target[i] 增大到 j
            left[j] -= 2;
            StringBuilder ans = new StringBuilder(target.substring(0, i + 1));
            ans.setCharAt(i, (char) ('a' + j));

            // 中间可以随便填
            for (int k = 0; k < 26; k++) {
                ans.repeat('a' + k, left[k] / 2);
            }

            // 镜像翻转
            StringBuilder rightS = new StringBuilder(ans).reverse();
            return ans.append(midCh).append(rightS).toString();
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

        string mid_ch;
        for (int i = 0; i < 26; i++) {
            int c = left[i];
            if (c % 2 == 0) {
                continue;
            }
            // s 不能有超过一个字母出现奇数次
            if (!mid_ch.empty()) {
                return "";
            }
            // 记录填在正中间的字母
            mid_ch = 'a' + i;
            left[i]--;
        }

        int n = s.size();
        // 先假设答案左半与 t 的左半（不含正中间）相同
        for (int i = 0; i < n / 2; i++) {
            left[target[i] - 'a'] -= 2;
        }

        int neg = 0, left_max = 0;
        for (int i = 0; i < 26; i++) {
            if (left[i] < 0) {
                neg++; // 统计 left 中的负数个数
            } else if (left[i] > 0) {
                left_max = max(left_max, i); // 剩余可用字母的最大值
            }
        }

        if (neg == 0) {
            // 特殊情况：把 target 左半翻转到右半，能否比 target 大？
            string right_s = target.substr(0, n / 2);
            ranges::reverse(right_s);
            right_s = mid_ch + right_s;
            if (right_s > target.substr(n / 2)) { // 由于左半是一样的，所以只需比右半
                return target.substr(0, n / 2) + right_s;
            }
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

            // 把 target[i] 增大到 j
            left[j] -= 2;
            target.resize(i + 1);
            target[i] = 'a' + j;

            // 中间的空位可以随便填
            for (int k = 0; k < 26; k++) {
                target += string(left[k] / 2, 'a' + k);
            }

            // 镜像翻转
            string right_s = target;
            ranges::reverse(right_s);
            target += mid_ch;
            target += right_s;

            return target;
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

	midCh := ""
	for i, c := range left {
		if c%2 == 0 {
			continue
		}
		// s 不能有超过一个字母出现奇数次
		if midCh != "" {
			return ""
		}
		// 记录填在正中间的字母
		midCh = string('a' + byte(i))
		left[i]--
	}

	n := len(s)
	// 先假设答案左半与 t 的左半（不含正中间）相同
	for _, b := range target[:n/2] {
		left[b-'a'] -= 2
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
		// 特殊情况：把 target 左半翻转到右半，能否比 target 大？
		leftS := target[:n/2]
		tmp := []byte(leftS)
		slices.Reverse(tmp)
		rightS := midCh + string(tmp)
		if rightS > target[n/2:] { // 由于左半是一样的，所以只需比右半
			return leftS + rightS
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

		// 把 target[i] 增大到 j
		left[j] -= 2
		ans := []byte(target[:i+1])
		ans[i] = 'a' + j

		// 中间可以随便填
		for k, c := range left {
			ch := string('a' + byte(k))
			ans = append(ans, strings.Repeat(ch, c/2)...)
		}

		// 镜像翻转
		rightS := slices.Clone(ans)
		slices.Reverse(rightS)
		ans = append(ans, midCh...)
		ans = append(ans, rightS...)

		return string(ans)
	}
	return ""
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(n+|\Sigma|)$。

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
