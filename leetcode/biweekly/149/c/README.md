**核心思路**：维护最大的三个空位，其中一定有一个空位不在会议左右两侧，把会议移到这个空位中。

设最大的三个空位所在的位置（下标）分别是 $a,b,c$。

枚举要重新安排（移走）的会议，分类讨论：

- 如果会议左右两侧的空位没有 $a$，那么把会议移到 $a$。
- 否则，如果会议左右两侧的空位没有 $b$，那么把会议移到 $b$。
- 否则，把会议移到 $c$。

继续分类讨论：

- 如果能移（有足够长的空位），那么用会议左右两侧的空位长度之和，再加上会议长度，更新答案的最大值。
- 如果不能移，那么只能移到左右相邻的会议旁边，用会议左右两侧的空位长度之和，更新答案的最大值。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1eUF6eaERQ/?t=4m28s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxFreeTime(self, eventTime: int, startTime: List[int], endTime: List[int]) -> int:
        n = len(startTime)

        def get(i: int) -> int:
            if i == 0:
                return startTime[0]
            if i == n:
                return eventTime - endTime[n - 1]
            return startTime[i] - endTime[i - 1]

        a, b, c = 0, -1, -1
        for i in range(1, n + 1):
            sz = get(i)
            if sz > get(a):
                a, b, c = i, a, b
            elif b < 0 or sz > get(b):
                b, c = i, b
            elif c < 0 or sz > get(c):
                c = i

        ans = 0
        for i, (s, e) in enumerate(zip(startTime, endTime)):
            sz = e - s
            if i != a and i + 1 != a and sz <= get(a) or \
               i != b and i + 1 != b and sz <= get(b) or \
               sz <= get(c):
                ans = max(ans, get(i) + sz + get(i + 1))
            else:
                ans = max(ans, get(i) + get(i + 1))
        return ans
```

```py [sol-Python3 写法二]
class Solution:
    def maxFreeTime(self, eventTime: int, startTime: List[int], endTime: List[int]) -> int:
        free = [startTime[0]] + [s - e for s, e in zip(startTime[1:], endTime)] + [eventTime - endTime[-1]]

        a = b = c = -1
        for i, sz in enumerate(free):
            if a < 0 or sz > free[a]:
                a, b, c = i, a, b
            elif b < 0 or sz > free[b]:
                b, c = i, b
            elif c < 0 or sz > free[c]:
                c = i

        ans = 0
        for i, (s, e) in enumerate(zip(startTime, endTime)):
            sz = e - s
            if i != a and i + 1 != a and sz <= free[a] or \
               i != b and i + 1 != b and sz <= free[b] or \
               sz <= free[c]:
                ans = max(ans, free[i] + sz + free[i + 1])
            else:
                ans = max(ans, free[i] + free[i + 1])
        return ans
```

```java [sol-Java]
class Solution {
    private int eventTime;
    private int[] startTime, endTime;

    public int maxFreeTime(int eventTime, int[] startTime, int[] endTime) {
        this.eventTime = eventTime;
        this.startTime = startTime;
        this.endTime = endTime;
        int n = startTime.length;

        int a = 0, b = -1, c = -1;
        for (int i = 1; i <= n; i++) {
            int sz = get(i);
            if (sz > get(a)) {
                c = b; b = a; a = i;
            } else if (b < 0 || sz > get(b)) {
                c = b; b = i;
            } else if (c < 0 || sz > get(c)) {
                c = i;
            }
        }

        int ans = 0;
        for (int i = 0; i < n; i++) {
            int sz = endTime[i] - startTime[i];
            if (i != a && i + 1 != a && sz <= get(a) ||
                i != b && i + 1 != b && sz <= get(b) ||
                sz <= get(c)) {
                ans = Math.max(ans, get(i) + sz + get(i + 1));
            } else {
                ans = Math.max(ans, get(i) + get(i + 1));
            }
        }
        return ans;
    }

    private int get(int i) {
        if (i == 0) {
            return startTime[0];
        }
        int n = startTime.length;
        if (i == n) {
            return eventTime - endTime[n - 1];
        }
        return startTime[i] - endTime[i - 1];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxFreeTime(int eventTime, vector<int>& startTime, vector<int>& endTime) {
        int n = startTime.size();
        auto get = [&](int i) -> int {
            if (i == 0) {
                return startTime[0];
            }
            if (i == n) {
                return eventTime - endTime[n - 1];
            }
            return startTime[i] - endTime[i - 1];
        };

        int a = 0, b = -1, c = -1;
        for (int i = 1; i <= n; i++) {
            int sz = get(i);
            if (sz > get(a)) {
                c = b; b = a; a = i;
            } else if (b < 0 || sz > get(b)) {
                c = b; b = i;
            } else if (c < 0 || sz > get(c)) {
                c = i;
            }
        }

        int ans = 0;
        for (int i = 0; i < n; i++) {
            int sz = endTime[i] - startTime[i];
            if (i != a && i + 1 != a && sz <= get(a) ||
                i != b && i + 1 != b && sz <= get(b) ||
                sz <= get(c)) {
                ans = max(ans, get(i) + sz + get(i + 1));
            } else {
                ans = max(ans, get(i) + get(i + 1));
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxFreeTime(eventTime int, startTime, endTime []int) (ans int) {
	n := len(startTime)
	get := func(i int) int {
		if i == 0 {
			return startTime[0]
		}
		if i == n {
			return eventTime - endTime[n-1]
		}
		return startTime[i] - endTime[i-1]
	}

	a, b, c := 0, -1, -1
	for i := 1; i <= n; i++ {
		sz := get(i)
		if sz > get(a) {
			a, b, c = i, a, b
		} else if b < 0 || sz > get(b) {
			b, c = i, b
		} else if c < 0 || sz > get(c) {
			c = i
		}
	}

	for i, e := range endTime {
		sz := e - startTime[i]
		if i != a && i+1 != a && sz <= get(a) ||
			i != b && i+1 != b && sz <= get(b) ||
			sz <= get(c) {
			ans = max(ans, get(i)+sz+get(i+1))
		} else {
			ans = max(ans, get(i)+get(i+1))
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{startTime}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
