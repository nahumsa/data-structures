package linkedlist

// This code is based on the book by Kommadi - Learn data structures and algorithms
// with golang.// This code is based on the book by Kommadi - Learn data structures and algorithms
// with golang.

import "fmt"

// Node class that has a int property and a link to
// the next node
type Node struct {
	Property int
	NextNode *Node
}

// LinkedList class that points to a headNode
type LinkedList struct {
	HeadNode *Node
}

// AddToHead adds a node to the start of the linked
// list.
func (linkedList *LinkedList) AddToHead(property int) {
	node := Node{}
	node.Property = property
	if node.NextNode != nil {
		node.NextNode = linkedList.HeadNode
	}
	linkedList.HeadNode = &node
}

// IterateList iterates from the head node into the final
// node of the linked list.
func (linkedList *LinkedList) IterateList() {
	for node := linkedList.HeadNode; node != nil; node = node.NextNode {
		fmt.Println(node.Property)
	}
}

// LastNode returns the last node of the linked list.
func (linkedList *LinkedList) LastNode() *Node {
	var lastNode *Node
	for node := linkedList.HeadNode; node != nil; node = node.NextNode {
		if node.NextNode == nil {
			lastNode = node
		}
	}

	return lastNode
}

// AddToEnd adds a node to the end of the linked list.
func (linkedList *LinkedList) AddToEnd(property int) {
	node := Node{
		Property: property,
		NextNode: nil,
	}
	lastNode := linkedList.LastNode()
	if lastNode != nil {
		lastNode.NextNode = &node
	}
}

// NodeWithProperty returns the node with property value.
func (linkedList *LinkedList) NodeWithProperty(property int) *Node {
	var node *Node
	var nodeWith *Node

	for node = linkedList.HeadNode; node != nil; node = node.NextNode {
		if node.Property == property {
			nodeWith = node
			break
		}
	}
	if nodeWith == nil {
		return nil
	}

	return nodeWith
}

// AddAfter adds a node after the a node with a given property.
func (linkedList *LinkedList) AddAfter(nodeProperty, property int) {
	node := &Node{
		Property: property,
		NextNode: nil,
	}
	nodeWith := linkedList.NodeWithProperty(nodeProperty)
	if nodeWith != nil {
		node.NextNode = nodeWith.NextNode
		nodeWith.NextNode = node
	}
}

// DeleteProperty deletes an element with a given property.
func (linkedList *LinkedList) DeleteProperty(property int) {
	var node *Node

	for node = linkedList.HeadNode; node != nil; node = node.NextNode {
		if node.NextNode.Property == property {
			node.NextNode = node.NextNode.NextNode
			break
		}
	}

	if node == nil {
		node = &Node{}
	}
}
