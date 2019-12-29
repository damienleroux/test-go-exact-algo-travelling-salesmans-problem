package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isSetInSets(setToFind []int, sets [][]int) bool {
	for _, set := range sets {
		//only set with same number of element are concerned
		if len(set) == len(setToFind) {
			for i, setIndex := range set {
				indexFound := false
				for _, setToFindIndex := range setToFind {
					if setToFindIndex == setIndex {
						indexFound = true
					}
				}
				if !indexFound {
					break
				}
				//if every indexes in set are verified without breaking
				if i == len(set)-1 {
					//the indexes set has been found
					return true
				}
			}
		}
	}
	return false
}

func areSaleSetsInSets(SalesSetToFind [][]int, SalesSets [][][]int) bool {
	for _, SalesSet := range SalesSets {
		//only set with same number of element are concerned
		if len(SalesSet) == len(SalesSetToFind) {
			for i, SaleSet := range SalesSet {
				setFound := isSetInSets(SaleSet, SalesSet)
				if !setFound {
					break
				}
				//if every indexes in set are verified without breaking
				if i == len(SalesSet)-1 {
					//the indexes set has been found
					return true
				}
			}
		}
	}
	return false
}

func testGetPossibleSaleIndexesSets(t *testing.T, salesNumber int, expectedSets [][][]int) {
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
		[][][]int{
			{
				{0},    //sale 1 one can handle first command
				{1, 2}, //while the other one can handle the toher ones
			}, {
				{0, 1},
				{2},
			}, {
				{0, 1, 2},
				{},
			}, {
				{0, 2},
				{1},
			}, {
				{1},
				{0, 2},
			}, {
				{1, 2},
				{0},
			}, {
				{2},
				{0, 1},
			},
		})
}

func TestGetPossibleSaleIndexesSets5(t *testing.T) {
	testGetPossibleSaleIndexesSets(
		t,
		//number of sales
		5,
		//expected set
		[][][]int{
			{
				{0},
				{1, 2, 3, 4},
			}, {
				{0, 1},
				{2, 3, 4},
			}, {
				{0, 1, 2},
				{3, 4},
			}, {
				{0, 1, 2, 3},
				{4},
			}, {
				{0, 1, 2, 3, 4},
				{},
			}, {
				{0, 1, 2, 4},
				{3},
			}, {
				{0, 1, 3},
				{2, 4},
			}, {
				{0, 1, 3, 4},
				{2},
			}, {
				{0, 1, 4},
				{2, 3},
			}, {
				{0, 2},
				{1, 3, 4},
			}, {
				{0, 2, 3},
				{1, 4},
			}, {
				{0, 2, 3, 4},
				{1},
			}, {
				{0, 2, 4},
				{1, 3},
			}, {
				{0, 3},
				{1, 2, 4},
			}, {
				{0, 3, 4},
				{1, 2},
			}, {
				{0, 4},
				{1, 2, 3},
			}, {
				{1},
				{0, 2, 3, 4},
			}, {
				{1, 2},
				{0, 3, 4},
			}, {
				{1, 2, 3},
				{0, 4},
			}, {
				{1, 2, 3, 4},
				{0},
			}, {
				{1, 2, 4},
				{0, 3},
			}, {
				{1, 3},
				{0, 2, 4},
			}, {
				{1, 3, 4},
				{0, 2},
			}, {
				{1, 4},
				{0, 2, 3},
			}, {
				{2},
				{0, 1, 3, 4},
			}, {
				{2, 3},
				{0, 1, 4},
			}, {
				{2, 3, 4},
				{0, 1},
			}, {
				{2, 4},
				{0, 1, 3},
			}, {
				{3},
				{0, 1, 2, 4},
			}, {
				{3, 4},
				{0, 1, 2},
			}, {
				{4},
				{0, 1, 2, 3},
			}})
}
