package lancher

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/breakD/algorithms/dstruct/stack"
)

// SimpleFixedStack example
func LanchSFStack() {
	sfstack := stack.NewSFStack(100)
	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadString('\n')

		cleanT := strings.TrimRight(text, "\n")
		if cleanT != "-" {
			sfstack.Push(cleanT)
		} else if !sfstack.IsEmpty() {
			fmt.Print(sfstack.Pop() + " ")
		}

		if err == io.EOF {
			break
		}
	}
	fmt.Printf("(%d left on stack)\n", sfstack.Size())
}
