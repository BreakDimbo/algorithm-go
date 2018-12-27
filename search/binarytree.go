package search

import "fmt"

type Tree struct {
	root *Node
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func NewNode(value int) *Node {
	return &Node{data: value}
}

// 不考虑值重复的情况
func (t *Tree) Add(value int) *Node {
	if t.root == nil {
		t.root = NewNode(value)
		return t.root
	}
	p := t.root

	for p != nil {
		if p.data > value {
			if p.left == nil {
				p.left = NewNode(value)
				return p.left
			}
			p = p.left
		} else {
			if p.right == nil {
				p.right = NewNode(value)
				return p.right
			}
			p = p.right
		}
	}
	return nil
}

// 三种情况
// 无子节点
// 有一个子节点
// 有两个子节点
func (t *Tree) Del(value int) {
	p := t.root
	var pp *Node // 记录父节点

	// 定位待删除节点
	for p != nil && p.data != value {
		pp = p
		if value > p.data {
			p = p.right
		} else {
			p = p.left
		}
	}

	if p == nil {
		return
	}

	// 如果待删除节点有两个子节点，找到右子树的最小节点，与当前节点替换，删除右子树最小节点
	if p.left != nil && p.right != nil {
		pMin := p.right
		ppMin := p

		for pMin.left != nil {
			ppMin = pMin
			pMin = pMin.left
		}

		// 更新 data，准备删除 pMin
		p.data = pMin.data

		p = pMin
		pp = ppMin
	}

	// 删除的是叶子节点或者只有一个子节点
	var child *Node
	if p.left != nil {
		child = p.left
	} else if p.right != nil {
		child = p.right
	} else {
		child = nil
	}

	if pp == nil {
		// 如果删除的是根节点
		t.root = child
	} else if pp.left == p {
		pp.left = child
	} else {
		pp.right = child
	}
}

func (t *Tree) Find(value int) *Node {
	p := t.root

	for p != nil {
		if p.data == value {
			return p
		}

		if value > p.data {
			p = p.right
		} else {
			p = p.left
		}
	}
	return nil
}

// current node -> left node -> right node
func (t *Tree) PrintPreOrder() {
	fmt.Print("preOrder: ")
	t.printPreOrder(t.root)
	fmt.Println()
}

func (t *Tree) printPreOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%d|", node.data)
	t.printPreOrder(node.left)
	t.printPreOrder(node.right)

}

// left node -> current node  -> right node
func (t *Tree) PrintInOrder() {
	fmt.Print("inOrder: ")
	t.printInOrder(t.root)
	fmt.Println()
}

func (t *Tree) printInOrder(node *Node) {
	if node == nil {
		return
	}
	t.printInOrder(node.left)
	fmt.Printf("%d|", node.data)
	t.printInOrder(node.right)
}

// left node  -> right node -> current node
func (t *Tree) PrintPostOrder() {
	fmt.Print("postOrder: ")
	t.printPostOrder(t.root)
	fmt.Println()
}

func (t *Tree) printPostOrder(node *Node) {
	if node == nil {
		return
	}
	t.printPostOrder(node.left)
	t.printPostOrder(node.right)
	fmt.Printf("%d|", node.data)
}
