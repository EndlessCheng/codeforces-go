由于操作不会增加或删除元素，只会改变元素的顺序，所以每次操作后，我们得到的是 $\textit{nums}_1$ 的一个排列。

根据数据范围，至多有 $6! = 720$ 个不同的排列，这很小，考虑暴力。

每次操作相当于从一个数组变成另一个数组。抽象成有向图，把数组视作节点，操作就是在一个点到另一个点之间连一条有向边。

问题变成：

- 计算从起点 $\textit{nums}_1$ 到终点 $\textit{nums}_2$ 的**最短路**长度。

由于边权都是 $1$，可以用 **BFS** 解决。暴力枚举所有子数组和所有插入位置。

关于 BFS 的原理和双列表写法，见[【基础算法精讲 13】](https://www.bilibili.com/video/BV1hG4y1277i/)。

#### 答疑

**问**：为什么本题一定有解？

**答**：我们可以把一个元素（当作子数组）提取出来，插在数组的任意位置。类似扑克牌，我们**可以随意洗牌**。所以可以把 $\textit{nums}_1$ 变成其任意排列。由于题目保证 $\textit{nums}_2$ 是 $\textit{nums}_1$ 的一个排列，所以一定可以把 $\textit{nums}_1$ 变成 $\textit{nums}_2$。 

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def minSplitMerge(self, nums1: List[int], nums2: List[int]) -> int:
        n = len(nums1)
        vis = {tuple(nums1)}
        q = deque([nums1])
        for ans in count(0):
            tmp = q
            q = []
            for a in tmp:
                if a == nums2:
                    return ans
                for l in range(n):
                    for r in range(l + 1, n + 1):
                        sub = a[l: r]
                        b = a[:l] + a[r:]  # 从 a 中移除 sub
                        for i in range(len(b) + 1):
                            c = b[:i] + sub + b[i:]
                            t = tuple(c)
                            if t not in vis:
                                vis.add(t)
                                q.append(c)
```

```java [sol-Java]
class Solution {
    public int minSplitMerge(int[] nums1, int[] nums2) {
        List<Integer> nums2List = toList(nums2);
        int n = nums1.length;
        Set<List<Integer>> vis = new HashSet<>();
        vis.add(toList(nums1));
        List<List<Integer>> q = List.of(toList(nums1));
        for (int ans = 0; ; ans++) {
            List<List<Integer>> tmp = q;
            q = new ArrayList<>();
            for (List<Integer> a : tmp) {
                if (a.equals(nums2List)) {
                    return ans;
                }
                for (int l = 0; l < n; l++) {
                    for (int r = l + 1; r <= n; r++) {
                        List<Integer> sub = a.subList(l, r);
                        List<Integer> b = new ArrayList<>(a);
                        b.subList(l, r).clear(); // 从 b 中移除 sub
                        for (int i = 0; i <= b.size(); i++) {
                            List<Integer> c = new ArrayList<>(b);
                            c.addAll(i, sub);
                            if (vis.add(c)) {
                                q.add(c);
                            }
                        }
                    }
                }
            }
        }
    }

    private List<Integer> toList(int[] a) {
        List<Integer> b = new ArrayList<>();
        for (int x : a) {
            b.add(x);
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minSplitMerge(vector<int>& nums1, vector<int>& nums2) {
        int n = nums1.size();
        set<vector<int>> vis = {nums1};
        vector<vector<int>> q = {nums1};
        for (int ans = 0; ; ans++) {
            vector<vector<int>> nxt;
            for (auto& a : q) {
                if (a == nums2) {
                    return ans;
                }
                for (int l = 0; l < n; l++) {
                    for (int r = l + 1; r <= n; r++) {
                        auto b = a;
                        vector<int> sub(b.begin() + l, b.begin() + r);
                        b.erase(b.begin() + l, b.begin() + r); // 从 b 中移除 sub
                        for (int i = 0; i <= b.size(); i++) {
                            auto c = b;
                            c.insert(c.begin() + i, sub.begin(), sub.end());
                            if (vis.insert(c).second) { // c 不在 vis 中
                                nxt.emplace_back(c);
                            }
                        }
                    }
                }
            }
            q = move(nxt);
        }
    }
};
```

```go [sol-Go]
func minSplitMerge(nums1, nums2 []int) (ans int) {
	n := len(nums1)
	t := [6]int{}
	for j, x := range nums1 {
		t[j] = x
	}
	vis := map[[6]int]bool{t: true}
	q := [][]int{nums1}
	for ; ; ans++ {
		tmp := q
		q = nil
		for _, a := range tmp {
			if slices.Equal(a, nums2) {
				return
			}
			for l := 0; l < n; l++ {
				for r := l + 1; r <= n; r++ {
					b := slices.Clone(a)
					sub := slices.Clone(b[l:r])
					b = append(b[:l], b[r:]...) // 从 b 中移除 sub
					for i := 0; i <= len(b); i++ {
						c := slices.Insert(slices.Clone(b), i, sub...)
						t := [6]int{}
						for j, x := range c {
							t[j] = x
						}
						if !vis[t] {
							vis[t] = true
							q = append(q, c)
						}
					}
				}
			}
		}
	}
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n!\cdot n^4)$，其中 $n$ 是 $\textit{nums}_1$ 的长度。
- 空间复杂度：$\mathcal{O}(n!\cdot n)$。最多有 $\mathcal{O}(n!)$ 个状态，每个状态占用 $\mathcal{O}(n)$ 的空间。

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
