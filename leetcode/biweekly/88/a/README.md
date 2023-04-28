## 方法一：暴力枚举删除的字符

先简单说说暴力的做法。

枚举 $\textit{word}[i]$，将其去掉后再统计剩余字符的出现次数。用数组或者哈希表统计都行。

如果剩余字符的出现次数都相同，则返回 `true`。

反之，如果无论去掉哪个 $\textit{word}[i]$，都无法让剩余字符的出现次数都相同，则返回 `false`。

```py [sol1-Python3]
class Solution:
    def equalFrequency(self, word: str) -> bool:
        for i in range(len(word)):  # 枚举删除的字符
            cnt = Counter(word[:i] + word[i + 1:])  # 统计出现次数
            if len(set(cnt.values())) == 1:  # 出现次数都一样
                return True
        return False
```

```java [sol1-Java]
class Solution {
    public boolean equalFrequency(String word) {
        var s = word.toCharArray();
        int n = s.length;
        for (int i = 0; i < n; ++i) { // 枚举删除的字符
            var cnt = new HashMap<Character, Integer>();
            for (int j = 0; j < n; ++j)
                if (j != i)
                    cnt.merge(s[j], 1, Integer::sum); // 统计出现次数
            if (isSame(cnt)) // 出现次数都一样
                return true;
        }
        return false;
    }

    private boolean isSame(Map<Character, Integer> cnt) {
        int c0 = cnt.entrySet().iterator().next().getValue();
        for (int c : cnt.values())
            if (c != c0)
                return false;
        return true;
    }
}
```

```cpp [sol1-C++]
class Solution {
    bool is_same(unordered_map<char, int> &cnt) {
        int c0 = cnt.begin()->second;
        for (auto &[_, c]: cnt)
            if (c != c0)
                return false;
        return true;
    }
public:
    bool equalFrequency(string word) {
        int n = word.length();
        for (int i = 0; i < n; ++i) { // 枚举删除的字符
            unordered_map<char, int> cnt;
            for (int j = 0; j < n; ++j)
                if (j != i)
                    ++cnt[word[j]]; // 统计出现次数
            if (is_same(cnt)) // 出现次数都一样
                return true;
        }
        return false;
    }
};
```

```go [sol1-Go]
func equalFrequency(word string) bool {
next:
	for i := range word { // 枚举删除的字符
		cnt := map[rune]int{}
		for j, c := range word {
			if j != i {
				cnt[c]++ // 统计出现次数
			}
		}
		c0 := 0
		for _, c := range cnt {
			if c0 == 0 {
				c0 = c
			} else if c != c0 { // 出现次数不一样
				continue next // 枚举下一个字符
			}
		}
		return true // 循环没有中途退出，说明出现次数都一样
	}
	return false
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{word}$ 长度。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$，其中 $|\Sigma|$ 为字符集的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。

## 方法二：分类讨论

如果有至少三种**不同**的出现次数，例如 $\texttt{"abccdddd"}$，出现次数分别为 $1,1,2,4$，无论去掉哪个字符，剩下的出现次数仍然至少有两种。例如去掉 $\texttt{c}$ 得到 $\texttt{"abcdddd"}$，出现次数分别为 $1,1,1,4$，不满足题目要求。

所以**只需要讨论出现次数至多两种**的情况：

- 如果出现次数只有一种：
   - 如果只有一种字符，例如 $\texttt{"aaaaa"}$，那么无论删除哪个都是满足要求的。
   - 如果每个字符恰好出现一次，例如 $\texttt{"abcde"}$，那么无论删除哪个都是满足要求的。代码实现时，这可以合并到下面的「较少的出现次数恰好是 $1$」的情况中。
   - 如果每个字符出现不止一次，例如 $\texttt{"aabbcc"}$，虽然出现次数均为 $2$，但是题目要求**恰好**去掉一个字符，所以无法满足要求。
- 如果出现次数有两种，那么必须变成一种：
   - 考虑去掉出现次数较少的字符：它的出现次数必须恰好是 $1$，且只有这一种字符出现一次。例如 $\texttt{"abbbccc"}$，去掉只出现一次的字符，就能满足要求。但如果它的出现次数大于 $1$，例如 $\texttt{"aabbbccc"}$，就无法满足要求。
   - 考虑去掉出现次数较多的字符：它的出现次数必须比出现次数较少的恰好多 $1$，且只有这一种字符出现一次。例如 $\texttt{"aabbccc"}$ 去掉 $\texttt{c}$ 是可以满足要求的。其它的情况例如 $\texttt{"abccc"}$ 或 $\texttt{"abccdd"}$ 都是无法满足要求的。

```py [sol2-Python3]
class Solution:
    def equalFrequency(self, word: str) -> bool:
        cnt = sorted(Counter(word).values())  # 出现次数从小到大排序
        # 只有一种字符 or 去掉次数最少的 or 去掉次数最多的
        return len(cnt) == 1 or \
               cnt[0] == 1 and len(set(cnt[1:])) == 1 or \
               cnt[-1] == cnt[-2] + 1 and len(set(cnt[:-1])) == 1
```

```java [sol2-Java]
class Solution {
    public boolean equalFrequency(String word) {
        var mCnt = new HashMap<Character, Integer>();
        for (var c : word.toCharArray())
            mCnt.merge(c, 1, Integer::sum);
        var cnt = new ArrayList<>(mCnt.values());
        Collections.sort(cnt); // 出现次数从小到大排序
        int m = cnt.size();
        // 只有一种字符 or 去掉次数最少的 or 去掉次数最多的
        return m == 1 ||
               cnt.get(0) == 1 && isSame(cnt.subList(1, m)) ||
               cnt.get(m - 1) == cnt.get(m - 2) + 1 && isSame(cnt.subList(0, m - 1));
    }

    private boolean isSame(List<Integer> cnt) {
        int c0 = cnt.get(0);
        for (int c : cnt)
            if (c != c0)
                return false;
        return true;
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    bool equalFrequency(string word) {
        unordered_map<char, int> m_cnt;
        for (char c: word)
            ++m_cnt[c];
        vector<int> cnt;
        for (auto &[_, c]: m_cnt)
            cnt.push_back(c);
        sort(cnt.begin(), cnt.end()); // 出现次数从小到大排序
        // 只有一种字符 or 去掉次数最少的 or 去掉次数最多的
        return cnt.size() == 1 ||
               cnt[0] == 1 && equal(cnt.begin() + 2, cnt.end(), cnt.begin() + 1) ||
               cnt.back() == cnt[cnt.size() - 2] + 1 && equal(cnt.begin() + 1, cnt.end() - 1, cnt.begin());
    }
};
```

```go [sol2-Go]
func equalFrequency(word string) bool {
	mCnt := map[rune]int{}
	for _, c := range word {
		mCnt[c]++
	}
	cnt := make([]int, 0, len(mCnt))
	for _, c := range mCnt {
		cnt = append(cnt, c)
	}
	sort.Ints(cnt) // 出现次数从小到大排序
	m := len(cnt)
	// 只有一种字符 or 去掉次数最少的 or 去掉次数最多的
	return m == 1 ||
		   cnt[0] == 1 && isSame(cnt[1:]) ||
		   cnt[m-1] == cnt[m-2]+1 && isSame(cnt[:m-1])
}

func isSame(a []int) bool {
	for _, x := range a[1:] {
		if x != a[0] {
			return false
		}
	}
	return true
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|\log |\Sigma|)$，其中 $n$ 为 $\textit{word}$ 长度，$|\Sigma|$ 为字符集的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

### 相似题目

- [1224. 最大相等频率](https://leetcode.cn/problems/maximum-equal-frequency/)

[往期每日一题题解](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

---

欢迎关注[ biIibiIi@灵茶山艾府](https://space.bilibili.com/206214)，高质量算法教学，持续输出中~
