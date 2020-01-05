package computeTour

/*
 * An intSet is a set of integer aka []int
 * This file aims to gather the functions useful to manipulate a set of integers
 */

type IntSet struct {
	elements []int
}

func (current *IntSet) Get() []int {
	return current.elements
}
func (current *IntSet) Len() int {
	return len(current.elements)
}

func (current *IntSet) AppendInts(elements []int) {
	current.elements = append(current.elements, elements...)
}

func (current *IntSet) AppendIntSet(other IntSet) {
	current.elements = append(current.elements, other.Get()...)
}

func (current *IntSet) AppendInt(element int) {
	current.elements = append(current.elements, element)
}

// Add every integer between beg and end, including beg and end
func (current *IntSet) AppendRangeInt(beg int, end int) {
	for i := beg; i <= end; i++ {
		current.AppendInt(i)
	}
}

func newIntSet() IntSet {
	current := IntSet{}
	return current
}

func NewIntSetFromInts(elements []int) IntSet {
	current := IntSet{}
	current.AppendInts(elements)
	return current
}

func NewIntSetFromIntSet(other IntSet) IntSet {
	current := IntSet{}
	current.AppendIntSet(other)
	return current
}
