package btree

// This code is based on the book by Kommadi - Learn data structures and algorithms
// with golang.

import (
	"fmt"
	"sync"
)

// TreeNode class
type TreeNode struct {
	key       int
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

// BinarySearchTree class
type BinarySearchTree struct {
	rootNode *TreeNode
	lock     sync.RWMutex
}

// InsertElement will insert a element to a BinarySearchTree class
func (tree *BinarySearchTree) InsertElement(key, value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	var treeNode *TreeNode
	treeNode = &TreeNode{key, value, nil, nil}

	if tree.rootNode == nil {
		tree.rootNode = treeNode
	} else {
		insertTreeNode(tree.rootNode, treeNode)
	}
}

// insertTreeNode inserts a node into the rootNode
func insertTreeNode(rootNode, newTreeNode *TreeNode) {
	// Insert low values to the left and high values
	// to the right.
	if newTreeNode.key < rootNode.key {
		if rootNode.leftNode == nil {
			rootNode.leftNode = newTreeNode
		} else {
			insertTreeNode(rootNode.leftNode, newTreeNode)
		}
	} else {
		if rootNode.rightNode == nil {
			rootNode.rightNode = newTreeNode
		} else {
			insertTreeNode(rootNode.rightNode, newTreeNode)
		}
	}

}

// InOrderTraverse visits all nodes in order.
func (tree *BinarySearchTree) InOrderTraverse(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	inOrderTraverseTree(tree.rootNode, function)
}

func inOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		inOrderTraverseTree(treeNode.leftNode, function)
		function(treeNode.value)
		inOrderTraverseTree(treeNode.rightNode, function)
	}
}

// PreOrderTraverse visits all tree nodes with preorder traversing.
func (tree *BinarySearchTree) PreOrderTraverse(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	preOrderTraverseTree(tree.rootNode, function)
}

func preOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		function(treeNode.value)
		preOrderTraverseTree(treeNode.leftNode, function)
		preOrderTraverseTree(treeNode.rightNode, function)
	}
}

// PostOrderTraverse visits all tree nodes with postorder traversing.
func (tree *BinarySearchTree) PostOrderTraverse(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	postOrderTraverseTree(tree.rootNode, function)
}

func postOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		postOrderTraverseTree(treeNode.leftNode, function)
		postOrderTraverseTree(treeNode.rightNode, function)
		function(treeNode.value)
	}
}

// MinNode finds the node with the maximum value on a Binary Search Tree.
func (tree *BinarySearchTree) MinNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	var treeNode *TreeNode
	treeNode = tree.rootNode
	if treeNode == nil {
		return nil
	}
	// Search for the end of all left nodes to find the minimum value
	for {
		if treeNode.leftNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.leftNode
	}
}

// MaxNode finds the node with the minimum value on a Binary Search Tree.
func (tree *BinarySearchTree) MaxNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	var treeNode *TreeNode
	treeNode = tree.rootNode
	if treeNode == nil {
		return nil
	}
	// Search for the end of all right nodes to find the minimum value
	for {
		if treeNode.rightNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.rightNode
	}
}

// SearchNode finds a specific node in the binary search tree
func (tree *BinarySearchTree) SearchNode(key int) bool {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	return searchNode(tree.rootNode, key)
}

func searchNode(treeNode *TreeNode, key int) bool {
	if treeNode == nil {
		return false
	}

	if key < treeNode.key {
		return searchNode(treeNode.leftNode, key)
	}

	if key > treeNode.key {
		return searchNode(treeNode.rightNode, key)
	}

	return true
}

// RemoveNode remove node for a given key
func (tree *BinarySearchTree) RemoveNode(key int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	removeNode(tree.rootNode, key)
}

func removeNode(treeNode *TreeNode, key int) *TreeNode {
	if treeNode == nil {
		return nil
	}

	if key < treeNode.key {
		treeNode.leftNode = removeNode(treeNode.leftNode, key)
		return treeNode
	}

	if key > treeNode.key {
		treeNode.rightNode = removeNode(treeNode.rightNode, key)
		return treeNode
	}

	// the key is matched
	if treeNode.leftNode == nil && treeNode.rightNode == nil {
		treeNode = nil
		return nil
	}

	if treeNode.leftNode == nil {
		treeNode = treeNode.rightNode
		return treeNode
	}

	if treeNode.rightNode == nil {
		treeNode = treeNode.leftNode
		return treeNode
	}

	// Search the smallest value on the right side
	var leftmostrightNode *TreeNode
	leftmostrightNode = treeNode.rightNode
	for {
		if leftmostrightNode != nil && leftmostrightNode.leftNode != nil {
			leftmostrightNode = leftmostrightNode.leftNode
		} else {
			break
		}
	}

	treeNode.key, treeNode.value = leftmostrightNode.key, leftmostrightNode.value
	treeNode.rightNode = removeNode(treeNode.rightNode, treeNode.key)
	return treeNode
}

// String shows a string representation of the Binary Tree
func (tree *BinarySearchTree) String() {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	fmt.Println("Binary Tree")
	fmt.Println("------------------------------------")
	stringify(tree.rootNode, 0)
	fmt.Println("------------------------------------")
}

func stringify(treeNode *TreeNode, level int) {
	if treeNode != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "	"
		}
		format += "---[ "
		level++
		stringify(treeNode.leftNode, level)
		fmt.Printf(format+"%v\n", treeNode.key)
		stringify(treeNode.rightNode, level)
	}
}

func main() {
	bt := BinarySearchTree{}

	norder := []int{8, 3, 1, 6, 10, 19}

	for _, val := range norder {
		bt.InsertElement(val, val)
	}

	bt.PostOrderTraverse(func(i int) {
		fmt.Println(i)
	})

	bt.String()

	// max := *bt.MaxNode()
	// min := *bt.MinNode()
	// fmt.Println("Max:", max, "Min:", min)

	// fmt.Println("Does the binary tree have Key=1 :", bt.SearchNode(1))
	// fmt.Println("Does the binary tree have Key=5 :", bt.SearchNode(5))
}
