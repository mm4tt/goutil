package ds

import "fmt"

type BitSet uint64

const _1 uint64 = 1

func (b *BitSet) IsEmpty() bool {
	return *b == 0
}

func (b *BitSet) Size() int {
	return len(b.Elems())
}

func (b *BitSet) Elems() []int {
	b1 := *b
	var ret []int
	for i := 0; !b1.IsEmpty(); i++ {
		if b1.Contains(0) {
			ret = append(ret, i)
		}
		b1 >>= 1
	}
	return ret
}

func (b *BitSet) Contains(i int) bool {
	return (*b & (_1 << i)) > 0
}

func (b *BitSet) Insert(i int) {
	*b |= _1 << i
}

func (b *BitSet) Delete(i int) {
	*b &= ^(_1 << i)
}

func (b *BitSet) Intersect(set BitSet) {
	*b &= set
}

func (b BitSet) String() string {
	return fmt.Sprintf("%v", b.Elems())
}
