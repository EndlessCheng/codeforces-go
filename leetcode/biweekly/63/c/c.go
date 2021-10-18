package main

/* BFS + 统计每个点变为空闲状态的最早秒数

题目可以抽象成一张图，我们可以用 BFS 求出 $0$ 号节点到其余节点的最短路长度，在 BFS 的同时求出每个节点变为空闲状态的最早秒数，其最大值就是整个计算机网络变为空闲状态的时间。

以其中某一节点 $v$ 为例，$v$ 变为空闲状态的最早秒数，就是其收到最后一条服务器给它的回复消息的时间的下一秒。记 $0$ 号节点到 $v$ 的最短路长度为 $d$，则从 $v$ 发出的第一条消息，到达 $0$ 后，回复消息又需要发给 $v$，这一共经过了 $2d$ 秒。此时我们需要知道 $v$ 发出的最后一条消息已经发出了多久（记作 $t$），这样就可以知道 $v$ 收到最后一条回复消息的时间。由于 $v$ 每 $p=\textit{patience}[v]$ 秒发出一条消息，我们按整除关系分类讨论：

- 若 $p$ 能整除 $2d$，则根据题意，此时最后一条消息恰好发出了 $p$ 秒；
- 若 $p$ 不能整除 $2d$，则最后一条消息发出了 $2d\bmod p$ 秒。例如 $2d=10,\ p=3$，则 $v$ 发出消息的时间为 $0,\ 3,\ 6,\ 9$，最后一条消息发出了 $10\bmod3=1$ 秒。

我们已经知道，从 $v$ 发出消息到收到回复要经过 $2d$ 秒，因此收到最后一条消息的时间为 $2d+(2d-t)=4d-t$ 秒，其下一秒为 $4d-t+1$ 秒，此时 $v$ 变为空闲状态。

所有节点变为空闲状态的时间的最大值即为整个计算机网络变为空闲状态的时间。

*/

// github.com/EndlessCheng/codeforces-go
func networkBecomesIdle(edges [][]int, patience []int) (ans int) {
	n := len(patience)
	g := make([][]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	vis := make([]bool, n)
	vis[0] = true
	q := []int{0}
	for d := 0; q != nil; d++ {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v > 0 {
				p := patience[v]
				t := p // 计算最后一条消息发出的秒数
				if 2*d%p > 0 {
					t = 2 * d % p
				}
				ans = max(ans, 4*d-t+1)
			}
			for _, w := range g[v] {
				if !vis[w] {
					vis[w] = true
					q = append(q, w)
				}
			}
		}
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
