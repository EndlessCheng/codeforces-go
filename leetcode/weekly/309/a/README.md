### 视频讲解

见[【周赛 309】](https://www.bilibili.com/video/BV1Dt4y1j7qh)

### 思路

遍历 $s$，记录 $s[i]$ 出现的位置 $\textit{last}[s[i]]$，如果再次遇到 $s[i]$，那么两个字母之间的字母个数即为 $i-\textit{last}[s[i]]-1$，与 $\textit{distance}$ 比较即可。

代码实现时，为了方便判断 $s[i]$ 是否第一次出现，可以记录 $\textit{last}[s[i]]=i+1$，这样 $\textit{last}[s[i]]$ 是正数即表示 $s[i]$ 出现过，此时两个字母之间的字母个数为 $i-\textit{last}[s[i]]$。

```py [sol1-Python3]
class Solution:
    def checkDistances(self, s: str, distance: List[int]) -> bool:
        last = [0] * 26
        for i, c in enumerate(s):
            c = ord(c) - ord('a')
            if last[c] and i - last[c] != distance[c]:
                return False
            last[c] = i + 1
        return True
```

```java [sol1-Java]
class Solution {
    public boolean checkDistances(String S, int[] distance) {
        var s = S.toCharArray();
        var last = new int[26];
        for (int i = 0; i < s.length; ++i) {
            int c = s[i] - 'a';
            if (last[c] != 0 && i - last[c] != distance[c])
                return false;
            last[c] = i + 1;
        }
        return true;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    bool checkDistances(string s, vector<int> &distance) {
        int last[26]{};
        for (int i = 0; i < s.size(); ++i) {
            int c = s[i] - 'a';
            if (last[c] && i - last[c] != distance[c])
                return false;
            last[c] = i + 1;
        }
        return true;
    }
};
```

```go [sol1-Go]
func checkDistances(s string, distance []int) bool {
	last := [26]int{}
	for i, c := range s {
		c -= 'a'
		if last[c] > 0 && i-last[c] != distance[c] {
			return false
		}
		last[c] = i + 1
	}
	return true
}
```

### 复杂度分析

- 时间复杂度：$O(|\Sigma|+n)$，其中 $|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$；$n$ 为 $s$ 的长度。
- 空间复杂度：$O(|\Sigma|)$。

---

欢迎关注[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)，高质量算法教学，持续更新中~
