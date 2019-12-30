package computeTour

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isSetInSets(setToFind IntSet, sets []IntSet) bool {
	for _, set := range sets {
		//only set with same number of element are concerned
		if set.Len() == setToFind.Len() {
			for i, setIndex := range set.Get() {
				indexFound := false
				for _, setToFindIndex := range setToFind.Get() {
					if setToFindIndex == setIndex {
						indexFound = true
					}
				}
				if !indexFound {
					break
				}
				//if every indexes in set are verified without breaking
				if i == set.Len()-1 {
					//the indexes set has been found
					return true
				}
			}
		}
	}
	return false
}

func areSaleSetsInSets(SalesSetToFind IntSets, SalesSets []IntSets) bool {
	for _, SalesSet := range SalesSets {
		//only set with same number of element are concerned
		if SalesSet.Len() == SalesSetToFind.Len() {
			for i, SaleSet := range SalesSet.Get() {
				setFound := isSetInSets(SaleSet, SalesSet.Get())
				if !setFound {
					break
				}
				//if every indexes in set are verified without breaking
				if i == SalesSet.Len()-1 {
					//the indexes set has been found
					return true
				}
			}
		}
	}
	return false
}

func testGetPossibleSaleIndexesSets(t *testing.T, salesNumber int, expectedSets []IntSets) {
	resultSaleIndexesSets := GetPossibleSaleIndexesSets(salesNumber, 100)
	//t.Log(resultSaleIndexesSets)
	//test number of solutions found
	expectedNumberOfSet := len(expectedSets)
	assert.Equal(t, len(resultSaleIndexesSets), expectedNumberOfSet, fmt.Sprintf("%d should have been found for %d sales", expectedNumberOfSet, salesNumber))

	//test solutions
	for _, expectedSet := range expectedSets {
		assert.Equal(t, areSaleSetsInSets(expectedSet, resultSaleIndexesSets), true, fmt.Sprintf("%v is expected to be found within results set ", expectedSet))
	}

}

func TestGetPossibleSaleIndexesSets3(t *testing.T) {

	// test how to share 3 sales among 2 salers
	testGetPossibleSaleIndexesSets(
		t,
		//number of sales
		3,
		//expected set
		[]IntSets{
			NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{0}),    //sale 1 one can handle first command
				NewIntSetFromInts([]int{1, 2}), //while the other one can handle the toher ones
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{0, 1}),
				NewIntSetFromInts([]int{2}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{0, 1, 2}),
				NewIntSetFromInts([]int{}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{0, 2}),
				NewIntSetFromInts([]int{1}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{1}),
				NewIntSetFromInts([]int{0, 2}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{1, 2}),
				NewIntSetFromInts([]int{0}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{2}),
				NewIntSetFromInts([]int{0, 1}),
			}),
		})
}

func TestGetPossibleSaleIndexesSets5(t *testing.T) {
	testGetPossibleSaleIndexesSets(
		t,
		//number of sales
		5,
		//expected set
		[]IntSets{
			NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{0}),
				NewIntSetFromInts([]int{1, 2, 3, 4}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{0, 1}),
				NewIntSetFromInts([]int{2, 3, 4}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{0, 1, 2}),
				NewIntSetFromInts([]int{3, 4}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{0, 1, 2, 3}),
				NewIntSetFromInts([]int{4}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{0, 1, 2, 3, 4}),
				NewIntSetFromInts([]int{}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{0, 1, 2, 4}),
				NewIntSetFromInts([]int{3}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{0, 1, 3}),
				NewIntSetFromInts([]int{2, 4}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{0, 1, 3, 4}),
				NewIntSetFromInts([]int{2}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{0, 1, 4}),
				NewIntSetFromInts([]int{2, 3}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{0, 2}),
				NewIntSetFromInts([]int{1, 3, 4}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{0, 2, 3}),
				NewIntSetFromInts([]int{1, 4}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{0, 2, 3, 4}),
				NewIntSetFromInts([]int{1}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{0, 2, 4}),
				NewIntSetFromInts([]int{1, 3}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{0, 3}),
				NewIntSetFromInts([]int{1, 2, 4}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{0, 3, 4}),
				NewIntSetFromInts([]int{1, 2}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{0, 4}),
				NewIntSetFromInts([]int{1, 2, 3}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{1}),
				NewIntSetFromInts([]int{0, 2, 3, 4}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{1, 2}),
				NewIntSetFromInts([]int{0, 3, 4}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{1, 2, 3}),
				NewIntSetFromInts([]int{0, 4}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{1, 2, 3, 4}),
				NewIntSetFromInts([]int{0}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{1, 2, 4}),
				NewIntSetFromInts([]int{0, 3}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{1, 3}),
				NewIntSetFromInts([]int{0, 2, 4}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{1, 3, 4}),
				NewIntSetFromInts([]int{0, 2}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{1, 4}),
				NewIntSetFromInts([]int{0, 2, 3}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{2}),
				NewIntSetFromInts([]int{0, 1, 3, 4}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{2, 3}),
				NewIntSetFromInts([]int{0, 1, 4}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{2, 3, 4}),
				NewIntSetFromInts([]int{0, 1}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{2, 4}),
				NewIntSetFromInts([]int{0, 1, 3}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{3}),
				NewIntSetFromInts([]int{0, 1, 2, 4}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{3, 4}),
				NewIntSetFromInts([]int{0, 1, 2}),
			}), NewIntSetsFromIntSetSlice([]IntSet{
				NewIntSetFromInts([]int{4}),
				NewIntSetFromInts([]int{0, 1, 2, 3}),
			}),
		})
}
