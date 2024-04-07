本题相当于：

- 每次操作可以把 $s[i]$ 加一或减一，求在操作次数不超过 $k$ 的前提下，$s$ 的最小字典序。

算法：

1. 从左到右遍历 $s$。
2. 如果把 $s[i]$ 变成 $\texttt{a}$ 的操作次数 $\textit{dis} \le k$，那么就把 $s[i]$ 变成 $\texttt{a}$，同时 $k$ 减少 $\textit{dis}$。
3. 否则，把 $s[i]$ 减少 $k$，退出循环。

请看 [视频讲解](https://www.bilibili.com/video/BV1ut421H7Wv/) 第二题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def getSmallestString(self, s: str, k: int) -> str:
        s = list(s)
        for i, c in enumerate(map(ord, s)):
            dis = min(c - ord('a'), ord('z') - c + 1)
            if dis > k:
                s[i] = chr(c - k)
                break
            s[i] = 'a'
            k -= dis
        return ''.join(s)
```

```java [sol-Java]
class Solution {
    public String getSmallestString(String s, int k) {
        char[] t = s.toCharArray();
        for (int i = 0; i < t.length; i++) {
            int dis = Math.min(t[i] - 'a', 'z' - t[i] + 1);
            if (dis > k) {
                t[i] -= k;
                break;
            }
            t[i] = 'a';
            k -= dis;
        }
        return new String(t);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string getSmallestString(string s, int k) {
        for (int i = 0; i < s.length(); i++) {
            int dis = min(s[i] - 'a', 'z' - s[i] + 1);
            if (dis > k) {
                s[i] -= k;
                break;
            }
            s[i] = 'a';
            k -= dis;
        }
        return s;
    }
};
```

```go [sol-Go]
func getSmallestString(s string, k int) string {
	t := []byte(s)
	for i, c := range t {
		dis := int(min(c-'a', 'z'-c+1))
		if dis > k {
			t[i] -= byte(k)
			break
		}
		t[i] = 'a'
		k -= dis
	}
	return string(t)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$，其中 C++ 可以原地修改字符串。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
