## 视频讲解

请看 [视频讲解](https://www.bilibili.com/video/BV18u411Y7Gt/) 第四题。

## 思路

分类讨论：

1. 如果答案只有 $1$ 行，那么必须全为 $0$。反过来，如果存在全为 $0$ 的，返回这一行的下标作为答案。
2. 如果答案有 $2$ 行，那么不能有同一列均为 $1$。从二进制角度来理解，就是这两行 AND 的结果等于 $0$。反过来，如果两行 AND 的结果等于 $0$，就返回这两行的下标。
3. 如果答案有 $3$ 行，那么每一列的和均不超过 $1$，去掉任意一行后，每一列的和仍然均不超过 $1$。所以无需考虑这种情况。
4. **假定上面的情况都没有找到答案**。如果答案有 $4$ 行，那么考虑 $1$ 最少的那一行，其它行必须与这一行有交集（AND 不为 $0$）。继续分类讨论，当成一个**数独**游戏来玩（假定列数 $n=5$）：
   1. 如果这一行是 $10000$，那么其它行第一列必须是 $1$，但列的数字和就不符合要求了。
   2. 如果这一行是 $11000$，那么第二行可以是 $10\texttt{\_\_\_}$，第三行可以是 $01\texttt{\_\_\_}$，但第四行无论怎么填都会有一列的和超过 $2$，不符合要求。
   3. 对于 $1$ 更多的情况，由于每**列**至多 $2$ 个 $1$，总共至多 $2n=10$ 个 $1$，无法满足。例如这一行是 $11100$，由于我们考虑的是 $1$ 最少的行，其余行至少有 $3$ 个 $1$。那么第二行可以是 $10011$，第三行可以是 $01011$，第四行可以是 $00111$，前三列是符合要求的，但后面两列不符合要求。也可以这样理解，总的 $1$ 的个数至少是 $3\cdot 4=12$，它是大于 $10$ 的，不满足要求。
5. 如果答案超过 $4$ 行，类似上面的方法可以证明答案是不存在的。

因此，答案至多两行。

注：当 $n=6$ 时，有如下合法构造：

$$
111000\\
100110\\
010101\\
001011
$$

此时就要考虑 $4$ 行的情况了。（本题 $n$ 至多为 $5$）

```py [sol-Python3]
class Solution:
    def goodSubsetofBinaryMatrix(self, grid: List[List[int]]) -> List[int]:
        idx = {}
        for i, row in enumerate(grid):
            mask = 0
            for j, x in enumerate(row):
                mask |= x << j
            idx[mask] = i
        if 0 in idx:
            return [idx[0]]
        for x, i in idx.items():
            for y, j in idx.items():
                if (x & y) == 0:
                    return sorted((i, j))
        return []
```

```java [sol-Java]
class Solution {
    public List<Integer> goodSubsetofBinaryMatrix(int[][] grid) {
        var idx = new HashMap<Integer, Integer>();
        for (int i = 0; i < grid.length; i++) {
            int mask = 0;
            for (int j = 0; j < grid[i].length; j++)
                mask |= grid[i][j] << j;
            idx.put(mask, i);
        }
        if (idx.containsKey(0))
            return List.of(idx.get(0));
        for (var e1 : idx.entrySet())
            for (var e2 : idx.entrySet())
                if ((e1.getKey() & e2.getKey()) == 0) {
                    int i = e1.getValue(), j = e2.getValue();
                    return i < j ? List.of(i, j) : List.of(j, i);
                }
        return List.of();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> goodSubsetofBinaryMatrix(vector<vector<int>> &grid) {
        unordered_map<int, int> idx;
        for (int i = 0; i < grid.size(); i++) {
            int mask = 0;
            for (int j = 0; j < grid[i].size(); j++)
                mask |= grid[i][j] << j;
            idx[mask] = i;
        }
        if (idx.count(0)) return {idx[0]};
        for (auto [x, i]: idx)
            for (auto [y, j]: idx)
                if ((x & y) == 0)
                    return {min(i, j), max(i, j)};
        return {};
    }
};
```

```go [sol-Go]
func goodSubsetofBinaryMatrix(grid [][]int) []int {
	idx := map[int]int{}
	for i, row := range grid {
		mask := 0
		for j, x := range row {
			mask |= x << j
		}
		idx[mask] = i
	}
	if i, ok := idx[0]; ok {
		return []int{i}
	}
	for x, i := range idx {
		for y, j := range idx {
			if x&y == 0 {
				if i < j {
					return []int{i, j}
				}
				return []int{j, i}
			}
		}
	}
	return nil
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn+2^{2n})$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(2^n)$。
