[视频讲解](https://www.bilibili.com/video/BV1x4421w7ba/)

## 方法一：正向思维

#### 提示 1

由于可以随意交换字母，先把所有字母都取出来，然后考虑如何填入各个字符串。

如果一个奇数长度字符串最终是回文串，那么它**正中间**的那个字母填什么都可以。

既然如此，不妨**先把左右的字母填了，最后在往正中间填入字母**。

#### 提示 2

字符串越短，需要的字母越少。

所以按照字符串的长度从小到大填。

#### 提示 3

统计所有字符串的长度之和，减去出现次数为奇数的字母，即为可以往左右填入的字母个数 $\textit{tot}$。

把字符串按照长度从小到大排序，然后遍历。注意长为奇数的字符串，由于不考虑正中间的字母，其长度要减一。

```py [sol-Python3]
class Solution:
    def maxPalindromesAfterOperations(self, words: List[str]) -> int:
        ans = tot = 0
        cnt = Counter()
        for w in words:
            tot += len(w)
            cnt += Counter(w)
        tot -= sum(c % 2 for c in cnt.values())  # 减去出现次数为奇数的字母

        words.sort(key=len)  # 按照长度从小到大排序
        for w in words:
            tot -= len(w) // 2 * 2  # 长为奇数的字符串，长度要减一
            if tot < 0: break
            ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int maxPalindromesAfterOperations(String[] words) {
        int tot = 0;
        int mask = 0; // 奇数个数的字母集合
        for (String w : words) {
            tot += w.length();
            for (char c : w.toCharArray()) {
                mask ^= 1 << (c - 'a');
            }
        }
        tot -= Integer.bitCount(mask); // 减去出现次数为奇数的字母

        Arrays.sort(words, (a, b) -> a.length() - b.length());
        int ans = 0;
        for (String w : words) {
            tot -= w.length() / 2 * 2; // 长为奇数的字符串，长度要减一
            if (tot < 0) break;
            ans++;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxPalindromesAfterOperations(vector<string> &words) {
        int ans = 0, tot = 0, mask = 0;
        for (auto &w : words) {
            tot += w.length();
            for (char c : w) {
                mask ^= 1 << (c - 'a');
            }
        }
        tot -= __builtin_popcount(mask); // 减去出现次数为奇数的字母

        ranges::sort(words, [](const auto &a, const auto &b) {
            return a.length() < b.length();
        });
        for (auto &w : words) {
            tot -= w.length() / 2 * 2; // 长为奇数的字符串，长度要减一
            if (tot < 0) break;
            ans++;
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxPalindromesAfterOperations(words []string) (ans int) {
	tot, mask := 0, 0
	for _, w := range words {
		tot += len(w)
		for _, c := range w {
			mask ^= 1 << (c - 'a')
		}
	}
	tot -= bits.OnesCount(uint(mask)) // 减去出现次数为奇数的字母

	slices.SortFunc(words, func(a, b string) int { return len(a) - len(b) })
	for _, w := range words {
		tot -= len(w) / 2 * 2 // 长为奇数的字符串，长度要减一
		if tot < 0 { break }
		ans++
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L+n\log n)$，其中 $L$ 为字符串的长度之和，$n$ 为 $\textit{words}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。利用位运算可以做到 $\mathcal{O}(1)$ 额外空间（忽略排序的栈开销）。

## 方法二：逆向思维

最少有多少个字符串，不是回文串？

#### 提示 1

统计每个字母在所有字符串的**总**出现次数。如果每个字母的总出现次数都是偶数，所有字符串都可以是回文串吗？

分类讨论：

- 如果所有字符串的长度都是偶数，我们可以把这些字母对称放置，这样所有字符串都是回文串，例如示例 1。
- 如果有奇数长度字符串，那么这样的字符串的个数一定是偶数个（因为所有字母的出现次数都是偶数），我们先把字母对称放置，再把剩余的字母放在奇数长度字符串的正中间（剩余的字母可以随便放，不影响回文），这样所有字符串都是回文串。

所有字符串都可以是回文串，答案是 $n$。

#### 提示 2

如果有字母的出现次数是奇数呢？

从出现次数为奇数的字母中，每种各取出一个。比如 a 出现了 $3$ 次，b 出现了 $5$ 次，我们先取出一个 a 和一个 b。

- 把取出来的字母集合记作 $S$，优先把 $S$ 中的字母放在奇数长度字符串的正中间。如果奇数长度字符串的个数，不低于 $S$ 的大小，那么把 $S$ 中的字母填入正中间后，剩下的每种字母都是偶数，像提示 1 那样放置即可，答案是 $n$。
- 如果奇数长度字符串的个数，小于 $S$ 的大小，那么把 $S$ 中的一部分字母填入正中间后，$S$ 中还有剩余字母，我们可以找一个长度**最长**的字符串，放入 $S$ 中剩下的字母。如果 $S$ 中还有剩余字母，就继续找次长的字符串放入。所以需要**按照字符串的长度从大到小排序**。由于奇数长度字符串的正中间已经填入了字母，所以在填入 $S$ 中的剩余字母时，奇数长度字符串的长度要减一。答案初始化为 $n$，每用一个字符串放 $S$ 剩余字母，就把答案减一。

```py [sol-Python3]
class Solution:
    def maxPalindromesAfterOperations(self, words: List[str]) -> int:
        odd_l = 0  # 奇数长度字符串个数
        cnt = Counter()
        for w in words:
            odd_l += len(w) % 2
            cnt += Counter(w)

        words.sort(key=lambda w: -len(w))  # 按照长度从大到小排序

        ans = len(words)
        left = sum(c % 2 for c in cnt.values()) - odd_l  # S 中的剩余字母个数
        for w in words:
            if left <= 0: break
            left -= len(w) // 2 * 2
            ans -= 1
        return ans
```

```java [sol-Java]
class Solution {
    public int maxPalindromesAfterOperations(String[] words) {
        int oddL = 0; // 奇数长度字符串个数
        int mask = 0; // 奇数个数的字母集合
        for (String w : words) {
            oddL += w.length() % 2;
            for (char c : w.toCharArray()) {
                mask ^= 1 << (c - 'a');
            }
        }

        Arrays.sort(words, (a, b) -> b.length() - a.length());

        int ans = words.length;
        int left = Integer.bitCount(mask) - oddL; // S 中的剩余字母个数
        for (String w : words) {
            if (left <= 0) break;
            left -= w.length() / 2 * 2;
            ans--;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxPalindromesAfterOperations(vector<string> &words) {
        int oddL = 0; // 奇数长度字符串个数
        int mask = 0; // 奇数个数的字母集合
        for (auto &w : words) {
            oddL += w.length() % 2;
            for (char c : w) {
                mask ^= 1 << (c - 'a');
            }
        }

        ranges::sort(words, [](const auto &a, const auto &b) {
            return a.length() > b.length();
        });

        int ans = words.size();
        int left = __builtin_popcount(mask) - oddL; // S 中的剩余字母个数
        for (auto &w : words) {
            if (left <= 0) break;
            left -= w.length() / 2 * 2;
            ans--;
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxPalindromesAfterOperations(words []string) int {
	oddL, mask := 0, 0
	for _, w := range words {
		oddL += len(w) % 2 // 统计奇数长度字符串个数
		for _, c := range w {
			mask ^= 1 << (c - 'a')
		}
	}

	slices.SortFunc(words, func(a, b string) int { return len(b) - len(a) })

	ans := len(words)
	left := bits.OnesCount(uint(mask)) - oddL // S 中的剩余字母个数
	for _, w := range words {
		if left <= 0 { break }
		left -= len(w) / 2 * 2
		ans--
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L+n\log n)$，其中 $L$ 为字符串的长度之和，$n$ 为 $\textit{words}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。利用位运算可以做到 $\mathcal{O}(1)$ 额外空间（忽略排序的栈开销）。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
