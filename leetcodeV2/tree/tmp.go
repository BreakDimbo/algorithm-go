package tree

/*
import "math"

var pre *TreeNode

func isValidBST(root *TreeNode) bool {
	pre = nil
	return helper(root)
}

func helper(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !helper(root.Left) {
		return false
	}
	if pre != nil && pre.Val >= root.Val {
		return false
	}
	pre = root
	return helper(root.Right)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return right
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	for root != nil {
		if root.Val > q.Val && root.Val > p.Val {
			root = root.Left
		} else if root.Val < q.Val && root.Val < p.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return root
}

func maxDepth(root *TreeNode) int {
	maxDepth := 0
	depthDfs(root, 0, &maxDepth)
	return maxDepth
}

func depthDfs(node *TreeNode, depth int, maxDepth *int) {
	if node == nil {
		if depth > *maxDepth {
			*maxDepth = depth
		}
		return
	}

	depthDfs(node.Left, depth+1, maxDepth)
	depthDfs(node.Right, depth+1, maxDepth)
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minDepth := math.MaxInt32
	minDepthDfs(root, 0, &minDepth)
	return minDepth
}

func minDepthDfs(node *TreeNode, depth int, minDepth *int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil && depth < *minDepth {
		*minDepth = depth
		return
	}
	minDepthDfs(node.Left, depth+1, minDepth)
	minDepthDfs(node.Right, depth+1, minDepth)
}
*/

/*
func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	stack := []*TreeNode{}
	pNode := root

	for pNode != nil || len(stack) != 0 {
		if pNode != nil {
			stack = append(stack, pNode)
			pNode = pNode.Left
		} else {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pre != nil && pre.Val >= v.Val {
				return false
			}
			pre = v
			pNode = v.Right
		}
	}
	return true
}
*/

/*
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}
	if l == nil {
		return r
	}
	return l
}
*/

/*
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if root.Val > q.Val && root.Val > p.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}
*/

/*
func preorderTraversal(root *TreeNode) []int {
	r := []int{}
	stack := []*TreeNode{}
	pNode := root

	for len(stack) != 0 || pNode != nil {
		if pNode != nil {
			r = append(r, pNode.Val)
			stack = append(stack, pNode)
			pNode = pNode.Left
		} else {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pNode = v.Right
		}
	}
	return r
}
*/

/*
func inorderTraversal(root *TreeNode) []int {
	r := []int{}
	stack := []*TreeNode{}

	for len(stack) != 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			r = append(r, v.Val)
			root = v.Right
		}
	}
	return r
}
*/
/*
func postorderTraversal(root *TreeNode) []int {
	r := []int{}
	stack := []*TreeNode{}
	var pre *TreeNode

	for len(stack) != 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			v := stack[len(stack)-1]
			if v.Right == nil || v.Right == pre {
				r = append(r, v.Val)
				stack = stack[:len(stack)-1]
				pre = v
			} else {
				root = v.Right
			}
		}
	}
	return r
}
*/

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	r := 1.0
	for n != 0 {
		if n&1 == 1 {
			r *= x
		}
		x *= x
		n >>= 1
	}
	return r
}

func majorityElement(nums []int) int {
	can := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			can = num
		}
		if can == num {
			count++
		} else {
			count--
		}
	}
	return can
}

func maxProfit(prices []int) int {

}
