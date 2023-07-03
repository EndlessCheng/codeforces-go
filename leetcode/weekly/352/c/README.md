代码框架是 [滑动窗口（双指针）](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。在遍历数组的同时，维护窗口内的数字。

由于绝对差至多为 $2$，所以用平衡树或者哈希表维护都行，反正至多维护 $3$ 个数，添加删除可以视作是 $\mathcal{O}(1)$ 的。（如果用哈希表，还需记录数字的出现次数。）

如果窗口内的最大值与最小值的差大于 $2$，就不断移动左端点 $\textit{left}$，减少窗口内的数字。

最后

$$
[\textit{left},\textit{right}],[\textit{left}+1,\textit{right}],\cdots,[\textit{right},\textit{right}]
$$

这一共 $\textit{right}-\textit{left}+1$ 个子数组都是合法的，加入答案。

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
                if cnt[y] == 0: del cnt[y]
                left += 1
            ans += right - left + 1
        return ans
```

```java [sol-Java]
class Solution {
    public long continuousSubarrays(int[] nums) {
        long ans = 0;
        var t = new TreeMap<Integer, Integer>();
        int left = 0;
        for (int right = 0; right < nums.length; right++) {
            t.merge(nums[right], 1, Integer::sum);
            while (t.lastKey() - t.firstKey() > 2) {
                int y = nums[left++];
                if (t.get(y) == 1) t.remove(y);
                else t.merge(y, -1, Integer::sum);
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
    long long continuousSubarrays(vector<int> &nums) {
        long long ans = 0;
        multiset<int> s;
        int left = 0, n = nums.size();
        for (int right = 0; right < n; right++) {
            s.insert(nums[right]);
            while (*s.rbegin() - *s.begin() > 2)
                s.erase(s.find(nums[left++]));
            ans += right - left + 1;
        }
        return ans;
    }
};
```

```go [sol-Go]
func continuousSubarrays(a []int) (ans int64) {
	cnt := map[int]int{}
	left := 0
	for right, x := range a {
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
			y := a[left]
			if cnt[y]--; cnt[y] == 0 {
				delete(cnt, y)
			}
			left++
		}
		ans += int64(right - left + 1)
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。注意至多维护 $3$ 个数，仅用到常量额外空间。

#### 相似题目

- [1438. 绝对差不超过限制的最长连续子数组](https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/)
