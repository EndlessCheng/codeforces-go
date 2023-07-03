既然每次选最大的数，那么干脆对每行排序，这样每次就选的是一列的最大值。

累加这些最大值，即为答案。

视频讲解见[【双周赛 104】](https://www.bilibili.com/video/BV1fV4y1r7e6/)第二题。

```py [sol1-Python3]
class Solution:
    def matrixSum(self, nums: List[List[int]]) -> int:
        for row in nums: row.sort()
        return sum(map(max, zip(*nums)))  # zip(*nums) 枚举每一列
```

```java [sol1-Java]
class Solution {
    public int matrixSum(int[][] nums) {
        for (var row : nums)
            Arrays.sort(row);
        int ans = 0, n = nums[0].length;
        for (int j = 0; j < n; j++) {
            int mx = 0;
            for (var row : nums)
                mx = Math.max(mx, row[j]);
            ans += mx;
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int matrixSum(vector<vector<int>> &nums) {
        for (auto &row: nums)
            sort(row.begin(), row.end());
        int ans = 0, n = nums[0].size();
        for (int j = 0; j < n; j++) {
            int mx = 0;
            for (auto &row: nums)
                mx = max(mx, row[j]);
            ans += mx;
        }
        return ans;
    }
};
```

```go [sol1-Go]
func matrixSum(nums [][]int) (ans int) {
	for _, row := range nums {
		sort.Ints(row)
	}
	for j := range nums[0] {
		mx := 0
		for _, row := range nums {
			mx = max(mx, row[j])
		}
		ans += mx
	}
	return
}

func max(a, b int) int { if a < b { return b }; return a }
```

```js [sol-JavaScript]
var matrixSum = function (nums) {
    for (let row of nums)
        row.sort((a, b) => a - b);
    let ans = 0, n = nums[0].length;
    for (let j = 0; j < n; j++) {
        let mx = 0;
        for (var row of nums)
            mx = Math.max(mx, row[j]);
        ans += mx;
    }
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn\log n)$，其中 $m$ 和 $n$ 分别为 $\textit{nums}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序时的栈开销，仅用到若干额外变量。
