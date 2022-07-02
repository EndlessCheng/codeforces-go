说一个简单的写法。

首先我们需要从上往下染色，因为如果先染下面的话，上面染色又会把下面的颜色给覆盖掉，这样又需要重新染色。

由于初始时，所有节点的颜色均为 $0$，而目标颜色 $c_i>0$，因此根节点肯定要染色，染色之后所有节点颜色均为 $c_1$。

然后顺着根节点往下走：

- 如果节点目标颜色和其父节点颜色相同，那么无需对该节点染色。
- 如果节点目标颜色和其父节点颜色不同，那么必须对该节点染色。

因此我们只需要判断有多少个节点和其父节点颜色不同就行了。

由于输入十分友好，告诉了每个节点的父节点，因此无需建图，直接判断即可。

```go
package main
import("bufio";."fmt";"os")

// https://space.bilibili.com/206214
func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	Fscan(in, &n)
	p := make([]int, n+1)
	for i := 2; i <= n; i++ { Fscan(in, &p[i]) }
	c := make([]int, n+1)
	for i := 1; i <= n; i++ { Fscan(in, &c[i]) }
	
	ans := 1
	for i := 2; i <= n; i++ {
		if c[i] != c[p[i]] { ans++ }
	}
	Print(ans)
}
```