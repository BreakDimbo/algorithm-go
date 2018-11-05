package item

type Item interface {
	Equal(Item) bool
}
