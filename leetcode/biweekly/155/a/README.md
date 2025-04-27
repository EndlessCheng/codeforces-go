把 $\textit{responses}[i]$ 去重后，用哈希表 $\textit{cnt}$ 统计每个字符串的出现次数。

统计的过程中，维护最大出现次数 $\textit{maxCnt}$ 以及字典序最小的答案。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def findCommonResponse(self, responses: List[List[str]]) -> str:
        ans, max_cnt = "", 0
        cnt = defaultdict(int)
        for resp in responses:
            for s in set(resp):
                cnt[s] += 1
                c = cnt[s]
                if c > max_cnt or c == max_cnt and s < ans:
                    max_cnt = c
                    ans = s
        return ans
```

```java [sol-Java]
class Solution {
    public String findCommonResponse(List<List<String>> responses) {
        String ans = "";
        int maxCnt = 0;
        Map<String, Integer> cnt = new HashMap<>();
        Set<String> vis = new HashSet<>();
        for (List<String> resp : responses) {
            vis.clear();
            for (String s : resp) {
                if (!vis.add(s)) { // s 在 vis 中
                    continue;
                }
                int c = cnt.merge(s, 1, Integer::sum); // c = ++cnt[s]
                if (c > maxCnt || c == maxCnt && s.compareTo(ans) < 0) {
                    maxCnt = c;
                    ans = s;
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string findCommonResponse(vector<vector<string>>& responses) {
        string ans;
        int max_cnt = 0;
        unordered_map<string, int> cnt;
        unordered_set<string> vis;
        for (auto& resp : responses) {
            vis.clear();
            for (auto& s : resp) {
                if (!vis.insert(s).second) { // s 在 vis 中
                    continue;
                }
                int c = ++cnt[s];
                if (c > max_cnt || c == max_cnt && s < ans) {
                    max_cnt = c;
                    ans = s;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func findCommonResponse(responses [][]string) (ans string) {
	maxCnt := 0
	cnt := map[string]int{}
	vis := map[string]struct{}{}
	for _, resp := range responses {
		clear(vis)
		for _, s := range resp {
			if _, ok := vis[s]; ok {
				continue
			}
			vis[s] = struct{}{}
			cnt[s]++
			c := cnt[s]
			if c > maxCnt || c == maxCnt && s < ans {
				maxCnt = c
				ans = s
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L)$，其中 $L$ 是所有 $\textit{responses}[i][j]$ 的长度之和。
- 空间复杂度：$\mathcal{O}(L)$。

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
