从易到难，先思考原问题的一个简单版本：

- 只统计满足 $\textit{nums}[l] \ge \textit{nums}[r]$ 的合法子数组的数量。

此时 $\min(\textit{nums}[l], \textit{nums}[r]) = \textit{nums}[r]$。

合法子数组必须满足

$$
\textit{nums}[r] > \max(\textit{nums}[l+1],\ldots, \textit{nums}[r-1])
$$

上式表明，$[l+1,r-1]$ 中的所有数都小于 $\textit{nums}[r]$。

换句话说，$\textit{nums}[l]$ 是 $\textit{nums}[r]$ 左侧**最近**的大于等于 $\textit{nums}[r]$ 的数。

为什么是最近？反证法，如果继续向左，让子数组包含更多元素，那么 $\textit{nums}[l]$ 就会变成子数组的中间元素，由于 $\textit{nums}[l] \ge \textit{nums}[r]$，不符合要求。

所以对于每个右端点 $r$，我们只需要找 $\textit{nums}[r]$ 左侧最近的大于等于 $\textit{nums}[r]$ 的数的下标 $l$。如果 $l$ 存在且 $i-l+1\ge 3$，那么找到了一个合法子数组，把答案加一。

这是**单调栈**的标准应用，请看 [单调栈【基础算法精讲 26】](https://www.bilibili.com/video/BV1VN411J7S7/)。

对于 $\textit{nums}[l] < \textit{nums}[r]$ 的情况，同样可以用单调栈计算。

⚠**注意**：小心子数组两端点元素相同的情况，不能重复统计。不妨规定，左侧找大于等于，右侧找严格大于。虽然本题保证「所有元素互不相同」，但我的做法也适用于有重复元素的情况。

> 此外，上述结论表明，答案的上界是 $2n$（粗略估计），所以返回值用 $\texttt{int}$ 就够了。本题没有重复元素，答案的上界是 $n$（见写法二）。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1heYGzWEUa/?t=7m35s)，欢迎点赞关注~

## 写法一

```py [sol-Python3]
class Solution:
    def bowlSubarrays(self, nums: List[int]) -> int:
        ans = 0
        st = []
        for i, x in enumerate(nums):
            while st and nums[st[-1]] < x:
                # j=st[-1] 右侧严格大于 nums[j] 的数的下标是 i
                if i - st.pop() > 1:  # 子数组的长度至少为 3
                    ans += 1
            # i 左侧大于等于 nums[i] 的数的下标是 st[-1]
            if st and i - st[-1] > 1:  # 子数组的长度至少为 3
                ans += 1
            st.append(i)
        return ans
```

```java [sol-Java]
class Solution {
    public long bowlSubarrays(int[] nums) {
        int ans = 0;
        Deque<Integer> st = new ArrayDeque<>(); // 更快的写法见【Java 数组】
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            while (!st.isEmpty() && nums[st.peek()] < x) {
                // j=st.peek() 右侧严格大于 nums[j] 的数的下标是 i
                if (i - st.pop() > 1) { // 子数组的长度至少为 3
                    ans++;
                }
            }
            // i 左侧大于等于 nums[i] 的数的下标是 st.peek()
            if (!st.isEmpty() && i - st.peek() > 1) { // 子数组的长度至少为 3
                ans++;
            }
            st.push(i);
        }
        return ans;
    }
}
```

```java [sol-Java 数组]
class Solution {
    public long bowlSubarrays(int[] nums) {
        int n = nums.length;
        long ans = 0;
        int[] st = new int[n]; // 模拟栈
        int top = -1; // 栈顶下标
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            while (top >= 0 && nums[st[top]] < x) {
                // j=st[top] 右侧严格大于 nums[j] 的数的下标是 i
                if (i - st[top] > 1) { // 子数组的长度至少为 3
                    ans++;
                }
                top--; // 出栈
            }
            // i 左侧大于等于 nums[i] 的数的下标是 st[top]
            if (top >= 0 && i - st[top] > 1) { // 子数组的长度至少为 3
                ans++;
            }
            st[++top] = i; // 入栈
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long bowlSubarrays(vector<int>& nums) {
        int ans = 0;
        stack<int> st;
        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];
            while (!st.empty() && nums[st.top()] < x) {
                // j=st.top() 右侧严格大于 nums[j] 的数的下标是 i
                if (i - st.top() > 1) { // 子数组的长度至少为 3
                    ans++;
                }
                st.pop();
            }
            // i 左侧大于等于 nums[i] 的数的下标是 st.top()
            if (!st.empty() && i - st.top() > 1) { // 子数组的长度至少为 3
                ans++;
            }
            st.push(i);
        }
        return ans;
    }
};
```

```go [sol-Go]
func bowlSubarrays(nums []int) (ans int64) {
	st := []int{}
	for i, x := range nums {
		for len(st) > 0 && nums[st[len(st)-1]] < x {
			// j=st[len(st)-1] 右侧严格大于 nums[j] 的数的下标是 i
			if i-st[len(st)-1] > 1 { // 子数组的长度至少为 3
				ans++
			}
			st = st[:len(st)-1]
		}
		// i 左侧大于等于 nums[i] 的数的下标是 st[len(st)-1]
		if len(st) > 0 && i-st[len(st)-1] > 1 { // 子数组的长度至少为 3
			ans++
		}
		st = append(st, i)
	}
	return
}
```

## 写法二

注意到，无论是左大右小（比如 $[3,1,2]$）还是左小右大（比如 $[2,1,3]$），当我们遍历到 $\textit{nums}[i]$ 时，如果出栈后栈不为空，说明栈顶和 $\textit{nums}[i]$ 是合法子数组的左右端点。「栈不为空」这个条件还说明子数组的长度至少是 $3$。

比如 $[5,3,2,1,4]$，遍历到 $4$ 时：

- 弹出 $1$，发现栈不为空，那么 $[2,1,4]$ 是符合要求的子数组。
- 弹出 $2$，发现栈不为空，那么 $[3,2,1,4]$ 是符合要求的子数组。
- 弹出 $3$，发现栈不为空，那么 $[5,3,2,1,4]$ 是符合要求的子数组。

注意本题保证所有元素互不相同。

```py [sol-Python3]
class Solution:
    def bowlSubarrays(self, nums: List[int]) -> int:
        ans = 0
        st = []
        for x in nums:
            while st and st[-1] < x:
                st.pop()
                if st:
                    ans += 1
            st.append(x)
        return ans
```

```java [sol-Java]
class Solution {
    public long bowlSubarrays(int[] nums) {
        int ans = 0;
        Deque<Integer> st = new ArrayDeque<>(); // 更快的写法见【Java 数组】
        for (int x : nums) {
            while (!st.isEmpty() && st.peek() < x) {
                st.pop();
                if (!st.isEmpty()) {
                    ans++;
                }
            }
            st.push(x);
        }
        return ans;
    }
}
```

```java [sol-Java 数组]
class Solution {
    public long bowlSubarrays(int[] nums) {
        long ans = 0;
        int[] st = new int[nums.length]; // 模拟栈
        int top = -1; // 栈顶下标
        for (int x : nums) {
            while (top >= 0 && st[top] < x) {
                top--; // 出栈
                if (top >= 0) {
                    ans++;
                }
            }
            st[++top] = x; // 入栈
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long bowlSubarrays(vector<int>& nums) {
        int ans = 0;
        stack<int> st;
        for (int x : nums) {
            while (!st.empty() && st.top() < x) {
                st.pop();
                if (!st.empty()) {
                    ans++;
                }
            }
            st.push(x);
        }
        return ans;
    }
};
```

```go [sol-Go]
func bowlSubarrays(nums []int) (ans int64) {
	st := []int{}
	for _, x := range nums {
		for len(st) > 0 && st[len(st)-1] < x {
			st = st[:len(st)-1]
			if len(st) > 0 {
				ans++
			}
		}
		st = append(st, x)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。虽然写了个二重循环，但站在每个元素的视角看，这个元素在二重循环中最多入栈出栈各一次，因此整个二重循环的总循环次数为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 空间优化

$\textit{nums}$ 遍历过的数不会再用到，可以充分利用遍历过的空间，把 $\textit{nums}$ 当作栈。

```py [sol-Python3]
class Solution:
    def bowlSubarrays(self, nums: List[int]) -> int:
        ans = 0
        top = -1  # 栈顶下标
        for x in nums:
            while top >= 0 and nums[top] < x:
                top -= 1  # 出栈
                if top >= 0:
                    ans += 1
            top += 1
            nums[top] = x  # 入栈
        return ans
```

```java [sol-Java]
class Solution {
    public long bowlSubarrays(int[] nums) {
        long ans = 0;
        int top = -1; // 栈顶下标
        for (int x : nums) {
            while (top >= 0 && nums[top] < x) {
                top--; // 出栈
                if (top >= 0) {
                    ans++;
                }
            }
            nums[++top] = x; // 入栈
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long bowlSubarrays(vector<int>& nums) {
        int ans = 0;
        int top = -1; // 栈顶下标
        for (int x : nums) {
            while (top >= 0 && nums[top] < x) {
                top--; // 出栈
                if (top >= 0) {
                    ans++;
                }
            }
            nums[++top] = x; // 入栈
        }
        return ans;
    }
};
```

```go [sol-Go]
func bowlSubarrays(nums []int) (ans int64) {
	st := nums[:0]
	for _, x := range nums {
		for len(st) > 0 && st[len(st)-1] < x {
			st = st[:len(st)-1]
			if len(st) > 0 {
				ans++
			}
		}
		st = append(st, x)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。虽然写了个二重循环，但站在每个元素的视角看，这个元素在二重循环中最多入栈出栈各一次，因此整个二重循环的总循环次数为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

本题要求子数组两个端点都大于中间元素，如果改成**至少**有一个端点大于中间元素，怎么做？

欢迎在评论区分享你的思路/代码。

## 专题训练

见下面单调栈题单的第一章。

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
