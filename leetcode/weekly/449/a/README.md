统计每个字母的出现次数，记到一个长为 $26$ 的 $\textit{cnt}$ 数组中。

优先去掉出现次数少的字母，也就是去掉出现次数前 $26-k$ 小的字母。（出现次数为 $0$ 的也算上）

> 或者说，保留出现次数最大的 $k$ 个字母。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1h3EuzrEqV/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minDeletion(self, s: str, k: int) -> int:
        cnt = sorted(Counter(s).values())
        return sum(cnt[:-k])  # 保留出现次数最大的 k 个字母
```

```py [sol-Python3 list]
class Solution:
    def minDeletion(self, s: str, k: int) -> int:
        cnt = [0] * 26
        for b in s:
            cnt[ord(b) - ord('a')] += 1

        cnt.sort()
        return sum(cnt[:-k])
```

```java [sol-Java]
class Solution {
    public int minDeletion(String s, int k) {
        int[] cnt = new int[26];
        for (char b : s.toCharArray()) {
            cnt[b - 'a']++;
        }

        Arrays.sort(cnt);
        int ans = 0;
        for (int i = 0; i < 26 - k; i++) {
            ans += cnt[i];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minDeletion(string s, int k) {
        int cnt[26]{};
        for (char b : s) {
            cnt[b - 'a']++;
        }

        ranges::sort(cnt);
        int ans = 0;
        for (int i = 0; i < 26 - k; i++) {
            ans += cnt[i];
        }
        return ans;
    }
};
```

```go [sol-Go]
func minDeletion(s string, k int) (ans int) {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}

	slices.Sort(cnt[:])
	for _, c := range cnt[:26-k] {
		ans += c
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|\log |\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
