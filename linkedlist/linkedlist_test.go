package linkedlist

import (
	"testing"
)

func TestHeadNode(t *testing.T) {
	want := []int{10, 50, 100}
	for _, property := range want {
		linkedlist := LinkedList{}
		linkedlist.AddToHead(property)
		if linkedlist.HeadNode.Property != property {
			t.Errorf("Expected was %v, we got %v", property, linkedlist.HeadNode.Property)
		}
	}
}

func TestAddToEnd(t *testing.T) {
	firstProperty := 5
	linkedlist := LinkedList{}
	linkedlist.AddToHead(firstProperty)
	want := []int{10, 50, 100}
	for _, property := range want {
		linkedlist.AddToEnd(property)
		if linkedlist.LastNode().Property != property {
			t.Errorf("Expected was %v, we got %v", property, linkedlist.LastNode().Property)
		}
	}
}

func TestAddAfter(t *testing.T) {
	firstProperty := 5
	linkedlist := LinkedList{}
	linkedlist.AddToHead(firstProperty)

	want := []int{10, 50, 100}
	adds := []int{5, 15, 25, 35}
	for _, values := range adds {
		linkedlist.AddToEnd(values)
	}
	for i, property := range want {
		linkedlist.AddAfter(adds[i], property)
		gotNode := linkedlist.NodeWithProperty(property)
		if gotNode.Property != property {
			t.Errorf("Expected was %v, we got %v", property, gotNode.Property)
		}
	}
}
