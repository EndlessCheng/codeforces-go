[视频讲解](https://www.bilibili.com/video/BV1KG4y1j73o/?t=8m51s) 已出炉，欢迎点赞三连~

---

统计 $\textit{word}_1$ 字符出现次数 $c_1$ 以及 $\textit{word}_2$ 字符出现次数 $c_2$：

- 如果 $x=y$，那么移动后不同字符数目不变，如果此时 $c_1$ 和 $c_2$ 的长度相同，那么返回 true；
- 如果 $x\ne y$，那么就看 $x$ 的个数是否为 $1$，$y$ 的个数是否为 $1$，$x$ 是否出现在 $\textit{word}_2$ 中，$y$ 是否出现在 $\textit{word}_1$ 中来计算不同字符的变化量：
    - 如果 $x$ 个数为 $1$，那么移动后 $c_1$ 的大小减少了 $1$；
    - 如果 $y$ 个数为 $1$，那么移动后 $c_2$ 的大小减少了 $1$；
    - 如果 $y$ 不在 $\textit{word}_1$ 中，那么移动后 $c_1$ 的大小增加了 $1$；
    - 如果 $x$ 不在 $\textit{word}_2$ 中，那么移动后 $c_2$ 的大小增加了 $1$。

```py [sol1-Python3]
class Solution:
    def isItPossible(self, word1: str, word2: str) -> bool:
        c1, c2 = Counter(word1), Counter(word2)
        for x, c in c1.items():
            for y, d in c2.items():
                if y == x:  # 无变化
                    if len(c1) == len(c2): return True
                elif len(c1) - (c == 1) + (y not in c1) == \
                     len(c2) - (d == 1) + (x not in c2):  # 基于长度计算变化量
                    return True
        return False
```

```java [sol1-Java]
class Solution {
    public boolean isItPossible(String word1, String word2) {
        Map<Character, Integer> c1 = new HashMap<>(), c2 = new HashMap<>();
        for (char c : word1.toCharArray()) c1.merge(c, 1, Integer::sum);
        for (char c : word2.toCharArray()) c2.merge(c, 1, Integer::sum);
        for (var e : c1.entrySet())
            for (var f : c2.entrySet()) {
                char x = e.getKey(), y = f.getKey();
                if (x == y) {
                    if (c1.size() == c2.size()) return true;
                } else if (c1.size() - (e.getValue() == 1 ? 1 : 0) + (c1.containsKey(y) ? 0 : 1) ==
                           c2.size() - (f.getValue() == 1 ? 1 : 0) + (c2.containsKey(x) ? 0 : 1)) // 基于长度计算变化量
                    return true;
            }
        return false;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    bool isItPossible(string word1, string word2) {
        unordered_map<char, int> c1, c2;
        for (char c: word1) ++c1[c];
        for (char c: word2) ++c2[c];
        for (auto[x, c]: c1)
            for (auto[y, d]: c2)
                if (y == x) {
                    if (c1.size() == c2.size()) return true;
                } else if (c1.size() - (c == 1) + !c1.count(y) ==
                           c2.size() - (d == 1) + !c2.count(x)) // 基于长度计算变化量
                    return true;
        return false;
    }
};
```

```go [sol1-Go]
func isItPossible(word1, word2 string) bool {
	c1 := map[rune]int{}
	for _, c := range word1 {
		c1[c]++
	}
	c2 := map[rune]int{}
	for _, c := range word2 {
		c2[c]++
	}
	for x, c := range c1 {
		for y, d := range c2 {
			if y == x { // 无变化
				if len(c1) == len(c2) {
					return true
				}
			} else if len(c1)-b2i(c == 1)+b2i(c1[y] == 0) ==
				      len(c2)-b2i(d == 1)+b2i(c2[x] == 0) { // 基于长度计算变化量
				return true
			}
		}
	}
	return false
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}
```

#### 复杂度分析

- 时间复杂度：$O(n+m+|\Sigma|^2)$，其中 $n$ 为 $\textit{word}_1$ 的长度，$m$ 为 $\textit{word}_2$ 的长度，$|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$O(|\Sigma|)$。
