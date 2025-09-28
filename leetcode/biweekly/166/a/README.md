1. 统计每种字母的出现次数，记录在一个哈希表（或者数组）$\textit{cnt}$ 中。
2. 遍历 $\textit{cnt}$，把出现次数相同的字母，分到同一组，这可以用一个哈希表套列表实现。这一步相当于反向映射，把 $\textit{cnt}$ 的 value 映射到 key 上。
3. 答案为包含字母最多的组。如果有多个组包含的字母都是最多的，那么返回其中出现次数最大的组。

[本题视频讲解](https://www.bilibili.com/video/BV1AKnRz8Ejn/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def majorityFrequencyGroup(self, s: str) -> str:
        cnt = Counter(s)
        groups = defaultdict(list)
        k = 0
        for ch, cnt in cnt.items():
            groups[cnt].append(ch)
            if (len(groups[cnt]), cnt) > (len(groups[k]), k):
                k = cnt
        return ''.join(groups[k])
```

```java [sol-Java]
class Solution {
    public String majorityFrequencyGroup(String s) {
        int[] cnt = new int[26];
        for (char b : s.toCharArray()) {
            cnt[b - 'a']++;
        }

        Map<Integer, StringBuilder> groups = new HashMap<>();
        int mx = 0;
        for (int i = 0; i < 26; i++) {
            int c = cnt[i];
            if (c == 0) {
                continue;
            }
            groups.computeIfAbsent(c, _ -> new StringBuilder()).append((char) ('a' + i));
            if (mx == 0 ||
                    groups.get(c).length() > groups.get(mx).length() ||
                    groups.get(c).length() == groups.get(mx).length() && c > mx) {
                mx = c;
            }
        }

        return groups.get(mx).toString();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string majorityFrequencyGroup(string s) {
        int cnt[26]{};
        for (char b : s) {
            cnt[b - 'a']++;
        }

        unordered_map<int, string> groups;
        int mx = 0;
        for (int i = 0; i < 26; i++) {
            int c = cnt[i];
            if (c == 0) {
                continue;
            }
            groups[c] += 'a' + i;
            if (pair(groups[c].size(), c) > pair(groups[mx].size(), mx)) {
                mx = c;
            }
        }

        return groups[mx];
    }
};
```

```go [sol-Go]
func majorityFrequencyGroup(s string) string {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}

	groups := map[int][]byte{}
	mx := 0
	for i, c := range cnt {
		if c == 0 {
			continue
		}
		groups[c] = append(groups[c], 'a'+byte(i))
		if len(groups[c]) > len(groups[mx]) || len(groups[c]) == len(groups[mx]) && c > mx {
			mx = c
		}
	}

	return string(groups[mx])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n+|\Sigma|)$，其中 $n$ 是 $\textit{nums}$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(\min(n, |\Sigma|))$ 或 $\mathcal{O}(|\Sigma|)$。

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
