package solvers

import (
	"context"
	"github.com/golang/glog"
	"math"
	"os"
)

type parallel struct {
	N              int
	vectors        []*Vector
	evaluatorFuncs []Evaluator
}

func NewParallel(vectors []*Vector, evaluatorFuncs []Evaluator, N int) Solver {
	if len(evaluatorFuncs) == 0 {
		evaluatorFuncs = []Evaluator{GetTheBestByLengthAndCost(vectors)}
	}
	return &parallel{N: N, vectors: vectors, evaluatorFuncs: evaluatorFuncs}
}

type jobResult struct {
	result [][]int
	err    error
}

func createJobFromFuncOut(result [][]int, err error) jobResult {
	return jobResult{
		result: result,
		err:    err,
	}
}

func (pl *parallel) Solve(ctx context.Context) ([][]int, error) {

	resultChan := make(chan jobResult)
	for i := 0; i < pl.N; i++ {
		for funcIndex := range pl.evaluatorFuncs {
			go func(slv Solver) {
				if r := recover(); r != nil {
					if glog.V(errorLvl) {
						glog.Exitf("Recovered from", r)
					}
					// correctly finish the program
					os.Exit(1)
				}
				resultChan <- createJobFromFuncOut(slv.Solve(ctx))
			}(NewNearestNeighborExp(pl.vectors, pl.evaluatorFuncs[funcIndex], i+1))
		}
	}
	var results []jobResult
	finished := false
	for {
		if finished {
			break
		}
		select {
		case job := <-resultChan:
			results = append(results, job)
			if len(results) == pl.N {
				finished = true
			}
		case <-ctx.Done():
			finished = true
		}
	}

	var finalResult [][]int
	var minimalCost = math.MaxFloat64

	for _, r := range results {
		if r.err != nil {
			// TODO handler
			continue
		}
		if Cost(r.result, pl.vectors) < minimalCost {
			finalResult = r.result
		}
	}
	return finalResult, nil
}
