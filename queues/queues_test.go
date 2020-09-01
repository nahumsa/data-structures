package queues

import "testing"

func TestAdd(t *testing.T) {
	queue := Queue{}

	order1 := &Order{}
	order1.New(2, 20, "Computer", "Greg White")
	queue.Add(order1)

	order2 := &Order{}
	order2.New(5, 20, "Cellphone", "Greg Red")
	queue.Add(order2)

	order3 := &Order{}
	order3.New(1, 20, "Cellphone", "Greg Red")
	queue.Add(order3)

	if queue[0].Priority != 5 {
		t.Errorf("Expected was %v, we got %v", 5, queue[0].Priority)
	}
}

func TestDoneQueue(t *testing.T) {
	queue := Queue{}

	order1 := &Order{}
	order1.New(2, 20, "Computer", "Greg White")
	queue.Add(order1)

	order2 := &Order{}
	order2.New(5, 20, "Cellphone", "Greg Red")
	queue.Add(order2)

	order3 := &Order{}
	order3.New(1, 20, "Cellphone", "Greg Red")
	queue.Add(order3)

	queue.DoneOrder()

	if queue[0].Priority != 2 {
		t.Errorf("Expected was %v, we got %v", 2, queue[0].Priority)
	}
}
