package solvers

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"math"
)

type nearestNeighbor struct {
	vectors            []*Vector
	visited            []bool
	precomputeToOrigin []float64
	precomputeDistance []float64
}

var origin = Point{0, 0}

func NewNearestNeighbor(vectors []*Vector) Solver {
	nn := &nearestNeighbor{}
	nn.precomputeToOrigin = make([]float64, len(vectors))
	nn.precomputeDistance = make([]float64, len(vectors))
	nn.visited = make([]bool, len(vectors))
	nn.vectors = vectors
	for i := range vectors {
		nn.precomputeToOrigin[i] = Distance(nn.vectors[i].End, origin)
		nn.precomputeDistance[i] = Distance(nn.vectors[i].Start, nn.vectors[i].End)
		//fmt.Printf("%.2f\t%.2f\n", nn.precomputeDistance[i], nn.precomputeToOrigin[i])
	}
	return nn
}

func (nn *nearestNeighbor) findNearestVector(current Point) (int, float64) {
	nearestIndex := -1
	minDistance := math.MaxFloat64
	for i, v := range nn.vectors {
		if !nn.visited[i] {
			d := Distance(current, v.Start)
			if d < minDistance {
				minDistance = d
				nearestIndex = i
			}
		}
	}
	return nearestIndex, minDistance
}

func (nn *nearestNeighbor) Solve(ctx context.Context) ([][]int, error) {
	var solution [][]int
	for {
		current := origin
		var route []int
		totalDistance := 0.0

		for {
			nearestIndex, distToNearest := nn.findNearestVector(current)

			//if nearestIndex != -1 {
			//	if glog.V(1) {
			//		d := totalDistance + distToNearest + nn.precomputeDistance[nearestIndex] + nn.precomputeToOrigin[nearestIndex]
			//		fmt.Printf("%.2f\t%.2f\t%t\t \n", d, d, d > RouteMaxShiftMinutes)
			//	}
			//}
			if nearestIndex == -1 || totalDistance+distToNearest+nn.precomputeDistance[nearestIndex]+nn.precomputeToOrigin[nearestIndex] > RouteMaxShiftMinutes {
				break
			}

			nearestVector := nn.vectors[nearestIndex]
			totalDistance += distToNearest + nn.precomputeDistance[nearestIndex]

			if glog.V(debugLvl) {
				fmt.Printf("%.2f\t%.2f\t%.2f\t \n", distToNearest, nn.precomputeDistance[nearestIndex], totalDistance)
			}

			current = nearestVector.End
			route = append(route, nearestIndex)
			nn.visited[nearestIndex] = true
		}

		solution = append(solution, route)

		// Check if all vectors have been visited
		allVisited := true
		for _, v := range nn.visited {
			if !v {
				allVisited = false
				break
			}
		}
		if allVisited {
			break
		}
	}

	return solution, nil
}
