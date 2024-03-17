请看 [视频讲解](https://www.bilibili.com/video/BV1RH4y1W7DP/)。

用一个 $26\times 26$ 的布尔数组（或哈希表）$\textit{vis}$ 记录是否遇到了 $\textit{vis}[x][y]$，其中 $x$ 和 $y$ 是一对相邻字母 $(s[i-1],s[i])$。

如果 $\textit{vis}[y][x]$ 为真，则说明 $x+y$ 在反转后的字符串中。

```py [sol-Python3]
class Solution:
    def isSubstringPresent(self, s: str) -> bool:
        st = set()
        for x, y in pairwise(s):
            st.add((x, y))
            if (y, x) in st:
                return True
        return False
```

```java [sol-Java]
class Solution {
    public boolean isSubstringPresent(String S) {
        char[] s = S.toCharArray();
        boolean[][] vis = new boolean[26][26];
        for (int i = 1; i < s.length; i++) {
            int x = s[i - 1] - 'a';
            int y = s[i] - 'a';
            vis[x][y] = true;
            if (vis[y][x]) {
                return true;
            }
        }
        return false;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool isSubstringPresent(string s) {
        bool vis[26][26]{};
        for (int i = 1; i < s.length(); i++) {
            int x = s[i - 1] - 'a', y = s[i] - 'a';
            vis[x][y] = true;
            if (vis[y][x]) {
                return true;
            }
        }
        return false;
    }
};
```

```go [sol-Go]
func isSubstringPresent(s string) bool {
	vis := [26][26]bool{}
	for i := 1; i < len(s); i++ {
		x, y := s[i-1]-'a', s[i]-'a'
		vis[x][y] = true
		if vis[y][x] {
			return true
		}
	}
	return false
}
```

也可以用位运算优化，把 $\textit{vis}$ 数组的第二维压缩成二进制数。原理见 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

```py [sol-Python3]
class Solution:
    def isSubstringPresent(self, s: str) -> bool:
        vis = [0] * 26
        for x, y in pairwise(map(ord, s)):
            x -= ord('a')
            y -= ord('a')
            vis[x] |= 1 << y
            if vis[y] >> x & 1:
                return True
        return False
```

```java [sol-Java]
class Solution {
    public boolean isSubstringPresent(String S) {
        char[] s = S.toCharArray();
        int[] vis = new int[26];
        for (int i = 1; i < s.length; i++) {
            int x = s[i - 1] - 'a';
            int y = s[i] - 'a';
            vis[x] |= 1 << y;
            if ((vis[y] >> x & 1) > 0) {
                return true;
            }
        }
        return false;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool isSubstringPresent(string s) {
        int vis[26]{};
        for (int i = 1; i < s.length(); i++) {
            int x = s[i - 1] - 'a', y = s[i] - 'a';
            vis[x] |= 1 << y;
            if (vis[y] >> x & 1) {
                return true;
            }
        }
        return false;
    }
};
```

```go [sol-Go]
func isSubstringPresent(s string) bool {
	vis := [26]int{}
	for i := 1; i < len(s); i++ {
		x, y := s[i-1]-'a', s[i]-'a'
		vis[x] |= 1 << y
		if vis[y]>>x&1 > 0 {
			return true
		}
	}
	return false
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|)$，其中 $n$ 为 $\textit{nums}$ 的长度，$|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)

更多题单，请点我个人主页 - 讨论发布。
