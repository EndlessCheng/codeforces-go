#### 提示 1

元素 $x$ 会被左边某个比他大的元素 $y$ 给移除（如果存在的话）。

我们需要计算在移除 $x$ 之前，移除了多少个比 $y$ 小的元素，从而算出移除 $x$ 的时间（第几步操作）。

#### 提示 2

注意移除操作是可以「接力」的，例如 $[9,1,2,3,4,1,5]$，$5$ 应该被 $9$ 删掉，所以在删除 $5$ 之前，我们需要知道 $9$ 到 $5$ 之间的所有元素被移除的时间的最大值，从而算出移除 $5$ 的时间。

这可以用单调栈 + 线段树来做，但还有更巧妙的办法。

#### 提示 3

对于一串连续非降的序列，每个元素被删除的时间是单调递增的。

这意味着我们只需要存储「关键元素」被删除的时间。

#### 提示 4

我们可以用一个**单调递减栈**存储每个元素及其被删除的时间，当遇到一个不小于栈顶的元素 $x$ 时，就不断弹出栈顶元素，并取弹出元素被删除时间的最大值，这样就得到了提示 2 中所需要计算的时间的最大值 $\textit{maxT}$。

然后将 $x$ 及 $\textit{maxT}+1$ 入栈（如果此时栈为空，则说明前面没有比 $x$ 大的元素，此时 $\textit{maxT}=0$，我们可以将 $x$ 及 $0$ 入栈）。

#### 复杂度分析

- 时间复杂度：$O(n)$。
- 空间复杂度：$O(n)$。

```Python [sol1-Python3]
class Solution:
    def totalSteps(self, nums: List[int]) -> int:
        ans, st = 0, []
        for num in nums:
            max_t = 0
            while st and st[-1][0] <= num:
                max_t = max(max_t, st.pop()[1])
            if st: max_t += 1
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
            var max_t = 0;
            while (!st.isEmpty() && st.peek()[0] <= num) {
                max_t = Math.max(max_t, st.peek()[1]);
                st.pop();
            }
            if (!st.isEmpty()) ++max_t;
            ans = Math.max(ans, max_t);
            st.push(new int[]{num, max_t});
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
            int max_t = 0;
            while (!st.empty() && st.top().first <= num) {
                max_t = max(max_t, st.top().second);
                st.pop();
            }
            if (!st.empty()) ++max_t;
            ans = max(ans, max_t);
            st.emplace(num, max_t);
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
		}
		ans = max(ans, maxT)
		st = append(st, pair{num, maxT})
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```
