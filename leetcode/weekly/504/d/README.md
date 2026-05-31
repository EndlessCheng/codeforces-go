**适用场景**：按照题目要求，数组会被分割成若干组，每一组的判断/处理逻辑是相同的。

**核心思想**：

- 外层循环负责遍历组之前的准备工作（记录开始位置），和遍历组之后的统计工作（更新答案）。
- 内层循环负责遍历组，找出这一组最远在哪结束。

这个写法的好处是，各个逻辑块分工明确，也不需要特判最后一组（易错点）。以我的经验，这个写法是所有写法中最不容易出 bug 的，推荐大家记住。

---

对于本题，先统计每个元素 $x$ 的所有下标，记在 $\textit{pos}[x]$ 中。

对于每一组（每一段），枚举这一段的 $\textit{mex}=0,1,2,\ldots$

- 如果 $\textit{nums}$ 的剩余元素包含 $\textit{mex}$，那么为了最大化字典序，这一段必须包含剩余元素中的最左边的 $\textit{mex}$。然后继续枚举 $\textit{mex}$。
- 如果 $\textit{nums}$ 的剩余元素不包含 $\textit{mex}$，跳出内层循环，把 $\textit{mex}$ 添加到答案中。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

## 写法一

```py [sol-Python3]
class Solution:
    def maximumMEX(self, nums: list[int]) -> list[int]:
        n = len(nums)
        # mex 最大是 n，>= n 的数无需考虑
        pos = [deque() for _ in range(n + 1)]  # n 作为哨兵
        for i, x in enumerate(nums):
            if x < n:
                pos[x].append(i)

        ans = []
        i = 0
        while i < n:
            start = i  # 这一段的起点
            # 枚举这一段的 mex，越大越好（字典序越大）
            mex = 0
            while True:
                # 清理在 start 之前的下标
                q = pos[mex]
                while q and q[0] < start:
                    q.popleft()
                if not q:
                    break
                # 这一段包含剩余元素中的最左边的 mex
                i = max(i, q[0])
                mex += 1
            ans.append(mex)
            i += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int[] maximumMEX(int[] nums) {
        int n = nums.length;
        // mex 最大是 n，>= n 的数无需考虑
        ArrayDeque<Integer>[] pos = new ArrayDeque[n + 1]; // n 作为哨兵
        Arrays.setAll(pos, _ -> new ArrayDeque<>());
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            if (x < n) {
                pos[x].addLast(i);
            }
        }

        int idx = 0;
        for (int i = 0; i < n; i++) {
            int start = i; // 这一段的起点
            // 枚举这一段的 mex，越大越好（字典序越大）
            int mex = 0;
            for (; ; mex++) {
                // 清理在 start 之前的下标
                while (!pos[mex].isEmpty() && pos[mex].peekFirst() < start) {
                    pos[mex].pollFirst();
                }
                if (pos[mex].isEmpty()) {
                    break;
                }
                // 这一段包含剩余元素中的最左边的 mex
                i = Math.max(i, pos[mex].peekFirst());
            }
            nums[idx++] = mex;
        }
        return Arrays.copyOf(nums, idx);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> maximumMEX(vector<int>& nums) {
        int n = nums.size();
        // mex 最大是 n，>= n 的数无需考虑
        vector<queue<int>> pos(n + 1); // n 作为哨兵
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            if (x < n) {
                pos[x].push(i);
            }
        }

        vector<int> ans;
        for (int i = 0; i < n; i++) {
            int start = i; // 这一段的起点
            // 枚举这一段的 mex，越大越好（字典序越大）
            int mex = 0;
            for (; ; mex++) {
                // 清理在 start 之前的下标
                auto& q = pos[mex];
                while (!q.empty() && q.front() < start) {
                    q.pop();
                }
                if (q.empty()) {
                    break;
                }
                // 这一段包含剩余元素中的最左边的 mex
                i = max(i, q.front());
            }
            ans.push_back(mex);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumMEX(nums []int) (ans []int) {
	n := len(nums)
	// mex 最大是 n，>= n 的数无需考虑
	pos := make([][]int, n+1) // n 作为哨兵
	for i, x := range nums {
		if x < n {
			pos[x] = append(pos[x], i)
		}
	}

	for i := 0; i < n; i++ {
		start := i // 这一段的起点
		// 枚举这一段的 mex，越大越好（字典序越大）
		mex := 0
		for ; ; mex++ {
			// 清理在 start 之前的下标
			for len(pos[mex]) > 0 && pos[mex][0] < start {
				pos[mex] = pos[mex][1:]
			}
			if len(pos[mex]) == 0 {
				break
			}
			// 这一段包含剩余元素中的最左边的 mex
			i = max(i, pos[mex][0])
		}
		ans = append(ans, mex)
	}
	return
}
```

## 写法二

把 $\textit{pos}[x]$ 反转，就可以用列表（数组）代替队列了。

Go 语言由于删除切片的第一个元素是 $\mathcal{O}(1)$ 的，不需要这样做。

```py [sol-Python3]
class Solution:
    def maximumMEX(self, nums: list[int]) -> list[int]:
        n = len(nums)
        # mex 最大是 n，>= n 的数无需考虑
        pos = [[] for _ in range(n + 1)]  # n 作为哨兵
        for i, x in enumerate(nums):
            if x < n:
                pos[x].append(i)
        for p in pos:
            p.reverse()

        ans = []
        i = 0
        while i < n:
            start = i  # 这一段的起点
            # 枚举这一段的 mex，越大越好（字典序越大）
            mex = 0
            while True:
                # 清理在 start 之前的下标
                q = pos[mex]
                while q and q[-1] < start:
                    q.pop()
                if not q:
                    break
                # 这一段包含剩余元素中的最左边的 mex
                i = max(i, q[-1])
                mex += 1
            ans.append(mex)
            i += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int[] maximumMEX(int[] nums) {
        int n = nums.length;
        // mex 最大是 n，>= n 的数无需考虑
        List<Integer>[] pos = new ArrayList[n + 1]; // n 作为哨兵
        Arrays.setAll(pos, _ -> new ArrayList<>());
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            if (x < n) {
                pos[x].addLast(i);
            }
        }
        for (List<Integer> p : pos) {
            Collections.reverse(p);
        }

        int idx = 0;
        for (int i = 0; i < n; i++) {
            int start = i; // 这一段的起点
            // 枚举这一段的 mex，越大越好（字典序越大）
            int mex = 0;
            for (; ; mex++) {
                // 清理在 start 之前的下标
                while (!pos[mex].isEmpty() && pos[mex].getLast() < start) {
                    pos[mex].removeLast();
                }
                if (pos[mex].isEmpty()) {
                    break;
                }
                // 这一段包含剩余元素中的最左边的 mex
                i = Math.max(i, pos[mex].getLast());
            }
            nums[idx++] = mex;
        }
        return Arrays.copyOf(nums, idx);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> maximumMEX(vector<int>& nums) {
        int n = nums.size();
        // mex 最大是 n，>= n 的数无需考虑
        vector<vector<int>> pos(n + 1); // n 作为哨兵
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            if (x < n) {
                pos[x].push_back(i);
            }
        }
        for (auto& p : pos) {
            ranges::reverse(p);
        }

        vector<int> ans;
        for (int i = 0; i < n; i++) {
            int start = i; // 这一段的起点
            int mex = 0;
            // 枚举这一段的 mex，越大越好（字典序越大）
            for (; ; mex++) {
                // 清理在 start 之前的下标
                auto& q = pos[mex];
                while (!q.empty() && q.back() < start) {
                    q.pop_back();
                }
                if (q.empty()) {
                    break;
                }
                // 这一段包含剩余元素中的最左边的 mex
                i = max(i, q.back());
            }
            ans.push_back(mex);
        }
        return ans;
    }
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。注意 $i$ 只会增大，不会减小或重置，以及 $\textit{pos}$ 至多有 $n$ 个下标，所以三重循环的总循环次数为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面双指针题单的「**六、分组循环**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
