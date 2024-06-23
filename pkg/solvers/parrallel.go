package solvers

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"math"
	"sync"
	"time"
)

type parallel struct {
	N              int
	vectors        []*Vector
	evaluatorFuncs []Evaluator
}

func NewParallel(vectors []*Vector, evaluatorFuncs []Evaluator, N int) Solver {
	if len(evaluatorFuncs) == 0 {
		evaluatorFuncs = []Evaluator{GetTheBestByLengthAndCostMax(vectors)}
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

// it's not hardcode! it's KNOWHOW
// visit "Algorithm" section in the README.md
const magicKNOWHOWTimeout = 5 * time.Second
const magicKNOWHOW_N_number = 3

func (pl *parallel) Solve(ctx context.Context) ([][]int, error) {

	// sloppy logic to ensure that we do not stack on cases which generate to deep recursion
	ctx, cancel := context.WithCancel(ctx)
	chanN3Done := make(chan jobResult)
	var closeOnce sync.Once
	passedCancelThreshold := func() {
		closeOnce.Do(func() {
			close(chanN3Done)
		})
	}
	go func() {
		innerCtx, _ := context.WithTimeout(ctx, magicKNOWHOWTimeout)
		select {
		case <-innerCtx.Done():
			if glog.V(debugLvl) {
				glog.Infof("Deep recursion detected; cancel calculations")
			}
			// let's do not wait on rest calculation; probably it's deep recursion
			cancel()
		case <-chanN3Done:
			//disable cancel logic
			return
		}
	}()

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
			if job.index == magicKNOWHOW_N_number {
				if glog.V(debugLvl) {
					glog.Infof("Deep recursion detected; send signal to cancel calculations")
				}
				passedCancelThreshold()
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
