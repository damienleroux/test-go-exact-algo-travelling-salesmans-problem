package computeTour

/*
 * An IntSets is a gathering of IntSet
 */

type IntSets struct {
	sets []IntSet
}

func (current *IntSets) Get() []IntSet {
	return current.sets
}

func (current *IntSets) GetSet(index int) IntSet {
	return current.sets[index]
}

func (current *IntSets) Len() int {
	return len(current.sets)
}

func (current *IntSets) AppendIntSetSlice(other []IntSet) {
	current.sets = append(current.sets, other...)
}

func (current *IntSets) AppendIntSets(other IntSets) {
	current.sets = append(current.sets, other.Get()...)
}
func (current *IntSets) AppendIntSet(set IntSet) {
	current.sets = append(current.sets, set)
}

func newIntSets() IntSets {
	current := IntSets{}
	return current
}

func NewIntSetsFromIntSetSlice(elements []IntSet) IntSets {
	current := IntSets{}
	current.AppendIntSetSlice(elements)
	return current
}

func NewIntSetsFromIntSets(other IntSets) IntSets {
	current := IntSets{}
	current.AppendIntSets(other)
	return current
}
