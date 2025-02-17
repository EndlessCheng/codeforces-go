对于每种字母，分别用一个栈维护未标记字母的下标。

遍历到 $s[i]$ 时，分类讨论：

- 如果在 $s[i]$ 左边没有 $s[i]$ 的镜像字母，那么把下标 $i$ 加到第 $i$ 个栈中。
- 否则，弹出 $s[i]$ 的镜像字母对应的栈顶，即为我们要找的 $j$。把 $i-j$ 加入答案。

[本题视频讲解](https://www.bilibili.com/video/BV18srKYLEd8/?t=5m59s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def calculateScore(self, s: str) -> int:
        stk = [[] for _ in range(26)]
        ans = 0
        for i, c in enumerate(map(ord, s)):
            c -= ord('a')
            if stk[25 - c]:
                ans += i - stk[25 - c].pop()
            else:
                stk[c].append(i)
        return ans
```

```java [sol-Java]
class Solution {
    public long calculateScore(String s) {
        Deque<Integer>[] stk = new ArrayDeque[26];
        Arrays.setAll(stk, i -> new ArrayDeque<>());
        long ans = 0;
        for (int i = 0; i < s.length(); i++) {
            int c = s.charAt(i) - 'a';
            if (!stk[25 - c].isEmpty()) {
                ans += i - stk[25 - c].pop();
            } else {
                stk[c].push(i);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long calculateScore(string s) {
        stack<int> stk[26];
        long long ans = 0;
        for (int i = 0; i < s.size(); i++) {
            int c = s[i] - 'a';
            if (!stk[25 - c].empty()) {
                ans += i - stk[25 - c].top();
                stk[25 - c].pop();
            } else {
                stk[c].push(i);                
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func calculateScore(s string) (ans int64) {
	stk := [26][]int{}
	for i, c := range s {
		c -= 'a'
		if st := stk[25-c]; len(st) > 0 {
			ans += int64(i - st[len(st)-1])
			stk[25-c] = st[:len(st)-1]
		} else {
			stk[c] = append(stk[c], i)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(n+|\Sigma|)$。

更多相似题目，见下面数据结构题单中的「**三、栈**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. 【本题相关】[常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
