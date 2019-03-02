package string

/*
	solution 1 traverse k times O(kn)
	solution 2 use hash O(n)
	solution 3 bitMap O(1)
*/

type BitMap struct {
	bits ['z' - 'A']uint8
}

func NewBitMap() *BitMap {
	return &BitMap{
		['z' - 'A']uint8{},
	}
}

func (b *BitMap) Add(n int) {
	word, bit := n/8, n%8
	b.bits[word] |= (1 << uint8(bit))
}

func (b *BitMap) Has(n int) bool {
	word, bit := n/8, n%8
	return b.bits[word]&(1<<uint8(bit)) != 0
}

func numJewelsInStones(J string, S string) int {
	bitM := NewBitMap()
	for _, r := range J {
		bitM.Add(int(r - 'A'))
	}
	count := 0
	for _, r := range S {
		if bitM.Has(int(r - 'A')) {
			count++
		}
	}
	return count
}
