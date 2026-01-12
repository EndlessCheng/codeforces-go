比如字符串 $\texttt{bcd}$ 和 $\texttt{cde}$。把字符串通过操作，改成第一个字母是 $\texttt{a}$，就能看出这两个字符串都是 $\texttt{abc}$ 通过循环替换得到的。

用「枚举右，维护左」的技巧：

- 创建一个哈希表 $\textit{cnt}$，统计每个字符串的出现次数。
- 设 $\textit{words}[j]$ 通过循环替换，把第一个字母变成 $\texttt{a}$ 后的字符串为 $t$。
- 先从 $\textit{cnt}$ 中查询 $t$ 的出现次数，即为 $i$ 的个数，加到答案中。
- 然后把 $\textit{cnt}[t]$ 加一。

可以用折线图的平移和重叠理解这个思路，见 [本题视频讲解](https://www.bilibili.com/video/BV1tv6dBME7K/?t=9m36s)。

> 注：也可以计算 $\textit{words}[j]$ 的长为 $m-1$ 的差分数组，哈希表统计差分数组的出现次数。

## 优化前

```py [sol-Python3]
class Solution:
    def countPairs(self, words: List[str]) -> int:
        cnt = defaultdict(int)
        ans = 0
        for s in words:
            t = list(s)
            base = ord(t[0])
            for i in range(len(t)):
                t[i] = chr((ord(t[i]) - base) % 26)  # 保证结果在 [0, 25] 中
            t = ''.join(t)
            ans += cnt[t]
            cnt[t] += 1
        return ans
```

```java [sol-Java]
class Solution {
    public long countPairs(String[] words) {
        Map<String, Integer> cnt = new HashMap<>();
        long ans = 0;
        for (String s : words) {
            char[] t = s.toCharArray();
            char base = t[0];
            for (int i = 0; i < t.length; i++) {
                t[i] = (char) ((t[i] - base + 26) % 26); // 保证结果在 [0, 25] 中
            }
            s = new String(t);
            int c = cnt.getOrDefault(s, 0);
            ans += c;
            cnt.put(s, c + 1);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countPairs(vector<string>& words) {
        unordered_map<string, int> cnt;
        long long ans = 0;
        for (auto& s : words) {
            char base = s[0];
            for (char& ch : s) {
                ch = (ch - base + 26) % 26; // 保证结果在 [0, 25] 中
            }
            ans += cnt[s];
            cnt[s]++;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countPairs(words []string) (ans int64) {
	cnt := map[string]int{}
	for _, s := range words {
		t := []byte(s)
		base := t[0]
		for i := range t {
			t[i] = (t[i] - base + 26) % 26 // 保证结果在 [0, 25] 中
		}
		s = string(t)
		ans += int64(cnt[s])
		cnt[s]++
	}
	return
}
```

## 优化

统计 $\textit{words}$ 每个字符串的出现次数，只对不同的字符串计算 $t$，从而减少计算 $t$ 的次数。

注意答案与 $\textit{words}$ 中的元素顺序无关，因为本质是从 $\textit{words}$ 中选两个相似字符串的方案数。无论 $\textit{words}=[A,B]$ 还是 $[B,A]$，答案是一样的。

```py [sol-Python3]
class Solution:
    def countPairs(self, words: List[str]) -> int:
        cnt = defaultdict(int)
        ans = 0
        for s, c in Counter(words).items():
            t = list(s)
            base = ord(t[0])
            for i in range(len(t)):
                t[i] = chr((ord(t[i]) - base) % 26)  # 保证结果在 [0, 25] 中
            t = ''.join(t)
            ans += cnt[t] * c + c * (c - 1) // 2  # c 个 s 中选 2 个有 C(c, 2) 种方案
            cnt[t] += c
        return ans
```

```java [sol-Java]
class Solution {
    public long countPairs(String[] words) {
        Map<String, Integer> cntWords = new HashMap<>();
        for (String s : words) {
            cntWords.merge(s, 1, Integer::sum);
        }

        Map<String, Integer> cnt = new HashMap<>();
        long ans = 0;
        for (Map.Entry<String, Integer> e : cntWords.entrySet()) {
            String s = e.getKey();
            int totalS = e.getValue();
            char[] t = s.toCharArray();
            char base = t[0];
            for (int i = 0; i < t.length; i++) {
                t[i] = (char) ((t[i] - base + 26) % 26); // 保证结果在 [0, 25] 中
            }
            s = new String(t);
            int c = cnt.getOrDefault(s, 0);
            // totalS 个 s 中选 2 个有 C(totalS, 2) 种方案
            ans += (long) c * totalS + (long) totalS * (totalS - 1) / 2;
            cnt.put(s, c + totalS);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countPairs(vector<string>& words) {
        unordered_map<string, int> cnt_words;
        for (auto& s : words) {
            cnt_words[s]++;
        }

        unordered_map<string, int> cnt;
        long long ans = 0;
        for (auto& [s, c] : cnt_words) {
            string t = move(s);
            char base = t[0];
            for (char& ch : t) {
                ch = (ch - base + 26) % 26; // 保证结果在 [0, 25] 中
            }
            // c 个 s 中选 2 个有 C(c, 2) 种方案
            ans += 1LL * cnt[t] * c + 1LL * c * (c - 1) / 2;
            cnt[t] += c;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countPairs(words []string) int64 {
	cntWords := map[string]int{}
	for _, s := range words {
		cntWords[s]++
	}

	ans := 0
	cnt := map[string]int{}
	for s, c := range cntWords {
		t := []byte(s)
		base := t[0]
		for i := range t {
			t[i] = (t[i] - base + 26) % 26 // 保证结果在 [0, 25] 中
		}
		s = string(t)
		ans += cnt[s]*c + c*(c-1)/2 // c 个 s 中选 2 个有 C(c, 2) 种方案
		cnt[s] += c
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 是 $\textit{words}$ 的长度，$m$ 是每个 $\textit{words}[i]$ 的长度。
- 空间复杂度：$\mathcal{O}(nm)$。

## 专题训练

见下面数据结构题单的「**§0.1 枚举右，维护左**」。

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
