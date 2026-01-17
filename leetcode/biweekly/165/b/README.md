## 贪心

当物品 $x = \textit{arrivals}[i]$ 进入窗口时，如果发现窗口中 $x$ 的出现次数超过 $m$，我们有两种选择：

- 丢弃刚进入窗口的物品 $x$。
- 如果窗口中还有另一个物品 $x'=x$，我们可以时光倒流到 $x'$ 进入窗口的那个时刻，将 $x'$ 丢弃。

哪种选择更好？

晚丢更好，如果后面还有更多的 $x$，那么晚丢会让窗口中的 $x$ 的个数比早丢更少，后面丢弃的物品数更少。

## 定长滑动窗口

关于定长滑动窗口的原理，见[【套路】教你解决定长滑窗](https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/solutions/2809359/tao-lu-jiao-ni-jie-jue-ding-chang-hua-ch-fzfo/)。

注意一开始的长为 $1,2,3,\dots,w-1$ 的窗口也要考察是否需要丢弃物品。

用一个哈希表（或者数组）$\textit{cnt}$ 统计窗口中的每个元素的出现次数。

元素 $x = \textit{arrivals}[i]$ 进入窗口时：

- 如果 $\textit{cnt}[x]=m$，那么丢弃 $x$，答案加一。⚠**易错点**：$x$ 在未来要离开窗口，但由于已经丢弃，不能在离开窗口时修改 $\textit{cnt}$。为了方便实现，我们可以把 $\textit{arrivals}[i]$ 改成 $0$（或者负数），表示已丢弃。
- 否则把 $\textit{cnt}[x]$ 加一。

元素 $x = \textit{arrivals}[i+1-w]$ 离开窗口时：

- 把 $\textit{cnt}[x]$ 减一。

[本题视频讲解](https://www.bilibili.com/video/BV1TXHZzUE3K/?t=2m11s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minArrivalsToDiscard(self, arrivals: List[int], w: int, m: int) -> int:
        cnt = [0] * (max(arrivals) + 1)  # 或者用 defaultdict(int)
        ans = 0
        for i, x in enumerate(arrivals):
            # x 进入窗口
            if cnt[x] == m:  # x 的个数已达上限
                # 注意 x 在未来要离开窗口，但由于已经丢弃，不能在离开窗口时修改 cnt
                # 这里直接置为 0，未来离开窗口就是 cnt[0]--，不影响答案
                arrivals[i] = 0  # 丢弃 arrivals[i]
                ans += 1
            else:
                cnt[x] += 1

            # 左端点元素离开窗口，为下一个循环做准备
            left = i + 1 - w
            if left >= 0:
                cnt[arrivals[left]] -= 1
        return ans
```

```java [sol-Java]
class Solution {
    public int minArrivalsToDiscard(int[] arrivals, int w, int m) {
        Map<Integer, Integer> cnt = new HashMap<>(); // 更快的写法见【Java 数组】
        int ans = 0;
        for (int i = 0; i < arrivals.length; i++) {
            int x = arrivals[i];
            // x 进入窗口
            int c = cnt.getOrDefault(x, 0);
            if (c == m) { // x 的个数已达上限
                // 注意 x 在未来要离开窗口，但由于已经丢弃，不能在离开窗口时修改 cnt
                // 这里直接置为 0，未来离开窗口就是 cnt[0]--，不影响答案
                arrivals[i] = 0; // 丢弃 arrivals[i]
                ans++;
            } else {
                cnt.put(x, c + 1);
            }

            // 左端点元素离开窗口，为下一个循环做准备
            int left = i + 1 - w;
            if (left >= 0) {
                cnt.merge(arrivals[left], -1, Integer::sum); // cnt[arrivals[left]]--
            }
        }
        return ans;
    }
}
```

```java [sol-Java 数组]
class Solution {
    public int minArrivalsToDiscard(int[] arrivals, int w, int m) {
        int mx = 0;
        for (int x : arrivals) {
            mx = Math.max(mx, x);
        }

        int[] cnt = new int[mx + 1];
        int ans = 0;
        for (int i = 0; i < arrivals.length; i++) {
            int x = arrivals[i];
            // x 进入窗口
            if (cnt[x] == m) { // x 的个数已达上限
                // 注意 x 在未来要离开窗口，但由于已经丢弃，不能在离开窗口时修改 cnt
                // 这里直接置为 0，未来离开窗口就是 cnt[0]--，不影响答案
                arrivals[i] = 0; // 丢弃 arrivals[i]
                ans++;
            } else {
                cnt[x]++;
            }

            // 左端点元素离开窗口，为下一个循环做准备
            int left = i + 1 - w;
            if (left >= 0) {
                cnt[arrivals[left]]--;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minArrivalsToDiscard(vector<int>& arrivals, int w, int m) {
        unordered_map<int, int> cnt;
        int ans = 0;
        for (int i = 0; i < arrivals.size(); i++) {
            int& x = arrivals[i];
            // x 进入窗口
            if (cnt[x] == m) { // x 的个数已达上限
                // 注意 x 在未来要离开窗口，但由于已经丢弃，不能在离开窗口时修改 cnt
                // 这里直接置为 0，未来离开窗口就是 cnt[0]--，不影响答案
                x = 0; // 丢弃 arrivals[i]
                ans++;
            } else {
                cnt[x]++;
            }

            // 左端点元素离开窗口，为下一个循环做准备
            int left = i + 1 - w;
            if (left >= 0) {
                cnt[arrivals[left]]--;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minArrivalsToDiscard(arrivals []int, w, m int) (ans int) {
	cnt := map[int]int{}
	for i, x := range arrivals {
		// x 进入窗口
		if cnt[x] == m { // x 的个数已达上限
			// 注意 x 在未来要离开窗口，但由于已经丢弃，不能在离开窗口时修改 cnt
			// 这里直接置为 0，未来离开窗口就是 cnt[0]--，不影响答案
			arrivals[i] = 0 // 丢弃 arrivals[i]
			ans++
		} else {
			cnt[x]++
		}

		// 左端点元素离开窗口，为下一个循环做准备
		left := i + 1 - w
		if left >= 0 {
			cnt[arrivals[left]]--
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n + U)$，其中 $n$ 是 $\textit{arrivals}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n + U)$。

## 专题训练

见下面滑动窗口题单的「**一、定长滑动窗口**」。

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
