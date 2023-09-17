[视频讲解](https://b23.tv/PDz9NBA) 第二题。

在选择学生人数固定的时候，选择方案是否唯一呢？

是的。比如选了 $k$ 个学生，那么所有 $\textit{nums}[i] < k$ 的学生都要选，所有 $\textit{nums}[i] > k$ 的都不能选，并且不能出现 $\textit{nums}[i] = k$ 的情况。这意味着**在选择学生人数固定的时候，选择方案是唯一的**。为方便判断，可以把 $\textit{nums}$ 从小到大排序。

既然「所有 $\textit{nums}[i] < k$ 的学生都要选」，那么我们必须选 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 这一共 $i+1$ 个学生。

由于数组已经排好序，$\textit{nums}[0]$ 到 $\textit{nums}[i]$ 都必须小于 $i+1$，且 $\textit{nums}[i+1]$ 到 $\textit{nums}[n-1]$ 都必须大于 $i+1$。

由于 $\textit{nums}[i]$ 都大于等于它左边的数，$\textit{nums}[i+1]$ 都小于等于它右边的数，所以只需要判断

$$
\textit{nums}[i] < i+1 < \textit{nums}[i+1]
$$

上式成立就意味着我们可以选 $i+1$ 个学生，算作一种方案。

特别地，如果 $i=n-1$，我们只需要判断是否满足 $\textit{nums}[i] < n$，在题目约束下，这是一定可以满足的。所以最后把答案加一。

此外，如果 $\textit{nums}[0] > 0$，那么所有 $\textit{nums}[i]$ 都是大于 $0$ 的，我们可以一个学生都不选。

```py [sol-Python3]
class Solution:
    def countWays(self, nums: List[int]) -> int:
        nums.sort()
        n = len(nums)
        ans = nums[0] > 0  # 一个学生都不选
        for i, (x, y) in enumerate(pairwise(nums)):
            if x < i + 1 < y:
                ans += 1
        return ans + 1  # +1 是因为可以都选
```

```java [sol-Java]
class Solution {
    public int countWays(List<Integer> nums) {
        int[] a = nums.stream().mapToInt(i -> i).toArray();
        Arrays.sort(a);
        int n = a.length;
        int ans = a[0] > 0 ? 1 : 0; // 一个学生都不选
        for (int i = 0; i < n - 1; i++) {
            if (a[i] < i + 1 && i + 1 < a[i + 1]) {
                ans++;
            }
        }
        return ans + 1; // +1 是因为可以都选
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countWays(vector<int> &nums) {
        sort(nums.begin(), nums.end());
        int n = nums.size();
        int ans = nums[0] > 0; // 一个学生都不选
        for (int i = 0; i < n - 1; i++) {
            if (nums[i] < i + 1 && i + 1 < nums[i + 1]) {
                ans++;
            }
        }
        return ans + 1; // +1 是因为可以都选
    }
};
```

```go [sol-Go]
func countWays(nums []int) (ans int) {
	sort.Ints(nums)
	if nums[0] > 0 { // 一个学生都不选
		ans++
	}
	for i, x := range nums[:len(nums)-1] {
		if x < i+1 && i+1 < nums[i+1] {
			ans++
		}
	}
	return ans + 1 // +1 是因为可以都选
}
```

```js [sol-JavaScript]
var countWays = function (nums) {
    nums.sort((a, b) => a - b);
    const n = nums.length;
    let ans = nums[0] > 0 ? 1 : 0; // 一个学生都不选
    for (let i = 0; i < n - 1; i++) {
        if (nums[i] < i + 1 && i + 1 < nums[i + 1]) {
            ans++;
        }
    }
    return ans + 1; // +1 是因为可以都选
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

注：如果采用桶排序，可以做到 $\mathcal{O}(n)$ 的时间复杂度。
