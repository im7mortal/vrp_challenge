package solvers

import (
	"math"
)

//// we know that negative mustn't be bigger then 12 hours shift
//const MaxFor12HoursShift float64 = 1000
//
//func Normalize(coord *[][4]float64, offset float64) {
//	for i := range *coord {
//		(*coord)[i][0] += offset
//		(*coord)[i][1] += offset
//		(*coord)[i][2] += offset
//		(*coord)[i][3] += offset
//	}
//}

func euclidDistance(coord [4]float64) float64 {
	return math.Sqrt(math.Pow(coord[2]-coord[0], 2) + math.Pow(coord[3]-coord[1], 2))
}

type Solver interface {
	Solve() [][]int
}

type nearestNeighbor struct {
	vals   [][4]float64
	center [2]float64
}

func NewNearestNeighbor(vals [][4]float64, center [2]float64) Solver {
	return &nearestNeighbor{vals, center}
}

func (nn *nearestNeighbor) Solve() (solution [][]int) {

	for i := range nn.vals {
		solution = append(solution, []int{i + 1})
	}

	return solution
}
