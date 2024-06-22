package solvers

import (
	"math"
)

const (
	errorLvl = 1
	debugLvl = 9
)

const (
	RouteCostMultiplier  = 500.
	RouteMaxShiftMinutes = 12. * 60.
)

type Point struct {
	X, Y float64
}

type Vector struct {
	Start, End Point
}

func Distance(p1, p2 Point) float64 {
	xDiff := p1.X - p2.X
	yDiff := p1.Y - p2.Y
	return math.Sqrt(xDiff*xDiff + yDiff*yDiff)
}

type Solver interface {
	Solve() [][]int
}

// TotalDistance calculates the total distance in minutes for one route
func TotalDistance(route []int, vectors []Vector) float64 {
	localOrigin := origin
	totalDistance := 0.
	for _, index := range route {
		totalDistance += Distance(localOrigin, vectors[index].Start)
		localOrigin = vectors[index].End
		totalDistance += Distance(vectors[index].Start, vectors[index].End)
	}
	totalDistance += Distance(vectors[route[len(route)-1]].End, origin)
	return totalDistance

}

// Cost calculates the cost of the solution
func Cost(routes [][]int, vectors []Vector) float64 {
	totalDistance := 0.
	for _, route := range routes {
		totalDistance += TotalDistance(route, vectors)
	}
	return RouteCostMultiplier*float64(len(routes)) + totalDistance
}
