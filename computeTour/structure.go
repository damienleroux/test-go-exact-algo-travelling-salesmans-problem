package computeTour

import algo "github.com/damienleroux/test-go-exact-algo-travelling-salesman-problem/computeTour"

type SetRoute struct {
	Set   IntSet
	Route algo.Route
}

type RoutesCombination struct {
	Routes               []algo.Route
	TotalCoveredDistance float64
}
