由于字典序要严格大于 $\textit{target}$，我们倒着枚举，看看能否把 $j = \textit{target}[i]$ 增大：

这要求：

- $s$ 的 $[0,i-1]$ 中的字母和 $\textit{target}$ 的这一段是一样的，消耗掉。
- 在 $[\textit{target}[i]+1,\texttt{z}]$ 中，$s$ 存在剩余可以用的字母。
  - 如果存在，那么可以把 $\textit{target}[i]$ 增大。
  - 剩余字母按照从小到大的顺序排在后面。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def lexGreaterPermutation(self, s: str, target: str) -> str:
        left = Counter(s)
        for c in target:
            left[c] -= 1

        ans = list(target)
        # 从右往左尝试
        for i in range(len(s) - 1, -1, -1):
            b = target[i]
            left[b] += 1
            if any(c < 0 for c in left.values()):
                continue  # 前面无法做到全部一样

            # target[i] 增大到 j
            for j in range(ord(b) - ord('a') + 1, 26):
                ch = ascii_lowercase[j]
                if left[ch] == 0:
                    continue

                left[ch] -= 1
                ans[i] = ch
                del ans[i + 1:]

                for ch in ascii_lowercase:
                    ans.extend(ch * left[ch])
                return ''.join(ans)
            # 增大失败，继续枚举
        return ""
```

```java [sol-Java]
class Solution {
    public String lexGreaterPermutation(String s, String target) {
        char[] t = target.toCharArray();
        int n = t.length;
        int[] left = new int[26];
        for (int i = 0; i < n; i++) {
            left[s.charAt(i) - 'a']++;
            left[t[i] - 'a']--;
        }
        StringBuilder ans = new StringBuilder(target);

        // 从右往左尝试
        next:
        for (int i = n - 1; i >= 0; i--) {
            int b = t[i] - 'a';
            left[b]++;
            for (int c : left) {
                if (c < 0) { // 前面无法做到全部一样
                    continue next;
                }
            }

            // target[i] 增大到 j
            for (int j = b + 1; j < 26; j++) {
                if (left[j] == 0) {
                    continue;
                }

                left[j]--;
                ans.setCharAt(i, (char) ('a' + j));
                ans.setLength(i + 1);

                for (int k = 0; k < 26; k++) {
                    ans.repeat('a' + k, left[k]);
                }
                return ans.toString();
            }
        }
        return "";
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string lexGreaterPermutation(string s, string target) {
        int left[26]{};
        for (int i = 0; i < s.size(); i++) {
            left[s[i] - 'a']++;
            left[target[i] - 'a']--;
        }

        // 从右往左尝试
        for (int i = s.size() - 1; i >= 0; i--) {
            int b = target[i] - 'a';
            left[b]++;

            bool ok = true;
            for (int c : left) {
                if (c < 0) { // 前面无法做到全部一样
                    ok = false;
                    break;
                }
            }
            if (!ok) {
                continue;
            }

            // target[i] 增大到 j
            for (int j = b + 1; j < 26; j++) {
                if (left[j] == 0) {
                    continue;
                }

                left[j]--;
                target[i] = 'a' + j;
                target.resize(i + 1);

                for (int k = 0; k < 26; k++) {
                    target += string(left[k], 'a' + k);
                }
                return target;
            }
            // 增大失败，继续枚举
        }
        return "";
    }
};
```

```go [sol-Go]
func lexGreaterPermutation(s, target string) string {
	left := make([]int, 26)
	for i, b := range s {
		left[b-'a']++
		left[target[i]-'a']--
	}
	ans := []byte(target)

next:
	for i := len(s) - 1; i >= 0; i-- {
		b := target[i] - 'a'
		left[b]++
		for _, c := range left {
			if c < 0 { // 前面无法做到全部一样
				continue next
			}
		}

		// target[i] 增大到 j
		for j := b + 1; j < 26; j++ {
			if left[j] == 0 {
				continue
			}

			left[j]--
			ans[i] = 'a' + j
			ans = ans[:i+1]

			for k, c := range left {
				ch := string('a' + byte(k))
				ans = append(ans, strings.Repeat(ch, c)...)
			}
			return string(ans)
		}
		// 增大失败，继续枚举
	}
	return ""
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|)$，其中 $n$ 是 $\textit{nums}$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

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
