package others

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	cache := Constructor(1)

	cache.Put(2, 1)
	cache.Get(2)
	cache.Put(3, 2)
	cache.Get(2)
	cache.Get(3)
}
