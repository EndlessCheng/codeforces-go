用哈希表 $\textit{color}$ 维护第 $x$ 个球的颜色，哈希表 $\textit{cnt}$ 维护每种颜色的出现次数。

当把球 $x$ 的颜色改成 $y$ 时：

1. 如果 $x$ 在 $\textit{color}$ 中，设 $c=\textit{color}[x]$，先把 $\textit{cnt}[c]$ 减少一，如果 $\textit{cnt}[c]$ 变成 $0$，则从 $\textit{cnt}$ 中删除 $c$。
2. 更新 $\textit{color}[x] = y$。
3. 把 $\textit{cnt}[y]$ 增加一。
4. 把 $\textit{cnt}$ 的大小加入答案。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1SU411d7wj/) 第三题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def queryResults(self, _: int, queries: List[List[int]]) -> List[int]:
        ans = []
        color = {}
        cnt = defaultdict(int)
        for x, y in queries:
            if x in color:
                c = color[x]
                cnt[c] -= 1
                if cnt[c] == 0:
                    del cnt[c]
            color[x] = y
            cnt[y] += 1
            ans.append(len(cnt))
        return ans
```

```java [sol-Java]
class Solution {
    public int[] queryResults(int limit, int[][] queries) {
        int[] ans = new int[queries.length];
        Map<Integer, Integer> color = new HashMap<>();
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int i = 0; i < queries.length; i++) {
            int[] q = queries[i];
            int x = q[0];
            int y = q[1];
            if (color.containsKey(x)) {
                int c = color.get(x);
                if (cnt.merge(c, -1, Integer::sum) == 0) {
                    cnt.remove(c);
                }
            }
            color.put(x, y);
            cnt.merge(y, 1, Integer::sum);
            ans[i] = cnt.size();
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> queryResults(int, vector<vector<int>>& queries) {
        vector<int> ans;
        unordered_map<int, int> color, cnt;
        for (auto& q : queries) {
            int x = q[0], y = q[1];
            if (auto it = color.find(x); it != color.end()) {
                int c = it->second;
                if (--cnt[c] == 0) {
                    cnt.erase(c);
                }
            }
            color[x] = y;
            cnt[y]++;
            ans.push_back(cnt.size());
        }
        return ans;
    }
};
```

```go [sol-Go]
func queryResults(_ int, queries [][]int) []int {
	ans := make([]int, len(queries))
	color := map[int]int{}
	cnt := map[int]int{}
	for i, q := range queries {
		x, y := q[0], q[1]
		if c := color[x]; c > 0 {
			cnt[c]--
			if cnt[c] == 0 {
				delete(cnt, c)
			}
		}
		color[x] = y
		cnt[y]++
		ans[i] = len(cnt)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(q)$，其中 $q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(q)$。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
