package main

import (
	"github.com/damienleroux/test-go-exact-algo-travelling-salesman-problem/computeTour"
	"math"
)

func createSet(combinaisonBases []int, i int) []int {
	var newCombinaison []int
	newCombinaison = append(newCombinaison, combinaisonBases...)
	newCombinaison = append(newCombinaison, i)
	return newCombinaison
}

func createPossibleSalesSets(
	setsChan chan<- [][]int,
	indexBase int,
	combinaisonsBases [][]int,
	salesNumber int,
	safeGuard chan int,
	maxGoroutines int,
	onEnd computeTour.CallbackEnd,
) {
	for index := indexBase; index < salesNumber; index++ {
		var otherCombinationBeforeIndex []int
		var otherCombination []int
		otherCombinationBeforeIndex = append(otherCombinationBeforeIndex, combinaisonsBases[1]...)
		//	otherCombination = append(otherCombination, otherCombinationBases...)
		for otherIndex := indexBase; otherIndex < index; otherIndex++ {
			otherCombinationBeforeIndex = append(otherCombinationBeforeIndex, otherIndex)
		}
		otherCombination = append(otherCombination, otherCombinationBeforeIndex...)
		for otherIndex := index + 1; otherIndex < salesNumber; otherIndex++ {
			otherCombination = append(otherCombination, otherIndex)
		}
		newCombinaison := createSet(combinaisonsBases[0], index)

		setsChan <- [][]int{newCombinaison, otherCombination}
		createPossibleSalesSets(setsChan, index+1, [][]int{newCombinaison, otherCombinationBeforeIndex}, salesNumber, safeGuard, maxGoroutines, onEnd)
	}

	if onEnd != nil {
		onEnd()
	}
}

// returns every possible combinaison when dividing X Sales into X bundle
func GetPossibleSaleIndexesSets(salesNumber, maxGoroutines int) [][][]int {
	safeGuard := make(chan int, maxGoroutines)
	setsChan := make(chan [][]int)

	// Create sets
	go createPossibleSalesSets(setsChan, 0, [][]int{{}, {}}, salesNumber, safeGuard, maxGoroutines, nil)

	// Find best route
	numberOfSolutions := int(math.Pow(2, float64(salesNumber)) - 1)
	possibleSaleIndexesSets := make([][][]int, numberOfSolutions)
	setIndex := 0
	for possibleSaleIndexesSet := range setsChan {
		//fmt.Println(possibleSaleIndexesSet[0], " vs ", possibleSaleIndexesSet[1])
		possibleSaleIndexesSets[setIndex] = possibleSaleIndexesSet
		setIndex++
		if setIndex == numberOfSolutions {
			close(setsChan)
		}
	}

	return possibleSaleIndexesSets
}
