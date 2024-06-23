package solvers

import (
	"context"
	"math"
	"sort"
)

type nearestNeighborExp struct {
	vectors            []*Vector
	visited            []bool
	precomputeToOrigin []float64
	precomputeDistance []float64
	N                  int
}

func NewNearestNeighborExp(vectors []*Vector, N int) Solver {
	nn := &nearestNeighborExp{N: N}
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

type neighbour struct {
	index    int
	distance float64
}

type neighbours []neighbour

func (n neighbours) Len() int { return len(n) }

func (n neighbours) Less(i, j int) bool { return n[i].distance < n[j].distance }

func (n neighbours) Swap(i, j int) { n[i], n[j] = n[j], n[i] }

func exist(i []int, v int) bool {
	for j := range i {
		if i[j] == v {
			return true
		}
	}
	return false
}

func (nn *nearestNeighborExp) find_N_NearestVector(current Point, visitedByCurrentIteration []int) neighbours {
	var ns neighbours
	for i, v := range nn.vectors {
		if !exist(visitedByCurrentIteration, i) && !nn.visited[i] {
			ns = append(ns, neighbour{
				index:    i,
				distance: Distance(current, v.Start), // TODO MUST BE MATRIX
			})
		}
	}
	sort.Sort(ns)
	if len(ns) > nn.N {
		return ns[:nn.N]
	}
	return ns
}

func (nn *nearestNeighborExp) salesmanRecursion(sequence []int, current Point, totalDistance float64) [][]int {

	var result [][]int

	nearestPoints := nn.find_N_NearestVector(current, sequence)
	if len(nearestPoints) == 0 {
		return [][]int{sequence}
	}
	//println(len(nearestPoints))
	appendedInCurrentIteration := false
	for _, nearestPoint := range nearestPoints {

		if totalDistance+nearestPoint.distance+nn.precomputeDistance[nearestPoint.index]+nn.precomputeToOrigin[nearestPoint.index] > RouteMaxShiftMinutes {
			if !appendedInCurrentIteration {
				result = append(result, sequence)
				appendedInCurrentIteration = true
			}
			continue
		}

		sequenceCopy := make([]int, len(sequence))
		copy(sequenceCopy, sequence)
		sequenceCopy = append(sequenceCopy, nearestPoint.index)
		//fmt.Printf("res %v\n", sequenceCopy)
		//time.Sleep(time.Second)

		res := nn.salesmanRecursion(sequenceCopy, nn.vectors[nearestPoint.index].End, totalDistance+nearestPoint.distance+nn.precomputeDistance[nearestPoint.index])

		result = append(result, res...)

	}
	return result

}

func (nn *nearestNeighborExp) getTheBestResult(results [][]int) []int {
	l := 0

	for _, r := range results {
		if len(r) > l {
			l = len(r)
		}
	}
	var minCost = math.MaxFloat64
	var out []int

	for _, r := range results {
		if len(r) == l {
			t := TotalDistance(r, nn.vectors)
			if t < minCost {
				minCost = t
				out = r
			}
		}
	}
	//fmt.Printf("%v\n", out)
	return out
}

func (nn *nearestNeighborExp) Solve(ctx context.Context) ([][]int, error) {
	var solution [][]int
	for {
		results := nn.salesmanRecursion([]int{}, origin, 0.0)

		route := nn.getTheBestResult(results)

		for _, index := range route {
			nn.visited[index] = true
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
