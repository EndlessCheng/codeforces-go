## 方法一：滑动窗口+有序集合/哈希表

**前置知识**：[滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

在遍历数组的同时，维护窗口内的元素及其出现次数。

由于绝对差至多为 $2$，所以用有序集合或者哈希表维护都行。

由于至多维护 $3$ 个数，无论用有序集合还是哈希表，添加、删除和查询最值都是 $\mathcal{O}(1)$ 的。

如果窗口内的最大值与最小值的差大于 $2$，就不断移动左端点 $\textit{left}$，减少窗口内的数字。

内层循环结束后，$[\textit{left},\textit{right}]$ 这个子数组是满足题目要求的。由于子数组越短，越能满足题目要求，所以除了 $[\textit{left},\textit{right}]$，还有 $[\textit{left}+1,\textit{right}],[\textit{left}+2,\textit{right}],\ldots,[\textit{right},\textit{right}]$ 都是满足要求的。也就是说，当右端点**固定**在 $\textit{right}$ 时，左端点在 $\textit{left},\textit{left}+1,\textit{left}+2,\ldots,\textit{right}$ 的所有子数组都是满足要求的，这一共有 $\textit{right}-\textit{left}+1$ 个，加入答案。

```py [sol-Python3]
class Solution:
    def continuousSubarrays(self, nums: List[int]) -> int:
        ans = left = 0
        cnt = Counter()
        for right, x in enumerate(nums):
            cnt[x] += 1
            while max(cnt) - min(cnt) > 2:
                y = nums[left]
                cnt[y] -= 1
                if cnt[y] == 0:
                    del cnt[y]
                left += 1
            ans += right - left + 1
        return ans
```

```java [sol-Java]
class Solution {
    public long continuousSubarrays(int[] nums) {
        long ans = 0;
        TreeMap<Integer, Integer> t = new TreeMap<>();
        int left = 0;
        for (int right = 0; right < nums.length; right++) {
            t.merge(nums[right], 1, Integer::sum); // t[nums[right]]++
            while (t.lastKey() - t.firstKey() > 2) {
                int out = nums[left];
                int c = t.get(out);
                if (c == 1) {
                    t.remove(out);
                } else {
                    t.put(out, c - 1);
                }
                left++;
            }
            ans += right - left + 1;
        }
        return ans;
    }
}
```

```cpp [sol-C++ map]
class Solution {
public:
    long long continuousSubarrays(vector<int>& nums) {
        long long ans = 0;
        map<int, int> cnt;
        int left = 0;
        for (int right = 0; right < nums.size(); right++) {
            cnt[nums[right]]++;
            while (cnt.rbegin()->first - cnt.begin()->first > 2) {
                int y = nums[left];
                if (--cnt[y] == 0) {
                    cnt.erase(y);
                }
                left++;
            }
            ans += right - left + 1;
        }
        return ans;
    }
};
```

```cpp [sol-C++ multiset]
class Solution {
public:
    long long continuousSubarrays(vector<int>& nums) {
        long long ans = 0;
        multiset<int> s;
        int left = 0, n = nums.size();
        for (int right = 0; right < n; right++) {
            s.insert(nums[right]);
            while (*s.rbegin() - *s.begin() > 2) {
                s.erase(s.find(nums[left])); // 删除一个 nums[left]
                left++;
            }
            ans += right - left + 1;
        }
        return ans;
    }
};
```

```go [sol-Go]
func continuousSubarrays(nums []int) (ans int64) {
	cnt := map[int]int{}
	left := 0
	for right, x := range nums {
		cnt[x]++
		for {
			mx, mn := x, x
			for k := range cnt {
				mx = max(mx, k)
				mn = min(mn, k)
			}
			if mx-mn <= 2 {
				break
			}
			out := nums[left]
			cnt[out]--
			if cnt[out] == 0 {
				delete(cnt, out)
			}
			left++
		}
		ans += int64(right - left + 1)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log D)$ 或 $\mathcal{O}(nD)$，其中 $n$ 为 $\textit{nums}$ 的长度，$D=2$ 表示最大值与最小值之差的上限。
- 空间复杂度：$\mathcal{O}(D)$。

## 方法二：滑动窗口+单调队列

可以做到和 $D$ 无关，且时间复杂度为 $\mathcal{O}(n)$。

**前置知识**：[单调队列【基础算法精讲 27】](https://www.bilibili.com/video/BV1bM411X72E/)

本质是计算滑动窗口的最小值和最大值，可以用两个单调队列 $\textit{minQ}$ 和 $\textit{maxQ}$ 分别维护窗口中的最小值（的下标）和最大值（的下标）。

如果最大值减去最小值大于 $2$，那么把 $\textit{left}$ 加一。然后检查队首是否在窗口外，如果在窗口外，就移出队首。

```py [sol-Python3]
class Solution:
    def continuousSubarrays(self, nums: List[int]) -> int:
        min_q = deque()
        max_q = deque()
        ans = left = 0
        for right, x in enumerate(nums):
            while min_q and x <= nums[min_q[-1]]:
                min_q.pop()
            min_q.append(right)

            while max_q and x >= nums[max_q[-1]]:
                max_q.pop()
            max_q.append(right)

            while nums[max_q[0]] - nums[min_q[0]] > 2:
                left += 1
                if min_q[0] < left:
                    min_q.popleft()
                if max_q[0] < left:
                    max_q.popleft()
            ans += right - left + 1
        return ans
```

```java [sol-Java]
class Solution {
    public long continuousSubarrays(int[] nums) {
        Deque<Integer> minQ = new ArrayDeque<>();
        Deque<Integer> maxQ = new ArrayDeque<>();
        long ans = 0;
        int left = 0;
        for (int right = 0; right < nums.length; right++) {
            int x = nums[right];
            while (!minQ.isEmpty() && x <= nums[minQ.peekLast()]) {
                minQ.pollLast();
            }
            minQ.addLast(right);

            while (!maxQ.isEmpty() && x >= nums[maxQ.peekLast()]) {
                maxQ.pollLast();
            }
            maxQ.addLast(right);

            while (nums[maxQ.peekFirst()] - nums[minQ.peekFirst()] > 2) {
                left++;
                if (minQ.peekFirst() < left) {
                    minQ.pollFirst();
                }
                if (maxQ.peekFirst() < left) {
                    maxQ.pollFirst();
                }
            }
            ans += right - left + 1;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long continuousSubarrays(vector<int>& nums) {
        deque<int> min_q, max_q;
        long long ans = 0;
        int left = 0;
        for (int right = 0; right < nums.size(); right++) {
            int x = nums[right];
            while (!min_q.empty() && x <= nums[min_q.back()]) {
                min_q.pop_back();
            }
            min_q.push_back(right);

            while (!max_q.empty() && x >= nums[max_q.back()]) {
                max_q.pop_back();
            }
            max_q.push_back(right);

            while (nums[max_q.front()] - nums[min_q.front()] > 2) {
                left++;
                if (min_q.front() < left) {
                    min_q.pop_front();
                }
                if (max_q.front() < left) {
                    max_q.pop_front();
                }
            }
            ans += right - left + 1;
        }
        return ans;
    }
};
```

```go [sol-Go]
func continuousSubarrays(nums []int) (ans int64) {
	var minQ, maxQ []int
	left := 0
	for right, x := range nums {
		for len(minQ) > 0 && x <= nums[minQ[len(minQ)-1]] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, right)
		
		for len(maxQ) > 0 && x >= nums[maxQ[len(maxQ)-1]] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, right)

		for nums[maxQ[0]]-nums[minQ[0]] > 2 {
			left++
			if minQ[0] < left {
				minQ = minQ[1:]
			}
			if maxQ[0] < left {
				maxQ = maxQ[1:]
			}
		}
		ans += int64(right - left + 1)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

- [1438. 绝对差不超过限制的最长连续子数组](https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/)

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
