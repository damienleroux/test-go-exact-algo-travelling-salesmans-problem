package computeTour

import (
	"math"

	algo "github.com/damienleroux/test-go-exact-algo-travelling-salesman-problem/computeTour"
)

func createPossibleSalesSets(
	setsChan chan<- IntSets,
	indexBase int,
	combinaisonsBases IntSets,
	salesNumber int,
	safeGuard chan int,
	maxGoroutines int,
	onEnd algo.CallbackEnd,
) {
	for index := indexBase; index < salesNumber; index++ {
		otherCombinationBeforeIndex := NewIntSetFromIntSet(combinaisonsBases.GetSet(1))
		otherCombinationBeforeIndex.AppendRangeInt(indexBase, index-1)

		otherCombination := NewIntSetFromIntSet(otherCombinationBeforeIndex)
		otherCombination.AppendRangeInt(index+1, salesNumber-1)

		newCombinaison := NewIntSetFromIntSet(combinaisonsBases.GetSet(0))
		newCombinaison.AppendInt(index)

		setsChan <- NewIntSetsFromIntSetSlice([]IntSet{newCombinaison, otherCombination})

		nextIndex := index + 1
		nextCombinaisonsBases := NewIntSetsFromIntSetSlice([]IntSet{newCombinaison, otherCombinationBeforeIndex})

		iterate := func(release algo.CallbackEnd) {
			createPossibleSalesSets(setsChan, nextIndex, nextCombinaisonsBases, salesNumber, safeGuard, maxGoroutines, release)
		}

		if len(safeGuard) < maxGoroutines {
			safeGuard <- 1
			releaseSafeGuard := func() { <-safeGuard }
			go iterate(releaseSafeGuard)
		} else {
			iterate(nil)
		}
	}

	if onEnd != nil {
		onEnd()
	}
}

// returns every possible combinaison when dividing X Sales into X bundle
func GetPossibleSaleIndexesSets(salesNumber, maxGoroutines int) []IntSets {
	safeGuard := make(chan int, maxGoroutines)
	setsChan := make(chan IntSets)

	// Create sets
	initialCombinaison := newIntSets()
	initialCombinaison.AppendIntSet(IntSet{})
	initialCombinaison.AppendIntSet(IntSet{})
	go createPossibleSalesSets(setsChan, 0, initialCombinaison, salesNumber, safeGuard, maxGoroutines, nil)

	// Find best route
	numberOfSolutions := int(math.Pow(2, float64(salesNumber)) - 1)
	possibleSaleIndexesSets := make([]IntSets, numberOfSolutions)
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

func findCombinaisonRoute(routes []SetRoute, salerSet IntSet) algo.Route {
	for _, setRoute := range routes {
		found := false
		for _, index := range setRoute.Set.Get() {
			for _, indexToCompare := range salerSet.Get() {
				if index == indexToCompare {
					found = true
				}
			}
		}
		if found && setRoute.Set.Len() == salerSet.Len() {
			return setRoute.Route
		}
	}
	return algo.Route{}
}

// Compute the best route to link a number of sales, each sale being represented by geo coordinates.
func GetBestRoute(sales []algo.Sale, maxGoroutines int) RoutesCombination {
	salesLen := len(sales)
	possibleSaleIndexesSets := GetPossibleSaleIndexesSets(salesLen, maxGoroutines)

	var combinaisons = make([]SetRoute, len(possibleSaleIndexesSets))

	// compute route for each set
	for i, possibleSaleIndexesSet := range possibleSaleIndexesSets {
		salerSet := possibleSaleIndexesSet.GetSet(0)
		salesEntry := make([]algo.Sale, salerSet.Len())
		for j, salerIndex := range salerSet.Get() {
			salesEntry[j] = sales[salerIndex]
		}
		salesRoute := algo.GetBestRoute(salesEntry, maxGoroutines)
		combinaisons[i] = SetRoute{salerSet, salesRoute}
	}

	var bestRoutesCombination RoutesCombination
	for _, possibleSaleIndexesSet := range possibleSaleIndexesSets {
		routesCombination := RoutesCombination{[]algo.Route{}, 0}
		for _, salerSet := range possibleSaleIndexesSet.Get() {
			route := findCombinaisonRoute(combinaisons, salerSet)
			routesCombination.Routes = append(routesCombination.Routes, route)
			routesCombination.TotalCoveredDistance += route.Steps[len(route.Steps)-1].TotalCoveredDistance
		}

		if routesCombination.TotalCoveredDistance > bestRoutesCombination.TotalCoveredDistance {
			bestRoutesCombination = routesCombination
		}
	}

	return bestRoutesCombination
}
