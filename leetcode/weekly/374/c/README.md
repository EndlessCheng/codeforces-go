「相邻字母相差至多为 $2$」这个约束把 $\textit{word}$ 划分成了多个子串 $s$，每个子串分别处理。可以用 [分组循环](https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/solution/jiao-ni-yi-ci-xing-ba-dai-ma-xie-dui-on-zuspx/) 找到每个子串 $s$。

对于每个子串，由于每个字符恰好出现 $k$ 次，我们可以枚举有 $m$ 种字符，这样问题就变成了：

- 长度固定为 $m\cdot k$ 的**滑动窗口**，判断每种字符是否都出现了恰好 $k$ 次。

[本题视频讲解](https://www.bilibili.com/video/BV1og4y1Z7SZ/?t=16m44s)

```py [sol-Python3]
class Solution:
    def countCompleteSubstrings(self, word: str, k: int) -> int:
        def f(s: str) -> int:
            res = 0
            for m in range(1, 27):
                if k * m > len(s):
                    break
                cnt = defaultdict(int)
                for right, c in enumerate(s):
                    cnt[c] += 1
                    left = right + 1 - k * m
                    if left >= 0:
                        res += all(c == 0 or c == k for c in cnt.values())
                        cnt[s[left]] -= 1
            return res

        n = len(word)
        ans = i = 0
        while i < n:
            st = i
            i += 1
            while i < n and abs(ord(word[i]) - ord(word[i - 1])) <= 2:
                i += 1
            ans += f(word[st:i])
        return ans
```

```java [sol-Java]
class Solution {
    public int countCompleteSubstrings(String word, int k) {
        int n = word.length();
        int ans = 0;
        for (int i = 0; i < n; ) {
            int st = i;
            for (i++; i < n && Math.abs(word.charAt(i) - word.charAt(i - 1)) <= 2; i++) ;
            ans += f(word.substring(st, i), k);
        }
        return ans;
    }

    private int f(String S, int k) {
        char[] s = S.toCharArray();
        int res = 0;
        for (int m = 1; m <= 26 && k * m <= s.length; m++) {
            int[] cnt = new int[26];
            for (int right = 0; right < s.length; right++) {
                cnt[s[right] - 'a']++;
                int left = right + 1 - k * m;
                if (left >= 0) {
                    boolean ok = true;
                    for (int i = 0; i < 26; i++) {
                        if (cnt[i] > 0 && cnt[i] != k) {
                            ok = false;
                            break;
                        }
                    }
                    if (ok) res++;
                    cnt[s[left] - 'a']--;
                }
            }
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    int f(string_view s, int k) {
        int res = 0;
        for (int m = 1; m <= 26 && k * m <= s.length(); m++) {
            int cnt[26]{};
            auto check = [&]() {
                for (int i = 0; i < 26; i++) {
                    if (cnt[i] && cnt[i] != k) {
                        return;
                    }
                }
                res++;
            };
            for (int right = 0; right < s.length(); right++) {
                cnt[s[right] - 'a']++;
                int left = right + 1 - k * m;
                if (left >= 0) {
                    check();
                    cnt[s[left] - 'a']--;
                }
            }
        }
        return res;
    }

public:
    int countCompleteSubstrings(string word, int k) {
        int n = word.length();
        int ans = 0;
        string_view s(word); // string_view 的 substr 没有拷贝
        for (int i = 0; i < n;) {
            int st = i;
            for (i++; i < n && abs(int(word[i]) - int(word[i - 1])) <= 2; i++);
            ans += f(s.substr(st, i - st), k);
        }
        return ans;
    }
};
```

```go [sol-Go]
func f(s string, k int) (res int) {
	for m := 1; m <= 26 && k*m <= len(s); m++ {
		cnt := [26]int{}
		check := func() {
			for i := range cnt {
				if cnt[i] > 0 && cnt[i] != k {
					return
				}
			}
			res++
		}
		for right, in := range s {
			cnt[in-'a']++
			if left := right + 1 - k*m; left >= 0 {
				check()
				cnt[s[left]-'a']--
			}
		}
	}
	return
}

func countCompleteSubstrings(word string, k int) (ans int) {
	for i, n := 0, len(word); i < n; {
		st := i
		for i++; i < n && abs(int(word[i])-int(word[i-1])) <= 2; i++ {
		}
		ans += f(word[st:i], k)
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|^2)$，其中 $n$ 是 $\textit{word}$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。忽略切片的开销。

**注**：可以进一步优化至 $\mathcal{O}(n|\Sigma|)$，具体见视频。

## 相似题目：枚举字母种类数+滑窗

- [395. 至少有 K 个重复字符的最长子串](https://leetcode.cn/problems/longest-substring-with-at-least-k-repeating-characters/)
- [1763. 最长的美好子字符串](https://leetcode.cn/problems/longest-nice-substring/)

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
