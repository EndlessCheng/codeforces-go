请看 [视频讲解](https://www.bilibili.com/video/BV1RH4y1W7DP/) 第二题。

设 $s$ 中有 $k$ 个 $c$。

每个 $c$ 都可以和它自己或左边的 $c$ 形成（对应着）一个符合要求的子串，所以一共有

$$
1 + 2 + \cdots + k = \dfrac{k(k+1)}{2}
$$

个符合要求的子串。

```py [sol-Python3]
class Solution:
    def countSubstrings(self, s: str, c: str) -> int:
        return comb(s.count(c) + 1, 2)
```

```java [sol-Java]
class Solution {
    public long countSubstrings(String s, char c) {
        long k = s.chars().filter(ch -> ch == c).count();
        return k * (k + 1) / 2;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countSubstrings(string s, char c) {
        long long k = ranges::count(s, c);
        return k * (k + 1) / 2;
    }
};
```

```go [sol-Go]
func countSubstrings(s string, c byte) int64 {
	k := int64(strings.Count(s, string(c)))
	return k * (k + 1) / 2
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)

更多题单，请点我个人主页 - 讨论发布。
