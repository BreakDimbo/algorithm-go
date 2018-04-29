package lancher

import (
	"fmt"

	"github.com/breakD/algorithms/dstruct/queue"
)

// QueueLancher example
func LanchQueue() {
	queue := &queue.Queue{}
	count := 10
	for index := 0; index < count; index++ {
		queue.Enqueue(index)
		queue.Traverse()
	}

	for index := 0; index < count; index++ {
		queue.Dequeue()
		queue.Traverse()
		fmt.Printf("dequeue result: queue size: %d, isEmpty: %t\n", queue.Size(), queue.IsEmpty())
	}
}
