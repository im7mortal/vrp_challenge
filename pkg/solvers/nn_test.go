package solvers_test

import (
	"fmt"
	"testing"
	"vorto/vpr/pkg/solvers"
	"vorto/vpr/pkg/solvers/utils"
)

// !!!
// !!! these values are only for rush developing as in real scenario we use the MEAN baseline
const (
	BaselinePetrFirstIteration = 100. * 1000.
	BaselineSimple             = 49197.9372133668
	BaselineStandard           = 45300.4127368728
)

var targetBaseline = BaselinePetrFirstIteration

func TestSolver(t *testing.T) {
	vectors, err := utils.Parse("../../problems/problem16.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	nn := solvers.NewNearestNeighbor(vectors)
	routes := nn.Solve()

	for _, route := range routes {
		if d := solvers.TotalDistance(route, vectors); d > solvers.RouteMaxShiftMinutes {
			t.Errorf("Route %v exceed 12 hours shift with duration %f", route, d)
		}
	}

	if cost := solvers.Cost(routes, vectors); cost > targetBaseline {
		t.Errorf("Target cost %f was exced %f", targetBaseline, cost)
	} else {
		fmt.Printf("Target cost %f was %f\n", targetBaseline, cost)
	}
}
