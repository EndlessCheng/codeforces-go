## 方法一：从左到右滑窗

题目要求把子数组操作成递增的（允许相等），这可以贪心地做，具体操作方法见 [3402 题的题解](https://leetcode.cn/problems/minimum-operations-to-make-columns-strictly-increasing/solutions/3033305/cong-shang-dao-xia-tan-xin-pythonjavacgo-dvhp/)。

由于子数组越长，操作次数越多，有单调性，可以用 [滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)解决。

### 右端点元素进入窗口

我们需要知道窗口内的最大值，即 [239. 滑动窗口最大值](https://leetcode.cn/problems/sliding-window-maximum/)，原理请看 [单调队列【基础算法精讲 27】](https://www.bilibili.com/video/BV1bM411X72E/)。

右端点元素 $x$ 进入窗口后，设窗口内的最大值为 $\textit{mx}$，那么操作次数增加了

$$
\textit{mx}-x
$$

### 左端点元素离开窗口

本题的难点。

例如 $\textit{nums}=[6,3,1,2,4,1,4]$，现在窗口内的数为 $[6,3,1,2,4,1]$，这些数都变成 $6$。如果 $6$ 离开了窗口，那么 $[3,1,2,4,1]$ 会变成 $[3,3,3,4,4]$，每个数的操作次数都会变少。

如果计算操作次数的减少量？换句话说，去掉 $6$ 之后，每个数都变成多少了？

**思考**：去掉 $6$ 之后，哪些数不变（即其最初的值）？哪些数仍然要变大？

**解答**：对于 $3$ 和 $4$ 来说，由于（去掉 $6$ 之后）前面没有它们大的数，所以 $3$ 和 $4$ 不用变。其他数左边仍然有更大的数，所以仍然要变大。

这启发我们引入如下概念。

设 $\textit{left}[i]$ 是 $i$ 左侧最近的大于 $\textit{nums}[i]$ 的数的下标。

把 $\textit{left}[i]$ 和 $i$ 连边，得到一棵树。$\textit{nums}=[6,3,1,2,4,1,4]$ 如下。为方便大家阅读，图中画的是元素值，实际代码中是下标。

![lc3420.png](https://pic.leetcode.cn/1736675203-LzPPhO-lc3420.png)

设 $\textit{out}=6$ 是离开窗口的数，分类讨论：

- 如果 $\textit{out}$ 就是在 $i$ 左侧的大于 $\textit{nums}[i]$ 的数，那么 $\textit{out}$ 离开窗口后，$\textit{nums}[i]$ 无需操作，仍然为 $\textit{nums}[i]$。例如上图中的 $3$ 和 $4$。
- 如果 $\textit{left}[i]$ 在 $\textit{out}$ 的右边，例如上图中的 $1$ 和 $2$，它们左侧大于其的数不是 $\textit{out}$，而是 $3$，所以这两个数操作后都变成了 $3$。

继续思考下去，可以得出如下结论：

- **树中每个节点都要变成其根节点的值。**

$\textit{out}$ 离开窗口，就相当于断开 $\textit{out}$ 及其子节点的边，这会生成三棵更小的树。每棵树中的节点，都要变成其根节点的值。

遍历这三棵 $\textit{out}$ 的子树，计算操作次数的减少量：

- 对于第一棵子树，其中每个节点都要变成 $3$，操作次数都减少了 $6-3=3$，子树内有 $3$ 个节点，所以总的操作次数减少了 $(6-3)\cdot 3 = 9$。
- 对于第二棵子树，其中每个节点都要变成 $4$，操作次数都减少了 $6-4=2$，子树内有 $2$ 个节点，所以总的操作次数减少了 $(6-4)\cdot 2 = 4$。
- 对于第三棵子树，如果元素 $4$ 的下标大于窗口右端点 $r$，结束遍历。否则计算方式同上。

怎么计算树的节点个数？可以 DFS。更简单的做法是，算出每个 $x=\textit{nums}[i]$ 右侧 $\ge x$ 的最近元素下标 $\textit{posR}[i]$（如果不存在则为 $n$）。于是子树节点下标范围为 $[i,\textit{posR}[i])$，子树大小就是区间大小，即

$$
\textit{posR}[i] - i
$$

特殊情况：如果窗口右端点 $r$ 在子树内，则上式改为

$$
r + 1 - i
$$

取最小值，得到最终的子树大小

$$
\min(\textit{posR}[i], r+1) - i
$$

对于该子树，操作次数的减小量为

$$
(\textit{out} - \textit{nums}[i]) \cdot (\min(\textit{posR}[i], r+1) - i)
$$

计算左右最近大于（大于等于）$\textit{nums}[i]$ 的元素下标，可以用 [单调栈【基础算法精讲 26】](https://www.bilibili.com/video/BV1VN411J7S7/)。

> 注：这个建树的思路，类似**笛卡尔树**。

### 子数组个数

滑动窗口的内层循环结束时，右端点**固定**在 $r$，左端点在 $l,l+1,\ldots,r$ 的所有子数组都是合法的，这一共有 $r-l+1$ 个。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1HKcue9ETm/?t=31m06s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countNonDecreasingSubarrays(self, nums: List[int], k: int) -> int:
        n = len(nums)
        g = [[] for _ in range(n)]
        pos_r = [n] * n
        st = []
        for i, x in enumerate(nums):
            while st and x >= nums[st[-1]]:
                pos_r[st[-1]] = i
                st.pop()
            # 循环结束后，栈顶就是左侧 > x 的最近元素了
            if st:
                g[st[-1]].append(i)
            st.append(i)

        ans = cnt = l = 0
        q = deque()  # 单调队列维护最大值
        for r, x in enumerate(nums):
            # x 进入窗口
            while q and nums[q[-1]] <= x:
                q.pop()  # 维护 q 的单调性
            q.append(r)

            # 由于队首到队尾单调递减，所以窗口最大值就是队首
            cnt += nums[q[0]] - x

            # 操作次数太多，缩小窗口
            while cnt > k:
                out = nums[l]  # 离开窗口的元素
                for i in g[l]:
                    if i > r:
                        break
                    cnt -= (out - nums[i]) * (min(pos_r[i], r + 1) - i)
                l += 1

                # 队首已经离开窗口了
                if q[0] < l:
                    q.popleft()

            ans += r - l + 1

        return ans
```

```java [sol-Java]
class Solution {
    public long countNonDecreasingSubarrays(int[] nums, int k) {
        int n = nums.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        int[] posR = new int[n];
        Arrays.fill(posR, n);
        Deque<Integer> st = new ArrayDeque<>();
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            while (!st.isEmpty() && x >= nums[st.peek()]) {
                posR[st.pop()] = i;
            }
            // 循环结束后，栈顶就是左侧 > x 的最近元素了
            if (!st.isEmpty()) {
                g[st.peek()].add(i);
            }
            st.push(i);
        }

        long ans = 0;
        int cnt = 0;
        int l = 0; // 窗口左端点
        Deque<Integer> q = new ArrayDeque<>(); // 单调队列维护最大值
        for (int r = 0; r < n; r++) { // 窗口右端点
            int x = nums[r];
            // x 进入窗口
            while (!q.isEmpty() && nums[q.peekLast()] <= x) {
                q.pollLast(); // 维护 q 的单调性
            }
            q.addLast(r);

            // 由于队首到队尾单调递减，所以窗口最大值就是队首
            cnt += nums[q.peekFirst()] - x;

            // 操作次数太多，缩小窗口
            while (cnt > k) {
                int out = nums[l]; // 离开窗口的元素
                for (int i : g[l]) {
                    if (i > r) {
                        break;
                    }
                    cnt -= (out - nums[i]) * (Math.min(posR[i], r + 1) - i);
                }
                l++;

                // 队首已经离开窗口了
                if (!q.isEmpty() && q.peekFirst() < l) {
                    q.pollFirst();
                }
            }

            ans += r - l + 1;
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countNonDecreasingSubarrays(vector<int>& nums, int k) {
        int n = nums.size();
        vector<vector<int>> g(n);
        vector<int> pos_r(n, n);
        stack<int> st;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            while (!st.empty() && x >= nums[st.top()]) {
                pos_r[st.top()] = i;
                st.pop();
            }
            // 循环结束后，栈顶就是左侧 > x 的最近元素了
            if (!st.empty()) {
                g[st.top()].push_back(i);
            }
            st.push(i);
        }

        long long ans = 0;
        int cnt = 0, l = 0;
        deque<int> q; // 单调队列维护最大值
        for (int r = 0; r < n; r++) {
            int x = nums[r];
            // x 进入窗口
            while (!q.empty() && nums[q.back()] <= x) {
                q.pop_back(); // 维护 q 的单调性
            }
            q.push_back(r);

            // 由于队首到队尾单调递减，所以窗口最大值就是队首
            cnt += nums[q.front()] - x;

            // 操作次数太多，缩小窗口
            while (cnt > k) {
                int out = nums[l]; // 离开窗口的元素
                for (int i : g[l]) {
                    if (i > r) {
                        break;
                    }
                    cnt -= (out - nums[i]) * (min(pos_r[i], r + 1) - i);
                }
                l++;

                // 队首已经离开窗口了
                if (!q.empty() && q.front() < l) {
                    q.pop_front();
                }
            }

            ans += r - l + 1;
        }

        return ans;
    }
};
```

```go [sol-Go]
func countNonDecreasingSubarrays(nums []int, k int) (ans int64) {
	n := len(nums)
	g := make([][]int, n)
	posR := make([]int, n)
	st := []int{}
	for i, x := range nums {
		for len(st) > 0 && x >= nums[st[len(st)-1]] {
			posR[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		// 循环结束后，栈顶就是左侧 > x 的最近元素了
		if len(st) > 0 {
			left := st[len(st)-1]
			g[left] = append(g[left], i)
		}
		st = append(st, i)
	}
	for _, i := range st {
		posR[i] = n
	}

	cnt := 0
	l := 0
	q := []int{} // 单调队列维护最大值
	for r, x := range nums {
		// x 进入窗口
		for len(q) > 0 && nums[q[len(q)-1]] <= x {
			q = q[:len(q)-1] // 维护 q 的单调性
		}
		q = append(q, r)

		// 由于队首到队尾单调递减，所以窗口最大值就是队首
		cnt += nums[q[0]] - x

		// 操作次数太多，缩小窗口
		for cnt > k {
			out := nums[l] // 离开窗口的元素
			for _, i := range g[l] {
				if i > r {
					break
				}
				cnt -= (out - nums[i]) * (min(posR[i], r+1) - i)
			}
			l++

			// 队首已经离开窗口了
			if q[0] < l {
				q = q[1:]
			}
		}

		ans += int64(r - l + 1)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：从右到左滑窗

从右到左滑窗，写起来更加简单。

再来看这棵树。

![lc3420.png](https://pic.leetcode.cn/1736675203-LzPPhO-lc3420.png)

外层循环**倒序**枚举窗口左端点 $l$，内层循环左移窗口右端点 $r$。

例如 $x=\textit{nums}[l]=6$ 进入窗口，那么方法一中的删边就变成了加边，所以操作次数的增加量就是

$$
(x - v)\cdot \textit{sz}
$$

其中 $v$ 是子树根节点的值，$\textit{sz}$ 是子树大小。例如子树 $3$ 的操作次数增加了 $(6-3)\cdot 3=9$。

右端点 $\textit{nums}[r]$ 离开窗口时，操作次数的减少量，等于 $\textit{nums}[r]$ 所属那棵树的根节点的值，减去 $\textit{nums}[r]$。例如上图中的 $4$ 离开窗口，那么操作次数的减少量就是 $6-4=2$。同时，把 $\textit{nums}[r]$ 所属子树的大小减一。

```py [sol-Python3]
class Solution:
    def countNonDecreasingSubarrays(self, nums: List[int], k: int) -> int:
        n = len(nums)
        ans = cnt = 0
        q = deque()  # (根节点的值, 树的大小)
        r = n - 1
        for l in range(n - 1, -1, -1):
            # x 进入窗口
            x = nums[l]
            size = 1  # 统计以 x 为根的树的大小
            while q and x >= q[-1][0]:
                # 以 v 为根的树，现在合并到 x 的下面（x 和 v 连一条边）
                v, sz = q.pop()
                size += sz
                cnt += (x - v) * sz  # 树 v 中的数都变成 x
            q.append([x, size])

            # 操作次数太多，缩小窗口
            while cnt > k:
                # 操作次数的减少量，等于 nums[r] 所在树的根节点值减去 nums[r]
                tree = q[0]  # 最右边的树
                cnt -= tree[0] - nums[r]
                r -= 1
                # nums[r] 离开窗口后，树的大小减一
                tree[1] -= 1
                if tree[1] == 0:  # 这棵树是空的
                    q.popleft()

            ans += r - l + 1

        return ans
```

```java [sol-Java]
class Solution {
    public long countNonDecreasingSubarrays(int[] nums, int k) {
        int n = nums.length;
        long ans = 0;
        long cnt = 0;
        Deque<int[]> q = new ArrayDeque<>(); // (根节点的值, 树的大小)
        int r = n - 1;
        for (int l = n - 1; l >= 0; l--) {
            // x 进入窗口
            int x = nums[l];
            int size = 1; // 统计以 x 为根的树的大小
            while (!q.isEmpty() && x >= q.peekFirst()[0]) {
                // 以 v 为根的树，现在合并到 x 的下面（x 和 v 连一条边）
                int[] tree = q.pollFirst();
                int v = tree[0], sz = tree[1];
                size += sz;
                cnt += (long) (x - v) * sz; // 树 v 中的数都变成 x
            }
            q.addFirst(new int[]{x, size});

            // 操作次数太多，缩小窗口
            while (cnt > k) {
                int[] tree = q.peekLast(); // 最右边的树
                // 操作次数的减少量，等于 nums[r] 所在树的根节点值减去 nums[r]
                cnt -= tree[0] - nums[r];
                r--;
                // nums[r] 离开窗口后，树的大小减一
                tree[1]--;
                if (tree[1] == 0) { // 这棵树是空的
                    q.pollLast();
                }
            }

            ans += r - l + 1;
        }
        return ans;
    }
}
```

```java [sol-Java 写法二]
class Solution {
    public long countNonDecreasingSubarrays(int[] nums, int k) {
        int n = nums.length;
        long ans = 0;
        long cnt = 0;
        int rTree = 0; // rTree 表示窗口最右边那棵树在 q 中的下标
        List<int[]> q = new ArrayList<>(); // (根节点的值, 树的大小)
        int r = n - 1;
        for (int l = n - 1; l >= 0; l--) {
            // x 进入窗口
            int x = nums[l];
            int size = 1; // 统计以 x 为根的树的大小
            while (!q.isEmpty() && x >= q.get(q.size() - 1)[0]) {
                // 以 v 为根的树，现在合并到 x 的下面（x 和 v 连一条边）
                int[] tree = q.remove(q.size() - 1);
                int v = tree[0], sz = tree[1];
                size += sz;
                cnt += (long) (x - v) * sz; // 树 v 中的数都变成 x
            }
            q.add(new int[]{x, size});

            // 如果从 q 中弹出树包含 rTree，那么 rTree 现在指向栈顶这棵树
            rTree = Math.min(rTree, q.size() - 1);

            // 操作次数太多，缩小窗口
            while (cnt > k) {
                int[] tree = q.get(rTree); // 最右边的树
                // 操作次数的减少量，等于 nums[r] 所在树的根节点值减去 nums[r]
                cnt -= tree[0] - nums[r];
                r--;
                // nums[r] 离开窗口后，树的大小减一
                tree[1]--;
                if (tree[1] == 0) { // 这棵树是空的
                    rTree++; // rTree 指向左边下一棵树
                }
            }

            ans += r - l + 1;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countNonDecreasingSubarrays(vector<int>& nums, int k) {
        int n = nums.size();
        long long ans = 0, cnt = 0;
        deque<pair<int, int>> q; // (根节点的值, 树的大小)
        int r = n - 1;
        for (int l = n - 1; l >= 0; l--) {
            // x 进入窗口
            int x = nums[l];
            int size = 1; // 统计以 x 为根的树的大小
            while (!q.empty() && x >= q.back().first) {
                // 以 v 为根的树，现在合并到 x 的下面（x 和 v 连一条边）
                auto [v, sz] = q.back();
                q.pop_back();
                size += sz;
                cnt += (long long) (x - v) * sz; // 树 v 中的数都变成 x
            }
            q.emplace_back(x, size);

            // 操作次数太多，缩小窗口
            while (cnt > k) {
                auto& [v, sz] = q.front(); // 最右边的树（注意这里是引用）
                // 操作次数的减少量，等于 nums[r] 所在树的根节点值减去 nums[r]
                cnt -= v - nums[r];
                r--;
                // nums[r] 离开窗口后，树的大小减一
                sz--;
                if (sz == 0) { // 这棵树是空的
                    q.pop_front();
                }
            }

            ans += r - l + 1;
        }
        return ans;
    }
};
```

```cpp [sol-C++ 写法二]
class Solution {
public:
    long long countNonDecreasingSubarrays(vector<int>& nums, int k) {
        int n = nums.size();
        long long ans = 0, cnt = 0;
        int r_tree = 0; // r_tree 表示窗口最右边那棵树在 q 中的下标
        vector<pair<int, int>> q; // (根节点的值, 树的大小)
        int r = n - 1;
        for (int l = n - 1; l >= 0; l--) {
            // x 进入窗口
            int x = nums[l];
            int size = 1; // 统计以 x 为根的树的大小
            while (!q.empty() && x >= q.back().first) {
                // 以 v 为根的树，现在合并到 x 的下面（x 和 v 连一条边）
                auto [v, sz] = q.back();
                q.pop_back();
                size += sz;
                cnt += (long long) (x - v) * sz; // 树 v 中的数都变成 x
            }
            q.emplace_back(x, size);

            // 如果从 q 中弹出的树包含 r_tree，那么 r_tree 现在指向栈顶这棵树
            r_tree = min(r_tree, (int) q.size() - 1);

            // 操作次数太多，缩小窗口
            while (cnt > k) {
                auto& [v, sz] = q[r_tree]; // 最右边的树（注意这里是引用）
                // 操作次数的减少量，等于 nums[r] 所在树的根节点值减去 nums[r]
                cnt -= v - nums[r];
                r--;
                // nums[r] 离开窗口后，树的大小减一
                sz--;
                if (sz == 0) { // 这棵树是空的
                    r_tree++; // r_tree 指向左边下一棵树
                }
            }

            ans += r - l + 1;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countNonDecreasingSubarrays(nums []int, k int) (ans int64) {
	n := len(nums)
	cnt := 0
	type pair struct{ val, size int } // 根节点的值, 树的大小
	q := []pair{}
	r := n - 1
	for l, x := range slices.Backward(nums) {
		// x 进入窗口
		size := 1 // 统计以 x 为根的树的大小
		for len(q) > 0 && x >= q[len(q)-1].val {
			// 以 p.val 为根的树，现在合并到 x 的下面（x 和 val 连一条边）
			p := q[len(q)-1]
			q = q[:len(q)-1]
			size += p.size
			cnt += (x - p.val) * p.size // 树 p.val 中的数都变成 x
		}
		q = append(q, pair{x, size})

		// 操作次数太多，缩小窗口
		for cnt > k {
			p := &q[0] // 最右边的树
			// 操作次数的减少量，等于 nums[r] 所在树的根节点值减去 nums[r]
			cnt -= p.val - nums[r]
			r--
			// nums[r] 离开窗口后，树的大小减一
			p.size--
			if p.size == 0 { // 这棵树是空的
				q = q[1:]
			}
		}

		ans += int64(r - l + 1)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

所有子数组的最小操作次数之和是多少？

> 提示：贡献法。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. 【本题相关】[滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. 【本题相关】[单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
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
