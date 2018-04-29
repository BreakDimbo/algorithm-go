package dstruct

import (
	"fmt"
)

func Traverse(i *Node) {
	fmt.Println("Start traversing the linked list.")
	for x := i; x != nil; x = x.Next {
		fmt.Printf("the item is: %s\n", x.Item)
	}
}
