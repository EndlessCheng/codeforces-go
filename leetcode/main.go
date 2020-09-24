package main

import (
    . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
    "math/bits"
    "sort"
)

func main() {
    toBytes := func(g [][]string) [][]byte {
        n, m := len(g), len(g[0])
        bytes := make([][]byte, n)
        for i := range bytes {
            bytes[i] = make([]byte, m)
            for j := range bytes[i] {
                bytes[i][j] = g[i][j][0]
            }
        }
        return bytes
    }

    _ = MustBuildTreeNode

    _ = []interface{}{toBytes, ListNode{}, TreeNode{}}
}

// LC 37
func solveSudoku(board [][]byte) {
    var line, column [9]int
    var block [3][3]int
    var spaces [][2]int

    flip := func(i, j int, digit byte) {
        line[i] ^= 1 << digit
        column[j] ^= 1 << digit
        block[i/3][j/3] ^= 1 << digit
    }

    for i, row := range board {
        for j, b := range row {
            if b != '.' {
                digit := b - '1'
                flip(i, j, digit)
            }
        }
    }

    for {
        modified := false
        for i, row := range board {
            for j, b := range row {
                if b != '.' {
                    continue
                }
                mask := 0x1ff &^ uint(line[i]|column[j]|block[i/3][j/3])
                if mask&(mask-1) == 0 {
                    digit := byte(bits.TrailingZeros(mask))
                    flip(i, j, digit)
                    board[i][j] = digit + '1'
                    modified = true
                }
            }
        }
        if !modified {
            break
        }
    }

    for i, row := range board {
        for j, b := range row {
            if b == '.' {
                spaces = append(spaces, [2]int{i, j})
            }
        }
    }

    var dfs func(int) bool
    dfs = func(pos int) bool {
        if pos == len(spaces) {
            return true
        }
        i, j := spaces[pos][0], spaces[pos][1]
        mask := 0x1ff &^ uint(line[i]|column[j]|block[i/3][j/3])
        for ; mask > 0; mask &= mask - 1 {
            digit := byte(bits.TrailingZeros(mask))
            flip(i, j, digit)
            board[i][j] = digit + '1'
            if dfs(pos + 1) {
                return true
            }
            flip(i, j, digit)
        }
        return false
    }
    dfs(0)
}

// LC 39
func combinationSum(a []int, target int) (ans [][]int) {
    b := []int{}
    var f func(p, rest int)
    f = func(p, rest int) {
        if p == len(a) {
            return
        }
        if rest == 0 {
            ans = append(ans, append([]int(nil), b...))
            return
        }
        f(p+1, rest)
        if rest-a[p] >= 0 {
            b = append(b, a[p])
            f(p, rest-a[p])
            b = b[:len(b)-1]
        }
    }
    f(0, target)
    return
}

// LC 40
func combinationSum2(a []int, target int) (ans [][]int) {
    min := func(a, b int) int {
        if a < b {
            return a
        }
        return b
    }

    sort.Ints(a)
    var freq [][2]int
    for _, v := range a {
        if freq == nil || v != freq[len(freq)-1][0] {
            freq = append(freq, [2]int{v, 1})
        } else {
            freq[len(freq)-1][1]++
        }
    }

    var b []int
    var f func(p, rest int)
    f = func(p, rest int) {
        if rest == 0 {
            ans = append(ans, append([]int(nil), b...))
            return
        }
        if p == len(freq) || rest < freq[p][0] {
            return
        }
        f(p+1, rest)
        most := min(rest/freq[p][0], freq[p][1])
        for i := 1; i <= most; i++ {
            b = append(b, freq[p][0])
            f(p+1, rest-i*freq[p][0])
        }
        b = b[:len(b)-most]
    }
    f(0, target)
    return
}

// LC 41
func firstMissingPositive(a []int) int {
    n := len(a)
    for i, v := range a {
        for 0 < v && v <= n && v != a[v-1] {
            a[i], a[v-1] = a[v-1], a[i]
            v = a[i]
        }
    }
    for i, v := range a {
        if i+1 != v {
            return i + 1
        }
    }
    return n + 1
}

// LC 47
// 给定一个可包含重复数字的序列，返回所有不重复的全排列
func permuteUnique(nums []int) (ans [][]int) {
    n := len(nums)
    sort.Ints(nums)
    perm := []int{}
    vis := make([]bool, n)
    var f func(int)
    f = func(p int) {
        if p == n {
            ans = append(ans, append([]int(nil), perm...))
            return
        }
        for i, v := range nums {
            if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
                continue
            }
            perm = append(perm, v)
            vis[i] = true
            f(p + 1)
            vis[i] = false
            perm = perm[:len(perm)-1]
        }
    }
    f(0)
    return
}

// LC 79
func exist(board [][]byte, word string) bool {
    type pair struct{ x, y int }
    var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    h, w := len(board), len(board[0])
    vis := make([][]bool, h)
    for i := range vis {
        vis[i] = make([]bool, w)
    }
    var f func(i, j, k int) bool
    f = func(i, j, k int) bool {
        if board[i][j] != word[k] {
            return false
        }
        if k == len(word)-1 {
            return true
        }
        vis[i][j] = true
        defer func() { vis[i][j] = false }() // 回溯时还原已访问的单元格
        for _, d := range dir4 {
            if x, y := i+d.x, j+d.y; 0 <= x && x < h && 0 <= y && y < w && !vis[x][y] {
                if f(x, y, k+1) {
                    return true
                }
            }
        }
        return false
    }
    for i, r := range board {
        for j := range r {
            if f(i, j, 0) {
                return true
            }
        }
    }
    return false
}

// LC 94
// Morris 中序遍历
func inorderTraversal(root *TreeNode) (res []int) {
    for root != nil {
        if root.Left != nil {
            // predecessor 节点表示当前 root 节点向左走一步，然后一直向右走至无法走为止的节点
            predecessor := root.Left
            for predecessor.Right != nil && predecessor.Right != root {
                // 有右子树且没有设置过指向 root，则继续向右走
                predecessor = predecessor.Right
            }
            if predecessor.Right == nil {
                // 将 predecessor 的右指针指向 root，这样后面遍历完左子树 root.Left 后，就能通过这个指向回到 root
                predecessor.Right = root
                // 遍历左子树
                root = root.Left
            } else { // predecessor 的右指针已经指向了 root，则表示左子树 root.Left 已经访问完了
                res = append(res, root.Val)
                // 恢复原样
                predecessor.Right = nil
                // 遍历右子树
                root = root.Right
            }
        } else { // 没有左子树
            res = append(res, root.Val)
            // 若有右子树，则遍历右子树
            // 若没有右子树，则整颗左子树已遍历完，root 会通过之前设置的指向回到这颗子树的父节点
            root = root.Right
        }
    }
    return
}

// LC 99
func recoverTree(root *TreeNode) {
    nodes := []*TreeNode{}
    var f func(o *TreeNode)
    f = func(o *TreeNode) {
        if o == nil {
            return
        }
        f(o.Left)
        nodes = append(nodes, o)
        f(o.Right)
    }
    f(root)
    so := make([]*TreeNode, len(nodes))
    copy(so, nodes)
    sort.Slice(so, func(i, j int) bool { return so[i].Val < so[j].Val })
    do := []*TreeNode{}
    for i, o := range nodes {
        if o.Val != so[i].Val {
            do = append(do, o)
        }
    }
    do[0].Val, do[1].Val = do[1].Val, do[0].Val
}

// LC 106
// 从中序与后序遍历序列构造二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 {
        return nil
    }
    rootVal := postorder[len(postorder)-1]
    for i, v := range inorder {
        if v == rootVal {
            return &TreeNode{
                rootVal,
                buildTree(inorder[:i], postorder[:i]),
                buildTree(inorder[i+1:], postorder[i:len(postorder)-1]),
            }
        }
    }
    panic(1)
}

// LC 124
func maxPathSum(root *TreeNode) int {
    max := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    ans := int(-1e18)
    var f func(*TreeNode) int
    f = func(o *TreeNode) int {
        if o == nil {
            return -1e18
        }
        l := max(f(o.Left), 0)
        r := max(f(o.Right), 0)
        ans = max(ans, o.Val+l+r)
        return o.Val + max(l, r)
    }
    f(root)
    return ans
}

// LC 152
func maxProduct(a []int) int {
    min := func(a, b int) int {
        if a < b {
            return a
        }
        return b
    }
    max := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    mi, mx, ans := a[0], a[0], a[0]
    for _, v := range a[1:] {
        mi, mx = min(v, min(v*mi, v*mx)), max(v, max(v*mi, v*mx))
        ans = max(ans, mx)
    }
    return ans
}

// LC 216
func combinationSum3(k int, n int) (ans [][]int) {
    var temp []int
    var dfs func(cur, rest int)
    dfs = func(cur, rest int) {
        // 找到一个答案
        if len(temp) == k && rest == 0 {
            ans = append(ans, append([]int(nil), temp...))
            return
        }
        // 剪枝：跳过的数字过多，后面已经无法选到 k 个数字
        if len(temp)+10-cur < k || rest < 0 {
            return
        }
        // 跳过当前数字
        dfs(cur+1, rest)
        // 选当前数字
        temp = append(temp, cur)
        dfs(cur+1, rest-cur)
        temp = temp[:len(temp)-1]
    }
    dfs(1, n)
    return
}

// LC 226
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    left := invertTree(root.Left)
    right := invertTree(root.Right)
    root.Left = right
    root.Right = left
    return root
}

// LC 332
func findItinerary(tickets [][]string) []string {
    g := map[string][]string{}
    for _, p := range tickets {
        g[p[0]] = append(g[p[0]], p[1])
    }
    for _, vs := range g {
        sort.Strings(vs)
    }

    path := make([]string, 0, len(tickets)+1)
    var f func(string)
    f = func(v string) {
        for len(g[v]) > 0 {
            w := g[v][0]
            g[v] = g[v][1:]
            f(w)
        }
        path = append(path, v)
    }
    f("JFK")

    for i, j := 0, len(path)-1; i < j; i++ {
        path[i], path[j] = path[j], path[i]
        j--
    }
    return path
}

// LC 538 1038
// 反序中序遍历
func convertBST(root *TreeNode) *TreeNode {
    sum := 0
    var f func(*TreeNode)
    f = func(o *TreeNode) {
        if o != nil {
            f(o.Right)
            sum += o.Val
            o.Val = sum
            f(o.Left)
        }
    }
    f(root)
    return root
}

// LC 968
func minCameraCover(root *TreeNode) int {
    min := func(a, b int) int {
        if a < b {
            return a
        }
        return b
    }
    var f func(*TreeNode) (a, b, c int)
    f = func(o *TreeNode) (a, b, c int) {
        if o == nil {
            return 1e9, 0, 0
        }
        la, lb, lc := f(o.Left)
        ra, rb, rc := f(o.Right)
        a = lc + rc + 1
        b = min(a, min(la+rb, ra+lb))
        c = min(a, lb+rb)
        return
    }
    _, ans, _ := f(root)
    return ans
}
