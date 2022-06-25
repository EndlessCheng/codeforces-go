本题 [视频讲解](https://www.bilibili.com/video/BV1iF41157dG/) 已出炉，欢迎三连~

---

#### 提示 1

元素 $x$ 会被左边某个比他大的元素 $y$ 给删除（如果存在的话）。

我们需要计算在删除 $x$ 之前，删除了多少个比 $y$ 小的元素，从而算出删除 $x$ 的时刻（第几步操作）。

答案可以转换成所有元素被删除的时刻的最大值。

#### 提示 2

以 $[20,1,9,1,2,3]$ 为例。

- 时刻一 $20$ 删掉 $1$，$9$ 删掉 $1$；
- 时刻二 $20$ 删掉 $9$，$9$ 删掉 $2$;
- 时刻三 $20$ 接替了 $9$ 的任务，来删除数字 $3$。

虽然说数字 $3$ 是被 $20$ 删除的，但是由于 $20$ 立马接替了 $9$，我们可以**等价转换**成 $3$ 是被 $9$ 删除的，也就是它左边离它最近且比它大的那个数。这一等价转换不会影响数字被删除的时刻。

#### 提示 3

再考虑这个例子 $[9,1,2,3,4,1,5]$。

$5$ 应该被 $9$ 删除。根据题目要求，在删除 $5$ 之前，需要把 $5$ 前面不超过 $5$ 的元素都删除，然后才能删除 $5$。所以在删除 $5$ 之前，我们需要知道 $9$ 到 $5$ 之间的所有元素被删除的时刻的最大值，这个时刻加一就是删除 $5$ 的时刻。

这可以用单调栈 + 线段树来做，单调栈求左边最近更大元素位置，线段树维护区间最大值。（[评论区](https://leetcode.cn/problems/steps-to-make-array-non-decreasing/comments/1587279)有人实现了这一思路）

但还有更巧妙的做法。

#### 提示 4

对于一串非降的序列，该序列每个元素被删除的时刻是单调递增的。（假设序列左侧有个更大的元素去删除序列中的元素）

利用这一单调性，我们只需要存储这串非降序列的**最后一个元素**被删除的时刻，提示 3 中所需要计算的最大值必然在这些元素中。

#### 提示 5

我们可以用一个**单调递减栈**存储元素及其被删除的时刻，当遇到一个不小于栈顶的元素 $x$ 时，就不断弹出栈顶元素，并取弹出元素被删除时刻的最大值，这样就得到了提示 3 中所需要计算的时刻的最大值 $\textit{maxT}$。

然后将 $x$ 及 $\textit{maxT}+1$ 入栈。注意如果此时栈为空，说明前面没有比 $x$ 大的元素，$x$ 无法被删除，即 $\textit{maxT}=0$，这种情况需要将 $x$ 及 $0$ 入栈。

#### 复杂度分析

- 时间复杂度：$O(n)$。每个元素至多入栈出栈各一次。
- 空间复杂度：$O(n)$。最坏情况下栈中有 $n$ 个元素。

```Python [sol1-Python3]
class Solution:
    def totalSteps(self, nums: List[int]) -> int:
        ans, st = 0, []
        for num in nums:
            max_t = 0
            while st and st[-1][0] <= num:
                max_t = max(max_t, st.pop()[1])
            max_t = max_t + 1 if st else 0
            ans = max(ans, max_t)
            st.append((num, max_t))
        return ans
```

```java [sol1-Java]
class Solution {
    public int totalSteps(int[] nums) {
        var ans = 0;
        var st = new ArrayDeque<int[]>();
        for (var num : nums) {
            var maxT = 0;
            while (!st.isEmpty() && st.peek()[0] <= num)
                maxT = Math.max(maxT, st.pop()[1]);
            maxT = st.isEmpty() ? 0 : maxT + 1;
            ans = Math.max(ans, maxT);
            st.push(new int[]{num, maxT});
        }
        return ans;
    }
}
```

```C++ [sol1-C++]
class Solution {
public:
    int totalSteps(vector<int> &nums) {
        int ans = 0;
        stack<pair<int, int>> st;
        for (int num : nums) {
            int maxT = 0;
            while (!st.empty() && st.top().first <= num) {
                maxT = max(maxT, st.top().second);
                st.pop();
            }
            maxT = st.empty() ? 0 : maxT + 1;
            ans = max(ans, maxT);
            st.emplace(num, maxT);
        }
        return ans;
    }
};
```

```go [sol1-Go]
func totalSteps(nums []int) (ans int) {
	type pair struct{ v, t int }
	st := []pair{}
	for _, num := range nums {
		maxT := 0
		for len(st) > 0 && st[len(st)-1].v <= num {
			maxT = max(maxT, st[len(st)-1].t)
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			maxT++
			ans = max(ans, maxT)
		} else {
			maxT = 0
		}
		st = append(st, pair{num, maxT})
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```
