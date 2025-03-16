## 方法一：二分查找

看示例 1，所有 $1$ 的下标列表是 $p=[0,2,4]$。

由于 $\textit{nums}$ 是循环数组：

- 在下标列表前面添加 $4-n=-3$，相当于认为在 $-3$ 下标处也有一个 $1$。
- 在下标列表末尾添加 $0+n=7$，相当于认为在 $7$ 下标处也有一个 $1$。

修改后的下标列表为 $p=[-3,0,2,4,7]$。

于是，我们在 $p$ 中二分查找下标 $i$，设二分返回值为 $j$，那么：

- $p[j-1]$ 就是在 $i$ 左边的最近位置。
- $p[j+1]$ 就是在 $i$ 右边的最近位置。

两个距离取最小值，答案为

$$
\min(i-p[j-1], p[j+1]-i)
$$

如果 $\textit{nums}[i]$ 在 $\textit{nums}$ 中只出现了一次，那么答案为 $-1$。

代码实现时，可以直接把答案记录在 $\textit{queries}$ 数组中。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1JYQ8YWEvD/?t=3m28s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def solveQueries(self, nums: List[int], queries: List[int]) -> List[int]:
        indices = defaultdict(list)
        for i, x in enumerate(nums):
            indices[x].append(i)

        n = len(nums)
        for p in indices.values():
            # 前后各加一个哨兵
            i0 = p[0]
            p.insert(0, p[-1] - n)
            p.append(i0 + n)

        for qi, i in enumerate(queries):
            p = indices[nums[i]]
            if len(p) == 3:
                queries[qi] = -1
            else:
                j = bisect_left(p, i)
                queries[qi] = min(i - p[j - 1], p[j + 1] - i)
        return queries
```

```java [sol-Java]
class Solution {
    public List<Integer> solveQueries(int[] nums, int[] queries) {
        Map<Integer, List<Integer>> indices = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            indices.computeIfAbsent(nums[i], k -> new ArrayList<>()).add(i);
        }

        int n = nums.length;
        for (List<Integer> p : indices.values()) {
            // 前后各加一个哨兵
            int i0 = p.get(0);
            p.add(0, p.get(p.size() - 1) - n);
            p.add(i0 + n);
        }

        List<Integer> ans = new ArrayList<>(queries.length); // 预分配空间
        for (int i : queries) {
            List<Integer> p = indices.get(nums[i]);
            if (p.size() == 3) {
                ans.add(-1);
            } else {
                int j = Collections.binarySearch(p, i);
                ans.add(Math.min(i - p.get(j - 1), p.get(j + 1) - i));
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> solveQueries(vector<int>& nums, vector<int>& queries) {
        unordered_map<int, vector<int>> indices;
        for (int i = 0; i < nums.size(); i++) {
            indices[nums[i]].push_back(i);
        }

        int n = nums.size();
        for (auto& [_, p] : indices) {
            // 前后各加一个哨兵
            int i0 = p[0];
            p.insert(p.begin(), p.back() - n);
            p.push_back(i0 + n);
        }

        for (int& i : queries) { // 注意这里是引用
            auto& p = indices[nums[i]];
            if (p.size() == 3) {
                i = -1;
            } else {
                int j = ranges::lower_bound(p, i) - p.begin();
                i = min(i - p[j - 1], p[j + 1] - i);
            }
        }
        return queries;
    }
};
```

```go [sol-Go]
func solveQueries(nums []int, queries []int) []int {
	indices := map[int][]int{}
	for i, x := range nums {
		indices[x] = append(indices[x], i)
	}

	n := len(nums)
	for x, p := range indices {
		// 前后各加一个哨兵
		i0 := p[0]
		p = slices.Insert(p, 0, p[len(p)-1]-n)
		indices[x] = append(p, i0+n)
	}

	for qi, i := range queries {
		p := indices[nums[i]]
		if len(p) == 3 {
			queries[qi] = -1
		} else {
			j := sort.SearchInts(p, i)
			queries[qi] = min(i-p[j-1], p[j+1]-i)
		}
	}
	return queries
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。每次二分需要 $\mathcal{O}(\log n)$ 的时间。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 方法二：预处理左右最近相同元素的下标

定义 $\textit{left}[i]$ 表示在 $i$ 左边的等于 $\textit{nums}[i]$ 的最近元素下标。注意数组是循环数组，我们可以像方法一那样，用 $-1$ 表示最后一个数的下标，$-2$ 表示倒数第二个数的下标，依此类推。

定义 $\textit{right}[i]$ 表示在 $i$ 右边的等于 $\textit{nums}[i]$ 的最近元素下标。注意数组是循环数组，我们可以像方法一那样，用 $n$ 表示第一个数的下标，$n+1$ 表示第二个数的下标，依此类推。

计算方式：

- 从 $-n$ 循环到 $n-1$，同时用一个哈希表记录每个数的最新位置。当 $i\ge 0$ 时，$\textit{left}[i]$ 就是哈希中记录的 $\textit{nums}[i]$ 的位置。注意先更新 $\textit{left}[i]$ 再更新哈希表。
- 从 $2n-1$ 循环到 $0$，同时用一个哈希表记录每个数的最新位置。当 $i < n$ 时，$\textit{right}[i]$ 就是哈希中记录的 $\textit{nums}[i]$ 的位置。注意先更新 $\textit{right}[i]$ 再更新哈希表。

答案为：

$$
\min(i-\textit{left}[i], \textit{right}[i]-i)
$$

如果上式等于 $n$，说明只有一个 $\textit{nums}[i]$，答案为 $-1$。

### 优化前

```py [sol-Python3]
class Solution:
    def solveQueries(self, nums: List[int], queries: List[int]) -> List[int]:
        n = len(nums)
        left = [0] * n
        pos = {}
        for i in range(-n, n):
            if i >= 0:
                left[i] = pos[nums[i]]
            pos[nums[i]] = i

        right = [0] * n
        pos.clear()
        for i in range(n * 2 - 1, -1, -1):
            if i < n:
                right[i] = pos[nums[i]]
            pos[nums[i % n]] = i

        for qi, i in enumerate(queries):
            l = left[i]
            queries[qi] = -1 if i - l == n else min(i - l, right[i] - i)
        return queries
```

```java [sol-Java]
class Solution {
    public List<Integer> solveQueries(int[] nums, int[] queries) {
        int n = nums.length;
        int[] left = new int[n];
        Map<Integer, Integer> pos = new HashMap<>();
        for (int i = -n; i < n; i++) {
            if (i >= 0) {
                left[i] = pos.get(nums[i]);
            }
            pos.put(nums[(i + n) % n], i);
        }

        int[] right = new int[n];
        pos.clear();
        for (int i = n * 2 - 1; i >= 0; i--) {
            if (i < n) {
                right[i] = pos.get(nums[i]);
            }
            pos.put(nums[i % n], i);
        }

        List<Integer> ans = new ArrayList<>(queries.length);
        for (int i : queries) {
            int l = left[i];
            ans.add(i - l == n ? -1 : Math.min(i - l, right[i] - i));
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> solveQueries(vector<int>& nums, vector<int>& queries) {
        int n = nums.size();
        vector<int> left(n), right(n);
        unordered_map<int, int> pos;
        for (int i = -n; i < n; i++) {
            if (i >= 0) {
                left[i] = pos[nums[i]];
            }
            pos[nums[(i + n) % n]] = i;
        }

        pos.clear();
        for (int i = n * 2 - 1; i >= 0; i--) {
            if (i < n) {
                right[i] = pos[nums[i]];
            }
            pos[nums[i % n]] = i;
        }

        for (int& i : queries) {
            int l = left[i];
            i = i - l == n ? -1 : min(i - l, right[i] - i);
        }
        return queries;
    }
};
```

```go [sol-Go]
func solveQueries(nums []int, queries []int) []int {
	n := len(nums)
	left := make([]int, n)
	pos := map[int]int{}
	for i := -n; i < n; i++ {
		if i >= 0 {
			left[i] = pos[nums[i]]
		}
		pos[nums[(i+n)%n]] = i
	}

	right := make([]int, n)
	clear(pos)
	for i := n*2 - 1; i >= 0; i-- {
		if i < n {
			right[i] = pos[nums[i]]
		}
		pos[nums[i%n]] = i
	}

	for qi, i := range queries {
		l := left[i]
		if i-l == n {
			queries[qi] = -1
		} else {
			queries[qi] = min(i-l, right[i]-i)
		}
	}
	return queries
}
```

### 优化

一次遍历同时计算 $\textit{left}$ 和 $\textit{right}$。（类似单调栈）

```py [sol-Python3]
class Solution:
    def solveQueries(self, nums: List[int], queries: List[int]) -> List[int]:
        n = len(nums)
        left = [0] * n
        right = [0] * n
        pos = {}
        for i in range(-n, n):
            x = nums[i]
            if i >= 0:
                j = pos[x]
                left[i] = j
                # 对于左边的 j 来说，它的 right 就是 i
                if j >= 0:
                    right[j] = i
                else:
                    right[j + n] = i + n
            pos[x] = i

        for qi, i in enumerate(queries):
            l = left[i]
            queries[qi] = -1 if i - l == n else min(i - l, right[i] - i)
        return queries
```

```java [sol-Java]
class Solution {
    public List<Integer> solveQueries(int[] nums, int[] queries) {
        int n = nums.length;
        int[] left = new int[n];
        int[] right = new int[n];
        Map<Integer, Integer> pos = new HashMap<>();
        for (int i = -n; i < n; i++) {
            if (i >= 0) {
                int j = pos.get(nums[i]);
                left[i] = j;
                // 对于左边的 j 来说，它的 right 就是 i
                if (j >= 0) {
                    right[j] = i;
                } else {
                    right[j + n] = i + n;
                }
            }
            pos.put(nums[(i + n) % n], i);
        }

        List<Integer> ans = new ArrayList<>(queries.length);
        for (int i : queries) {
            int l = left[i];
            ans.add(i - l == n ? -1 : Math.min(i - l, right[i] - i));
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> solveQueries(vector<int>& nums, vector<int>& queries) {
        int n = nums.size();
        vector<int> left(n), right(n);
        unordered_map<int, int> pos;
        for (int i = -n; i < n; i++) {
            if (i >= 0) {
                int j = pos[nums[i]];
                left[i] = j;
                // 对于左边的 j 来说，它的 right 就是 i
                if (j >= 0) {
                    right[j] = i;
                } else {
                    right[j + n] = i + n;
                }
            }
            pos[nums[(i + n) % n]] = i;
        }

        for (int& i : queries) {
            int l = left[i];
            i = i - l == n ? -1 : min(i - l, right[i] - i);
        }
        return queries;
    }
};
```

```go [sol-Go]
func solveQueries(nums []int, queries []int) []int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	pos := map[int]int{}
	for i := -n; i < n; i++ {
		if i >= 0 {
			j := pos[nums[i]]
			left[i] = j
			// 对于左边的 j 来说，它的 right 就是 i
			if j >= 0 {
				right[j] = i
			} else {
				right[j+n] = i + n
			}
		}
		pos[nums[(i+n)%n]] = i
	}

	for qi, i := range queries {
		l := left[i]
		if i-l == n {
			queries[qi] = -1
		} else {
			queries[qi] = min(i-l, right[i]-i)
		}
	}
	return queries
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 相似题目

- [2080. 区间内查询数字的频率](https://leetcode.cn/problems/range-frequency-queries/) 1702

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. 【本题相关】[二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
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
