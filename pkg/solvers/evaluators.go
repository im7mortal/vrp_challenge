package solvers

import (
	"math"
	"math/rand/v2"
)

func GetTheBestByLengthAndCostMin(vectors []*Vector) Evaluator {

	return func(results [][]int) []int {
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

}
func GetTheBestByLengthAndCostMax(vectors []*Vector) Evaluator {

	return func(results [][]int) []int {
		l := 0

		for _, r := range results {
			if len(r) > l {
				l = len(r)
			}
		}
		var maxCost = 0.
		var out []int

		for _, r := range results {
			if len(r) == l {
				t := TotalDistance(r, vectors)
				if t > maxCost {
					maxCost = t
					out = r
				}
			}
		}
		//fmt.Printf("%v\n", out)
		return out
	}

}

func GetTheBestByLengthAndRandom(rnd *rand.Rand) Evaluator {

	return func(results [][]int) []int {
		l := 0

		for _, r := range results {
			if len(r) > l {
				l = len(r)

			}
		}
		var filtered []int
		for i, r := range results {
			if len(r) == l {
				filtered = append(filtered, i)
			}
		}

		if len(filtered) == 1 {
			return results[filtered[0]]
		}
		return results[filtered[int(rnd.Int64N(int64(len(filtered))))]]
	}

}
