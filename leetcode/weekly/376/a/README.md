用数组统计每个数的出现次数。

```py [sol-Python3]
class Solution:
    def findMissingAndRepeatedValues(self, grid: List[List[int]]) -> List[int]:
        n = len(grid)
        cnt = [0] * (n * n + 1)
        for row in grid:
            for x in row:
                cnt[x] += 1

        ans = [0, 0]
        for i in range(1, n * n + 1):
            if cnt[i] == 2:
                ans[0] = i
            elif cnt[i] == 0:
                ans[1] = i
        return ans
```

```java [sol-Java]
public class Solution {
    public int[] findMissingAndRepeatedValues(int[][] grid) {
        int n = grid.length;
        int[] cnt = new int[n * n + 1];
        for (int[] row : grid) {
            for (int x : row) {
                cnt[x]++;
            }
        }

        int[] ans = new int[2];
        for (int i = 1; i <= n * n; i++) {
            if (cnt[i] == 2) {
                ans[0] = i;
            } else if (cnt[i] == 0) {
                ans[1] = i;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> findMissingAndRepeatedValues(vector<vector<int>> &grid) {
        int n = grid.size();
        vector<int> cnt(n * n + 1);
        for (const vector<int> &row: grid) {
            for (int x: row) {
                cnt[x]++;
            }
        }

        vector<int> ans(2, 0);
        for (int i = 1; i <= n * n; i++) {
            if (cnt[i] == 2) {
                ans[0] = i;
            } else if (cnt[i] == 0) {
                ans[1] = i;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	cnt := make([]int, n*n+1)
	for _, row := range grid {
		for _, x := range row {
			cnt[x]++
		}
	}

	ans := [2]int{}
	for i := 1; i <= n*n; i++ {
		if cnt[i] == 2 {
			ans[0] = i
		} else if cnt[i] == 0 {
			ans[1] = i
		}
	}
	return ans[:]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(n^2)$。
