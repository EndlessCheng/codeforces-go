分层 BFS 的同时排序

不了解网格图 BFS 的同学可以先看看 1765 题。

对于此题，我们可以直接 BFS 然后按照题目要求排序。但由于排序时距离为第一关键字，因此我们可以采用分层 BFS 的做法，即每次 BFS 时向外扩展一层，然后对扩展出的这些位置排序，取将价格在 $[\textit{low},\textit{high}]$ 中的物品位置加入答案。

```go [sol1-Go]
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func highestRankedKItems(grid [][]int, pricing, start []int, k int) (ans [][]int) {
	m, n := len(grid), len(grid[0])
	low, high := pricing[0], pricing[1]
	sx, sy := start[0], start[1]
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	vis[sx][sy] = true
	q := [][]int{{sx, sy}}
	for len(q) > 0 { // 分层 BFS
		// 此时 q 内所有元素到起点的距离均相同，因此按照题目中的第 2~4 关键字排序后，就可以将价格在 [low,high] 内的位置加入答案
		sort.Slice(q, func(i, j int) bool {
			ax, ay, bx, by := q[i][0], q[i][1], q[j][0], q[j][1]
			pa, pb := grid[ax][ay], grid[bx][by]
			return pa < pb || pa == pb && (ax < bx || ax == bx && ay < by)
		})
		l := sort.Search(len(q), func(i int) bool { return grid[q[i][0]][q[i][1]] >= low })
		r := sort.Search(len(q), func(i int) bool { return grid[q[i][0]][q[i][1]] > high })
		ans = append(ans, q[l:r]...)
		if len(ans) >= k {
			return ans[:k]
		}
		tmp := q
		q = nil
		for _, p := range tmp {
			for _, d := range dirs {
				if x, y := p[0]+d.x, p[1]+d.y; 0 <= x && x < m && 0 <= y && y < n && !vis[x][y] && grid[x][y] != 0 {
					vis[x][y] = true
					q = append(q, []int{x, y})
				}
			}
		}
	}
	return
}
```

```C++ [sol1-C++]
int dirs[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

class Solution {
public:
    vector<vector<int>> highestRankedKItems(vector<vector<int>> &grid, vector<int> &pricing, vector<int> &start, int k) {
        vector<vector<int>> ans;
        int m = grid.size(), n = grid[0].size();
        int low = pricing[0], high = pricing[1];
        int sx = start[0], sy = start[1];
        vector<vector<int>> vis(m, vector<int>(n));
        vis[sx][sy] = 1;
        vector<vector<int>> q = {{sx, sy}};
        while (!q.empty()) { // 分层 BFS
            // 此时 q 内所有元素到起点的距离均相同，因此按照题目中的第 2~4 关键字排序后，就可以将价格在 [low,high] 内的位置加入答案
            sort(q.begin(), q.end(), [&](auto &a, auto &b) {
                int pa = grid[a[0]][a[1]], pb = grid[b[0]][b[1]];
                return pa < pb || pa == pb && a < b;
            });
            for (auto &p : q) {
                if (low <= grid[p[0]][p[1]] && grid[p[0]][p[1]] <= high) {
                    ans.emplace_back(p);
                    if (ans.size() == k) {
                        return ans;
                    }
                }
            }
            vector<vector<int>> qq;
            for (auto &p : q) {
                for (auto &d : dirs) {
                    int x = p[0] + d[0], y = p[1] + d[1];
                    if (0 <= x && x < m && 0 <= y && y < n && !vis[x][y] && grid[x][y]) {
                        vis[x][y] = true;
                        vector<int> p = {x, y};
                        qq.emplace_back(p);
                    }
                }
            }
            q = move(qq);
        }
        return ans;
    }
};
```

```Python [sol1-Python3]
dirs = ((-1, 0), (1, 0), (0, -1), (0, 1))

class Solution:
    def highestRankedKItems(self, grid: List[List[int]], pricing: List[int], start: List[int], k: int) -> List[List[int]]:
        ans = []
        m, n = len(grid), len(grid[0])
        low, high = pricing
        sx, sy = start
        vis = {(sx, sy)}
        q = [(sx, sy)]
        while q:  # 分层 BFS
            # 此时 q 内所有元素到起点的距离均相同，因此按照题目中的第 2~4 关键字排序后，就可以将价格在 [low,high] 内的位置加入答案
            q.sort(key=lambda p: (grid[p[0]][p[1]], p))
            ans.extend(p for p in q if low <= grid[p[0]][p[1]] <= high)
            if len(ans) >= k:
                return ans[:k]
            tmp = q
            q = []
            for p in tmp:
                for d in dirs:
                    x, y = p[0] + d[0], p[1] + d[1]
                    if 0 <= x < m and 0 <= y < n and grid[x][y] and (x, y) not in vis:
                        vis.add((x, y))
                        q.append((x, y))
        return ans
```
