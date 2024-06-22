package solvers

import (
	"math"
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

func Distance(a, b Point) float64 {
	return math.Sqrt((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y))
}

type Solver interface {
	Solve() [][]int
}

// TotalDistance calculates the total distance in minutes for one route
func TotalDistance(route []int, vectors []Vector) float64 {
	totalDistance := Distance(origin, vectors[0].Start)
	for _, index := range route {
		totalDistance += Distance(vectors[index].Start, vectors[index].End)
	}
	return totalDistance + Distance(vectors[len(vectors)-1].End, origin)

}

// Cost calculates the cost of the solution
func Cost(routes [][]int, vectors []Vector) float64 {
	totalDistance := 0.
	for _, route := range routes {
		totalDistance += TotalDistance(route, vectors)
	}
	return RouteCostMultiplier*float64(len(routes)) + totalDistance
}
