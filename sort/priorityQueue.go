package sort

type PQueue struct {
	pq         []Comparable
	elementNum int
}

func New(length int) *PQueue {
	return &PQueue{
		pq:         make([]Comparable, length+1),
		elementNum: 0,
	}
}

func (p *PQueue) IsEmpty() bool {
	return p.elementNum == 0
}

func (p *PQueue) Size() int {
	return p.elementNum
}

func (p *PQueue) Insert(e Comparable) {
	p.elementNum++
	p.pq[p.elementNum] = e
	p.swim(p.elementNum)
}

func (p *PQueue) DelMax() (max Comparable) {
	max = p.pq[1]
	p.exch(1, p.elementNum)
	p.elementNum--
	p.pq[p.elementNum+1] = nil
	p.sink(1)
	return
}

func (p *PQueue) less(i, j int) bool {
	return p.pq[i].CompareTo(p.pq[j]) < 0
}

func (p *PQueue) exch(i, j int) {
	t := p.pq[i]
	p.pq[i] = p.pq[j]
	p.pq[j] = t
}

func (p *PQueue) swim(k int) {
	for k > 1 && p.less(k/2, k) {
		p.exch(k, k/2)
		k = k / 2
	}
}

func (p *PQueue) sink(k int) {

	for 2*k <= p.elementNum {
		j := 2 * k
		if j < p.elementNum && p.less(j, j+1) {
			j++
		}

		if !p.less(k, j) {
			break
		}

		p.exch(k, j)
		k = j
	}
}
