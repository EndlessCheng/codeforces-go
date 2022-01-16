逆向思维 + 贪心

由于加倍次数有限，我们应该将其留在数字更大的时候加倍，这样可以节省更多的行动次数。

从 $v=\textit{target}$ 倒推思考：

- 若 $v$ 为奇数，则其只能由 $v-1$ 递增得到。
- 若 $v$ 为偶数且还有加倍次数，则其可以由 $\dfrac{v}{2}$ 加倍得到。由于我们是从大往小考虑，因此若能加倍则必加倍。若此时无加倍次数，则后面都是递增操作，此时可以直接算出答案并返回。

```go
func minMoves(v, maxDoubles int) (ans int) {
	for v > 1 {
		if maxDoubles == 0 {
			return ans + v - 1
		}
		if v%2 > 0 {
			v--
			ans++
		}
		maxDoubles--
		v /= 2
		ans++
	}
	return
}
```
