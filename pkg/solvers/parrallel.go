package solvers

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"math"
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
	index  int
}

func createJobFromFuncOut(result [][]int, err error) jobResult {
	return jobResult{
		result: result,
		err:    err,
	}
}

func (pl *parallel) Solve(ctx context.Context) ([][]int, error) {

	resultChan := make(chan jobResult)
	counter := 0
	for i := 0; i < pl.N; i++ {
		for funcIndex := range pl.evaluatorFuncs {
			go func(n, fIndex int) {
				j := jobResult{index: n}
				defer func() {
					if r := recover(); r != nil {
						j.err = fmt.Errorf("Recovered from %s", r)
						if glog.V(errorLvl) {
							glog.Error(j.err)
						}
						// correctly finish the program
						resultChan <- j
					}
				}()
				slv := NewNearestNeighborExp(pl.vectors, pl.evaluatorFuncs[funcIndex], n)
				j.result, j.err = slv.Solve(ctx)
				resultChan <- j
			}(i+1, funcIndex)
			counter++
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
			// TODO I don't like counter pattern
			counter--
			if counter == 0 {
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
