package solvers

import "math"

func GetTheBestByLengthAndCost(vectors []*Vector, results [][]int) []int {
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
			t := TotalDistance(r, vectors)
			if t < minCost {
				minCost = t
				out = r
			}
		}
	}
	//fmt.Printf("%v\n", out)
	return out
}
