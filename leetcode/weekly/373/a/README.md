[本题视频讲解](https://www.bilibili.com/video/BV19N411j7Dj/)

由于循环左移 $n$ 次等价于循环左移 $0$ 次，左移 $n+1$ 次等价于循环左移 $1$ 次，……，所以可以先把 $k$ 模上 $m$。

如果此时 $k=0$，那么操作不会影响矩阵，直接返回 `true`。

否则模拟就行。

如果左移 $k$ 次可以让数组相等，那么右移 $k$ 次也可以让数组相等（证明见视频），所以怎么写都是对的，可以全部改成左移 $k$ 次或者右移 $k$ 次。

```py [sol-Python3]
class Solution:
    def areSimilar(self, mat: List[List[int]], k: int) -> bool:
        k %= len(mat[0])
        return k == 0 or all(r == r[k:] + r[:k] for r in mat)
```

```java [sol-Java]
class Solution {
    public boolean areSimilar(int[][] mat, int k) {
        int n = mat[0].length;
        k %= n;
        if (k == 0) {
            return true;
        }
        for (int[] r : mat) {
            for (int j = 0; j < n; j++) {
                if (r[j] != r[(j + k) % n]) {
                    return false;
                }
            }
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool areSimilar(vector<vector<int>> &mat, int k) {
        int n = mat[0].size();
        k %= n;
        if (k == 0) {
            return true;
        }
        for (auto &r: mat) {
            for (int j = 0; j < n; j++) {
                if (r[j] != r[(j + k) % n]) {
                    return false;
                }
            }
        }
        return true;
    }
};
```

```go [sol-Go]
func areSimilar(mat [][]int, k int) bool {
	n := len(mat[0])
	k %= n
	if k == 0 {
		return true
	}
	for _, r := range mat {
		if !slices.Equal(r, append(r[k:], r[:k]...)) {
			return false
		}
	}
	return true
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{mat}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。
