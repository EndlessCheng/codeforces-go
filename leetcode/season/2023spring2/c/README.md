晚上 8:30[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

先简单记录一下思路，直播结束后继续更新题解和其它语言。

1. BFS。
2. 状态为 $(i,j,k)$，表示当前位置在 $(i,j)$，**要去**提取 $\textit{mantra}[k]$。
3. 如果 $\textit{matrix}[i][j]=\textit{mantra}[k]$，则移动到状态 $(i,j,k+1)$。
4. 枚举周围四个格子，移动到状态 $(i',j',k)$。
3. 初始状态：$(0,0,0)$。
4. 终点为 $k=l$，这里 $l$ 为 $\textit{mantra}$ 的长度。

```py [sol1-Python3]
class Solution:
    def extractMantra(self, matrix: List[str], mantra: str) -> int:
        m, n = len(matrix), len(matrix[0])
        q = [(0, 0, 0)]  # 起点
        vis = {q[0]}
        step = 1
        while q:
            tmp = q
            q = []
            for i, j, k in tmp:
                if matrix[i][j] == mantra[k]:  # 可以提取
                    if k == len(mantra) - 1:  # 下一步就是终点，直接返回
                        return step
                    p = (i, j, k + 1)
                    if p not in vis:
                        vis.add(p)
                        q.append(p)
                # 枚举周围四个格子
                for x, y in (i + 1, j), (i - 1, j), (i, j + 1), (i, j - 1):
                    if 0 <= x < m and 0 <= y < n:
                        p = (x, y, k)
                        if p not in vis:
                            vis.add(p)
                            q.append(p)
            step += 1
        return -1  # 无法到达终点
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(mnl)$，其中 $m$ 和 $n$ 分别为 $\textit{matrix}$ 的行数和列数，$l$ 为 $\textit{mantra}$ 的长度。
- 空间复杂度：$\mathcal{O}(mnl)$。
