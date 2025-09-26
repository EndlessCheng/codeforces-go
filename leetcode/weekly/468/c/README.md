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

[本题视频讲解](https://www.bilibili.com/video/BV19GWcziEYE/?t=7m23s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Solution:
    def minSplitMerge(self, nums1: List[int], nums2: List[int]) -> int:
        n = len(nums1)
        vis = {tuple(nums1)}
        q = [nums1]
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
        int n = nums1.length;
        List<Integer> nums2List = toList(nums2);
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
                            if (vis.add(c)) { // c 不在 vis 中
                                q.add(c);
                            }
                        }
                    }
                }
            }
        }
    }

    private List<Integer> toList(int[] nums) {
        return Arrays.stream(nums).boxed().collect(Collectors.toList());
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minSplitMerge(vector<int>& nums1, vector<int>& nums2) {
        int n = nums1.size();
        set<vector<int>> vis = {nums1}; // 哈希集合写法见【C++ 写法二】
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

```cpp [sol-C++ 写法二]
template<typename T>
struct VectorHash {
    static void hash_combine(size_t& seed, const T& v) {
        // 参考 boost::hash_combine
        seed ^= hash<T>{}(v) + 0x9e3779b9 + (seed << 6) + (seed >> 2);
    }

    size_t operator()(const vector<T>& vec) const {
        size_t seed = 0;
        for (const T& v : vec) {
            hash_combine(seed, v);
        }
        return seed;
    }
};

class Solution {
public:
    int minSplitMerge(vector<int>& nums1, vector<int>& nums2) {
        int n = nums1.size();
        unordered_set<vector<int>, VectorHash<int>> vis = {nums1};
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

- 时间复杂度：$\mathcal{O}(n!\cdot n^4)$，其中 $n$ 是 $\textit{nums}_1$ 的长度。最多有 $\mathcal{O}(n!)$ 个状态，每个状态有 $\mathcal{O}(n^3)$ 种操作，每种操作需要 $\mathcal{O}(n)$ 的时间拼接数组。
- 空间复杂度：$\mathcal{O}(n!\cdot n)$。最多有 $\mathcal{O}(n!)$ 个状态，每个状态占用 $\mathcal{O}(n)$ 的空间。

## 位运算优化

先把 $\textit{nums}_1$ 和 $\textit{nums}_2$ 离散化，把元素值映射为 $[0,5]$ 中的整数。

由于一个在 $[0,5]$ 中的整数只需要 $3$ 个比特存储，所以一个数组至多要用 $3\times 6 = 18$ 个比特存储，所以可以用一个 $\texttt{int}$ 表示一个数组。

然后就是用位运算代替数组的操作了。

此外，可以在入队前判断是否到达终点，而不是等到出队时再判断。

```py [sol-Python3]
class Solution:
    def minSplitMerge(self, nums1: List[int], nums2: List[int]) -> int:
        if nums1 == nums2:
            return 0

        n = len(nums1)
        mp = {x: i for i, x in enumerate(set(nums1))}  # 用于离散化的映射表
        val1 = sum(mp[x] << (i * 3) for i, x in enumerate(nums1))
        val2 = sum(mp[x] << (i * 3) for i, x in enumerate(nums2))

        vis = {val1}
        q = [val1]
        for ans in count(1):
            tmp = q
            q = []
            for a in tmp:
                for r in range(1, n + 1):  # 为方便实现，先枚举 r，再枚举 l
                    v = a & ((1 << (r * 3)) - 1)
                    for l in range(r):
                        sub = v >> (l * 3)
                        b = a & ((1 << (l * 3)) - 1) | a >> (r * 3) << (l * 3)  # 从 a 中移除 sub
                        for i in range(n - r + l + 1):
                            c = b & ((1 << (i * 3)) - 1) | sub << (i * 3) | b >> (i * 3) << ((i + r - l) * 3)
                            if c == val2:
                                return ans
                            if c not in vis:
                                vis.add(c)
                                q.append(c)
```

```java [sol-Java]
class Solution {
    public int minSplitMerge(int[] nums1, int[] nums2) {
        if (Arrays.equals(nums1, nums2)) {
            return 0;
        }

        int n = nums1.length;
        int[] sorted = nums1.clone(); // 用于离散化
        Arrays.sort(sorted);

        int val1 = encode(nums1, sorted);
        int val2 = encode(nums2, sorted);

        Set<Integer> vis = new HashSet<>();
        vis.add(val1);
        List<Integer> q = List.of(val1);
        for (int ans = 1; ; ans++) {
            List<Integer> tmp = q;
            q = new ArrayList<>();
            for (int a : tmp) {
                for (int r = 1; r <= n; r++) { // 为方便实现，先枚举 r，再枚举 l
                    int t = a & ((1 << (r * 3)) - 1);
                    for (int l = 0; l < r; l++) {
                        int sub = t >> (l * 3);
                        int b = (a & ((1 << (l * 3)) - 1)) | (a >> (r * 3) << (l * 3)); // 从 a 中移除 sub
                        for (int i = 0; i <= n - r + l; i++) {
                            int c = (b & ((1 << (i * 3)) - 1)) | (sub << (i * 3)) | (b >> (i * 3) << ((i + r - l) * 3));
                            if (c == val2) {
                                return ans;
                            }
                            if (vis.add(c)) { // c 不在 vis 中
                                q.add(c);
                            }
                        }
                    }
                }
            }
        }
    }

    private int encode(int[] nums, int[] sorted) {
        int res = 0;
        for (int i = 0; i < nums.length; i++) {
            int x = Arrays.binarySearch(sorted, nums[i]);
            res |= x << (i * 3);
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    int encode(vector<int>& nums, vector<int>& sorted) {
        int res = 0;
        for (int i = 0; i < nums.size(); i++) {
            int x = ranges::lower_bound(sorted, nums[i]) - sorted.begin();
            res |= x << (i * 3);
        }
        return res;
    }

public:
    int minSplitMerge(vector<int>& nums1, vector<int>& nums2) {
        if (nums1 == nums2) {
            return 0;
        }

        int n = nums1.size();
        auto sorted = nums1; // 用于离散化
        ranges::sort(sorted);

        int val1 = encode(nums1, sorted);
        int val2 = encode(nums2, sorted);

        unordered_set<int> vis = {val1};
        vector<int> q = {val1};
        for (int ans = 1; ; ans++) {
            vector<int> nxt;
            for (int a : q) {
                for (int r = 1; r <= n; r++) { // 为方便实现，先枚举 r，再枚举 l
                    int t = a & ((1 << (r * 3)) - 1);
                    for (int l = 0; l < r; l++) {
                        int sub = t >> (l * 3);
                        int b = (a & ((1 << (l * 3)) - 1)) | (a >> (r * 3) << (l * 3)); // 从 a 中移除 sub
                        for (int i = 0; i <= n - r + l; i++) {
                            int c = (b & ((1 << (i * 3)) - 1)) | (sub << (i * 3)) | (b >> (i * 3) << ((i + r - l) * 3));
                            if (c == val2) {
                                return ans;
                            }
                            if (vis.insert(c).second) { // c 不在 vis 中
                                nxt.push_back(c);
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
func encode(nums, sorted []int) (res int) {
	for i, x := range nums {
		res |= sort.SearchInts(sorted, x) << (i * 3)
	}
	return
}

func minSplitMerge(nums1, nums2 []int) int {
	if slices.Equal(nums1, nums2) {
		return 0
	}

	n := len(nums1)
	sorted := slices.Clone(nums1) // 用于离散化
	slices.Sort(sorted)

	val1 := encode(nums1, sorted)
	val2 := encode(nums2, sorted)

	vis := map[int]bool{val1: true}
	q := []int{val1}
	for ans := 1; ; ans++ {
		tmp := q
		q = nil
		for _, a := range tmp {
			for r := 1; r <= n; r++ { // 为方便实现，先枚举 r，再枚举 l
				t := a & (1<<(r*3) - 1)
				for l := range r {
					sub := t >> (l * 3)
					b := a&(1<<(l*3)-1) | a>>(r*3)<<(l*3) // 从 a 中移除 sub
					for i := range n - r + l + 1 {
						c := b&(1<<(i*3)-1) | sub<<(i*3) | b>>(i*3)<<((i+r-l)*3)
						if c == val2 {
							return ans
						}
						if !vis[c] {
							vis[c] = true
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

- 时间复杂度：$\mathcal{O}(n!\cdot n^3)$，其中 $n$ 是 $\textit{nums}_1$ 的长度。最多有 $\mathcal{O}(n!)$ 个状态，每个状态有 $\mathcal{O}(n^3)$ 种操作，每种操作需要 $\mathcal{O}(1)$ 的时间拼接二进制数。
- 空间复杂度：$\mathcal{O}(n!)$。最多有 $\mathcal{O}(n!)$ 个状态，每个状态占用 $\mathcal{O}(1)$ 的空间。

## 专题训练

见下面图论题单的「**§1.3 图论建模 + BFS 最短路**」。

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
