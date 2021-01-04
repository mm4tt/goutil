package ds

import "fmt"

type BitSet uint64

// CreateBitSet creates a BitSet containg the provided values.
func CreateBitSet(i ...int) BitSet {
	var b BitSet
	for _, i := range i {
		b.Insert(i)
	}
	return b
}

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
	b.checkInput(i)
	return (*b & (1 << i)) > 0
}

func (b *BitSet) Insert(i int) {
	b.checkInput(i)
	*b |= 1 << i
}

func (b *BitSet) Delete(i int) {
	b.checkInput(i)
	*b &= ^(1 << i)
}

func (b *BitSet) Intersect(set BitSet) {
	*b &= set
}

func (b BitSet) String() string {
	return fmt.Sprintf("%v", b.Elems())
}

func (b BitSet) checkInput(i int) {
	if !(0 <= i && i < 64) {
		panic(fmt.Sprintf("Invalid BitSet item: %d not in [0, 63]", i))
	}
}
