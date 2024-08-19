遍历 $s$，用哈希表（或者数组）记录字母 $c=s[i]$ 的下标 $\textit{pos}[c] = i$。注意 $s$ 和 $t$ 中没有重复字母。

然后遍历 $t$，累加 $|i-pos[t[i]]|$，即为答案。

请看 [视频讲解](https://www.bilibili.com/video/BV1bx4y1i7rP/)，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def findPermutationDifference(self, s: str, t: str) -> int:
        pos = {c: i for i, c in enumerate(s)}
        return sum(abs(i - pos[c]) for i, c in enumerate(t))
```

```java [sol-Java]
class Solution {
    public int findPermutationDifference(String s, String t) {
        int[] pos = new int[26];
        for (int i = 0; i < s.length(); i++) {
            pos[s.charAt(i) - 'a'] = i;
        }
        int ans = 0;
        for (int i = 0; i < t.length(); i++) {
            ans += Math.abs(i - pos[t.charAt(i) - 'a']);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findPermutationDifference(string s, string t) {
        int pos[26];
        for (int i = 0; i < s.length(); i++) {
            pos[s[i] - 'a'] = i;
        }
        int ans = 0;
        for (int i = 0; i < t.length(); i++) {
            ans += abs(i - pos[t[i] - 'a']);
        }
        return ans;
    }
};
```

```go [sol-Go]
func findPermutationDifference(s, t string) (ans int) {
	pos := [26]int{}
	for i, c := range s {
		pos[c-'a'] = i
	}
	for i, c := range t {
		ans += abs(i - pos[c-'a'])
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

分析哈希表写法（Python 代码）的复杂度：

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
