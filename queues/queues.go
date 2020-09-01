package queues

// Queue is a slice of orders.
type Queue []Order

// Order is a struct with the priority quantity, product and costumer name.
type Order struct {
	Priority     int
	Quantity     int
	Product      string
	CustomerName string
}

// New initializes with order with priority, quantity,
// costumerName
func (order *Order) New(priority int, quantity int, product string, costumerName string) {
	order.Priority = priority
	order.Quantity = quantity
	order.Product = product
	order.CustomerName = costumerName
}

// Add adds the order to the queue
func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, *order)
	} else {
		appended := false
		for i, addedOrder := range *queue {
			if order.Priority > addedOrder.Priority {
				*queue = append((*queue)[:i], append(Queue{*order}, (*queue)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*queue = append(*queue, *order)
		}
	}
}

// DoneOrder delete the order with the most priority
func (queue *Queue) DoneOrder() {
	*queue = (*queue)[1:]
}
