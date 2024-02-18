对于每个单元格，枚举八个方向，生成数字，统计其中质数个数。

最后返回出现次数最多的质数，如果有多个这样的质数，返回最大的那个。

```py [sol-Python3]
class Solution:
    def mostFrequentPrime(self, mat: List[List[int]]) -> int:
        m, n = len(mat), len(mat[0])
        cnt = Counter()
        for i, row in enumerate(mat):
            for j, v in enumerate(row):
                for dx, dy in (1, 0), (1, 1), (0, 1), (-1, 1), (-1, 0), (-1, -1), (0, -1), (1, -1):
                    x, y, val = i + dx, j + dy, v
                    while 0 <= x < m and 0 <= y < n:
                        val = val * 10 + mat[x][y]
                        # 如果 val 在 cnt 中，那么 val 一定是质数
                        if val in cnt or self.is_prime(val):
                            cnt[val] += 1
                        x += dx
                        y += dy

        ans, max_cnt = -1, 0
        for v, c in cnt.items():
            if c > max_cnt:
                ans, max_cnt = v, c
            elif c == max_cnt:
                ans = max(ans, v)
        return ans

    def is_prime(self, n: int) -> bool:
        return all(n % i for i in range(2, isqrt(n) + 1))
```

```java [sol-Java]
class Solution {
    private static final int[][] DIRS = {{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}};

    public int mostFrequentPrime(int[][] mat) {
        int m = mat.length;
        int n = mat[0].length;
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                for (int[] d : DIRS) {
                    int x = i + d[0];
                    int y = j + d[1];
                    int v = mat[i][j];
                    while (x >= 0 && x < m && y >= 0 && y < n) {
                        v = v * 10 + mat[x][y];
                        if (isPrime(v)) {
                            cnt.merge(v, 1, Integer::sum);
                        }
                        x += d[0];
                        y += d[1];
                    }
                }
            }
        }

        int ans = -1;
        int maxCnt = 0;
        for (Map.Entry<Integer, Integer> e : cnt.entrySet()) {
            int v = e.getKey();
            int c = e.getValue();
            if (c > maxCnt) {
                ans = v;
                maxCnt = c;
            } else if (c == maxCnt) {
                ans = Math.max(ans, v);
            }
        }
        return ans;
    }

    private boolean isPrime(int n) {
        for (int i = 2; i * i <= n; i++) {
            if (n % i == 0) {
                return false;
            }
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
    static constexpr int dirs[8][2] = {{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}};

    bool is_prime(int n) {
        for (int i = 2; i * i <= n; i++) {
            if (n % i == 0) {
                return false;
            }
        }
        return true;
    }

public:
    int mostFrequentPrime(vector<vector<int>> &mat) {
        int m = mat.size(), n = mat[0].size();
        unordered_map<int, int> cnt;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                for (auto &d: dirs) {
                    int x = i + d[0], y = j + d[1], v = mat[i][j];
                    while (x >= 0 && x < m && y >= 0 && y < n) {
                        v = v * 10 + mat[x][y];
                        // 如果 v 在 cnt 中，那么 v 一定是质数
                        if (cnt.contains(v) || is_prime(v)) {
                            cnt[v]++;
                        }
                        x += d[0];
                        y += d[1];
                    }
                }
            }
        }

        int ans = -1, max_cnt = 0;
        for (auto &[v, c]: cnt) {
            if (c > max_cnt) {
                ans = v;
                max_cnt = c;
            } else if (c == max_cnt) {
                ans = max(ans, v);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func mostFrequentPrime(mat [][]int) int {
	dirs := []struct{ x, y int }{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	m, n := len(mat), len(mat[0])
	cnt := map[int]int{}
	for i, row := range mat {
		for j, v := range row {
			for _, d := range dirs {
				x, y, v := i+d.x, j+d.y, v
				for 0 <= x && x < m && 0 <= y && y < n {
					v = v*10 + mat[x][y]
					// 如果 v 在 cnt 中，那么 v 一定是质数
					if cnt[v] > 0 || isPrime(v) {
						cnt[v]++
					}
					x += d.x
					y += d.y
				}
			}
		}
	}

	ans, maxCnt := -1, 0
	for v, c := range cnt {
		if c > maxCnt {
			ans, maxCnt = v, c
		} else if c == maxCnt {
			ans = max(ans, v)
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mnk\cdot 10^{k/2})$，其中 $m$ 和 $n$ 分别为 $\textit{mat}$ 的行数和列数，$k=\max(m,n)$。总共有 $\mathcal{O}(mnk)$ 个数，判断质数需要 $\mathcal{O}(10^{k/2})$ 的时间。
- 空间复杂度：$\mathcal{O}(mnk)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
