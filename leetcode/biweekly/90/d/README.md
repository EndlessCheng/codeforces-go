下午两点在B站讲这场双周赛的题目，[欢迎关注](https://space.bilibili.com/206214)~

---

**前置知识**：[496. 下一个更大元素 I](https://leetcode.cn/problems/next-greater-element-i/)。

从左往右遍历 $\textit{nums}$，用（递减）单调栈 $s$ 记录元素，如果 $x=\textit{nums}[i]$ 比 $s$ 的栈顶大，则 $x$ 是栈顶的**下个**更大元素，弹出栈顶。最后把 $x$ 入栈（实际入栈的是下标 $i$）。

把弹出的元素加到另一个栈 $t$ 中（注意保持原始顺序），后续循环时，如果 $y=\textit{nums}[j]$ 比 $t$ 的栈顶大，则 $y$ 是栈顶的**下下个**更大元素，记录答案，弹出栈顶。

```py [sol1-Python3]
class Solution:
    def secondGreaterElement(self, nums: List[int]) -> List[int]:
        ans, s, t = [-1] * len(nums), [], []
        for i, x in enumerate(nums):
            while t and nums[t[-1]] < x:
                ans[t.pop()] = x
            j = len(s) - 1
            while j >= 0 and nums[s[j]] < x:
                j -= 1
            t += s[j + 1:]
            del s[j + 1:]
            s.append(i)
        return ans
```

```cpp [sol1-C++]
class Solution {
public:
    vector<int> secondGreaterElement(vector<int> &nums) {
        int n = nums.size();
        vector<int> ans(n, -1), s, t;
        for (int i = 0; i < n; ++i) {
            int x = nums[i];
            while (!t.empty() && nums[t.back()] < x) {
                ans[t.back()] = x;
                t.pop_back();
            }
            int j = (int) s.size() - 1;
            while (j >= 0 && nums[s[j]] < x) --j;
            t.insert(t.end(), s.begin() + (j + 1), s.end());
            s.resize(j + 1);
            s.push_back(i);
        }
        return ans;
    }
};
```

```go [sol1-Go]
func secondGreaterElement(nums []int) []int {
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = -1
	}
	s, t := []int{}, []int{}
	for i, x := range nums {
		for len(t) > 0 && nums[t[len(t)-1]] < x {
			ans[t[len(t)-1]] = x
			t = t[:len(t)-1]
		}
		j := len(s) - 1
		for j >= 0 && nums[s[j]] < x {
			j--
		}
		t = append(t, s[j+1:]...)
		s = append(s[:j+1], i)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。每个元素至多入栈出栈两次。
- 空间复杂度：$O(n)$。
