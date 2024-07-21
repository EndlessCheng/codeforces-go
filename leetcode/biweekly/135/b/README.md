操作次数取决于每种字母的出现次数，与字母的位置无关。

假设某个字母出现了 $c$ 次，那么操作后该字母最少能剩下多少？

根据题意，只有当 $c\ge 3$ 时才能操作，每次操作可以把 $c$ 减少 $2$。

- 如果 $c=3,5,7,\cdots$ 是奇数，那么不断减 $2$，最终 $c=1$。
- 如果 $c=4,6,8,\cdots$ 是偶数，那么不断减 $2$，最终 $c=2$。

这两种情况可以合并，最终剩下

$$
(c-1)\bmod 2 + 1
$$

个字母。注意上式同时兼顾 $c=0,1,2$ 的情况。

累加每种字母最终剩下的 $c$，即为答案。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1JE4m1d7br/) 第二题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def minimumLength(self, s: str) -> int:
        return sum((c - 1) % 2 + 1 for c in Counter(s).values())
```

```java [sol-Java]
class Solution {
    public int minimumLength(String s) {
        int[] cnt = new int[26];
        for (char b : s.toCharArray()) {
            cnt[b - 'a']++;
        }
        int ans = 0;
        for (int c : cnt) {
            ans += (c - 1) % 2 + 1;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumLength(string s) {
        int cnt[26]{};
        for (char b : s) {
            cnt[b - 'a']++;
        }
        int ans = 0;
        for (int c : cnt) {
            ans += (c - 1) % 2 + 1;
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumLength(s string) (ans int) {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}
	for _, c := range cnt {
		ans += (c-1)%2 + 1
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|$ 是字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
