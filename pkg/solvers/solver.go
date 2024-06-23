package solvers

import (
	"context"
	"fmt"
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
	Solve(ctx context.Context) ([][]int, error)
}

type Evaluator func([]*Vector, [][]int) []int

// TotalDistance calculates the total distance in minutes for one route
func TotalDistance(route []int, vectors []*Vector) float64 {
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

// for debug purposes; exactly matching with python
func pythonStyleArrayPrint(i []int) string {
	ss := "["
	for g, d := range i {
		if g != 0 {
			ss += ", "
		}
		ss += fmt.Sprintf("'%d'", d+1)
	}

	return ss + "]"
}

// Cost calculates the cost of the solution
func Cost(routes [][]int, vectors []*Vector) float64 {
	totalDistance := 0.
	for _, route := range routes {

		//fmt.Printf("%s\n", pythonStyleArrayPrint(route))
		//fmt.Printf("%.12f\n", TotalDistance(route, vectors))

		totalDistance += TotalDistance(route, vectors)
	}
	//println(RouteCostMultiplier * float64(len(routes)))
	//println(totalDistance)
	return RouteCostMultiplier*float64(len(routes)) + totalDistance
}
