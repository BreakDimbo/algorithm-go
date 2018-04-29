package lancher

import (
	"github.com/breakD/algorithms/dstruct"
)

// SimpleFixedStack example
func LanchLL() {
	first := dstruct.NewNode()
	second := dstruct.NewNode()
	third := dstruct.NewNode()

	first.Item = "to"
	second.Item = "be"
	third.Item = "or"

	first.Next = second
	second.Next = third

	dstruct.Traverse(first)
}
