## 提示 1

最后一次操作时，剩下的字母互不相同。

因为如果有相同字母，那么操作后还有剩余字母。

## 提示 2

设字母的最大出现次数为 $\textit{mx}$。

由于删除是从左到右进行的，最后剩下的就是出现次数等于 $\textit{mx}$ 的靠右字母（相同字母取出现位置最右的）。

[视频讲解](https://www.bilibili.com/video/BV1Sm411U7cR/)

```py [sol-Python3]
class Solution:
    def lastNonEmptyString(self, s: str) -> str:
        last = {c: i for i, c in enumerate(s)}
        cnt = Counter(s)
        mx = max(cnt.values())
        ids = sorted(last[ch] for ch, c in cnt.items() if c == mx)
        return ''.join(s[i] for i in ids)
```

```java [sol-Java]
class Solution {
    public String lastNonEmptyString(String S) {
        int[] cnt = new int[26];
        int[] last = new int[26];
        char[] s = S.toCharArray();
        for (int i = 0; i < s.length; i++) {
            int b = s[i] - 'a';
            cnt[b]++;
            last[b] = i;
        }

        // 注：也可以再遍历一次 s 直接得到答案，但效率不如下面，毕竟至多 26 个数
        List<Integer> ids = new ArrayList<>();
        int mx = Arrays.stream(cnt).max().getAsInt();
        for (int i = 0; i < 26; i++) {
            if (cnt[i] == mx) {
                ids.add(last[i]);
            }
        }
        Collections.sort(ids);

        StringBuilder t = new StringBuilder(ids.size()); // 预分配空间
        for (int i : ids) {
            t.append(s[i]);
        }
        return t.toString();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string lastNonEmptyString(string s) {
        int cnt[26]{}, last[26]{};
        for (int i = 0; i < s.size(); i++) {
            int b = s[i] - 'a';
            cnt[b]++;
            last[b] = i;
        }

        // 注：也可以再遍历一次 s 直接得到答案，但效率不如下面，毕竟至多 26 个数
        vector<int> ids;
        int mx = ranges::max(cnt);
        for (int i = 0; i < 26; i++) {
            if (cnt[i] == mx) {
                ids.push_back(last[i]);
            }
        }
        ranges::sort(ids);

        string t(ids.size(), 0);
        for (int i = 0; i < ids.size(); i++) {
            t[i] = s[ids[i]];
        }
        return t;
    }
};
```

```go [sol-Go]
func lastNonEmptyString(s string) string {
	var cnt, last [26]int
	for i, b := range s {
		b -= 'a'
		cnt[b]++
		last[b] = i
	}

	// 注：也可以再遍历一次 s 直接得到答案，但效率不如下面，毕竟至多 26 个数
	ids := []int{}
	mx := slices.Max(cnt[:])
	for i, c := range cnt {
		if c == mx {
			ids = append(ids, last[i])
		}
	}
	slices.Sort(ids)

	t := make([]byte, len(ids))
	for i, id := range ids {
		t[i] = s[id]
	}
	return string(t)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|\log |\Sigma|)$，其中 $n$ 为 $s$ 的长度，$|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
