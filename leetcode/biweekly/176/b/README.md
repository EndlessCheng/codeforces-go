题意有点绕，说人话就是：

- 统计 $\textit{words}[i][..k]$（长为 $k$ 的前缀）的出现次数，有多少个出现次数 $\ge 2$ 的**不同**前缀？

用哈希表统计即可。

跳过长度小于 $k$ 的字符串。

[本题视频讲解](https://www.bilibili.com/video/BV15TZ4B1Eev/?t=2m48s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def prefixConnected(self, words: List[str], k: int) -> int:
        cnt = Counter(w[:k] for w in words if len(w) >= k)
        return sum(c > 1 for c in cnt.values())
```

```java [sol-Java]
class Solution {
    public int prefixConnected(String[] words, int k) {
        Map<String, Integer> cnt = new HashMap<>();
        for (String w : words) {
            if (w.length() >= k) {
                cnt.merge(w.substring(0, k), 1, Integer::sum);
            }
        }

        int ans = 0;
        for (int c : cnt.values()) {
            if (c > 1) {
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int prefixConnected(vector<string>& words, int k) {
        unordered_map<string, int> cnt;
        for (auto& w : words) {
            if (w.size() >= k) {
                cnt[w.substr(0, k)]++;
            }
        }

        int ans = 0;
        for (auto& [_, c] : cnt) {
            ans += c > 1;
        }
        return ans;
    }
};
```

```go [sol-Go]
func prefixConnected(words []string, k int) (ans int) {
	cnt := map[string]int{}
	for _, w := range words {
		if len(w) >= k {
			cnt[w[:k]]++
		}
	}

	for _, c := range cnt {
		if c > 1 {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk)$，其中 $n$ 是 $\textit{words}$ 的长度。
- 空间复杂度：$\mathcal{O}(nk)$。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
