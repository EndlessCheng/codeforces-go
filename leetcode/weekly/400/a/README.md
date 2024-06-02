把 E 看成上车，L 看成下车，问：车上最多有多少人？

模拟即可。

```py [sol-Python3]
class Solution:
    def minimumChairs(self, s: str) -> int:
        ans = cnt = 0
        for c in s:
            if c == 'E':
                cnt += 1
                ans = max(ans, cnt)
            else:
                cnt -= 1
        return ans
```

```java [sol-Java]
class Solution {
    public int minimumChairs(String s) {
        int ans = 0;
        int cnt = 0;
        for (char c : s.toCharArray()) {
            if (c == 'E') {
                ans = Math.max(ans, ++cnt);
            } else {
                cnt--;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumChairs(string s) {
        int ans = 0, cnt = 0;
        for (char c : s) {
            if (c == 'E') {
                ans = max(ans, ++cnt);
            } else {
                cnt--;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumChairs(s string) (ans int) {
	cnt := 0
	for _, c := range s {
		if c == 'E' {
			cnt++
			ans = max(ans, cnt)
		} else {
			cnt--
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{s}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
