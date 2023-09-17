更详细的思考过程梳理，请看 [视频讲解](https://b23.tv/PDz9NBA) 第二题。

为了方便判断，把 $\textit{nums}$ 从小到大排序。

如果 $\textit{nums}[0] > 0$，那么所有 $\textit{nums}[i]$ 都是大于 $0$ 的，我们可以一个学生都不选。

如果 $\textit{nums}[i] < i+1 < \textit{nums}[i+1]$，这意味着选择 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 这一共 $i+1$ 个学生，是满足要求的：由于数组已经排好序，$\textit{nums}[0]$ 到 $\textit{nums}[i]$ 都是小于 $i+1$ 的，而 $\textit{nums}[i+1]$ 到 $\textit{nums}[n-1]$ 都是大于 $i+1$ 的。

特别地，如果 $i=n-1$，我们只需要判断是否满足 $\textit{nums}[i] < n$，在题目约束下，这是一定可以满足的。所以最后把答案加一。

```py [sol-Python3]
class Solution:
    def countWays(self, nums: List[int]) -> int:
        nums.sort()
        n = len(nums)
        ans = nums[0] > 0
        for i, (x, y) in enumerate(pairwise(nums)):
            if x < i + 1 < y:
                ans += 1
        return ans + 1  
```

```java [sol-Java]
class Solution {
    public int countWays(List<Integer> nums) {
        int[] a = nums.stream().mapToInt(i -> i).toArray();
        Arrays.sort(a);
        int n = a.length;
        int ans = a[0] > 0 ? 1 : 0;
        for (int i = 0; i < n - 1; i++) {
            if (a[i] < i + 1 && i + 1 < a[i + 1]) {
                ans++;
            }
        }
        return ans + 1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countWays(vector<int> &nums) {
        sort(nums.begin(), nums.end());
        int n = nums.size();
        int ans = nums[0] > 0;
        for (int i = 0; i < n - 1; i++) {
            if (nums[i] < i + 1 && i + 1 < nums[i + 1]) {
                ans++;
            }
        }
        return ans + 1;
    }
};
```

```go [sol-Go]
func countWays(nums []int) (ans int) {
	sort.Ints(nums)
	if nums[0] > 0 {
		ans++
	}
	for i, x := range nums {
		if x < i+1 && (i == len(nums)-1 || i+1 < nums[i+1]) {
			ans++
		}
	}
	return
}
```

```js [sol-JavaScript]
var countWays = function (nums) {
    nums.sort((a, b) => a - b);
    const n = nums.length;
    let ans = nums[0] > 0 ? 1 : 0;
    for (let i = 0; i < n - 1; i++) {
        if (nums[i] < i + 1 && i + 1 < nums[i + 1]) {
            ans++;
        }
    }
    return ans + 1;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

注：如果采用桶排序，可以做到 $\mathcal{O}(n)$ 的时间复杂度。
