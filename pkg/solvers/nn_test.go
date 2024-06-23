package solvers_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"
	"vorto/vpr/pkg/solvers"
	"vorto/vpr/pkg/solvers/utils"
)

func init() {
	// there is high probability that during development we stack in infinite loop on all cores
	// if we use all cores then it can parallelize UI; I have Ubuntu and it happened with me many times
	time.AfterFunc(30*time.Second, func() {
		os.Exit(0)
	})
}

// !!!
// !!! these values are only for rush developing as in real scenario we use the MEAN baseline
const (
	BaselinePetrFirstIteration = 100. * 1000.
	BaselineSimple             = 49197.9372133668
	BaselineStandard           = 45300.4127368728
)

var targetBaseline = BaselinePetrFirstIteration

// I prefer to have test cases in list instead of dynamically parsing directory
// in this case I can exclude quickly test cases
// for example problem5.txt and problem6.txt will stack with N salesman algorithm with N equal >2
var files = []string{
	"../../Training Problems/problem1.txt",
	"../../Training Problems/problem2.txt",
	"../../Training Problems/problem3.txt",
	"../../Training Problems/problem4.txt",
	"../../Training Problems/problem5.txt",
	"../../Training Problems/problem6.txt",
	"../../Training Problems/problem7.txt",
	"../../Training Problems/problem8.txt",
	"../../Training Problems/problem9.txt",
	"../../Training Problems/problem10.txt",
	"../../Training Problems/problem11.txt",
	"../../Training Problems/problem12.txt",
	"../../Training Problems/problem13.txt",
	"../../Training Problems/problem14.txt",
	"../../Training Problems/problem15.txt",
	"../../Training Problems/problem16.txt",
	"../../Training Problems/problem17.txt",
	"../../Training Problems/problem18.txt",
	"../../Training Problems/problem19.txt",
	"../../Training Problems/problem20.txt",
}

func TestSolver(t *testing.T) {
	vectors, err := utils.Parse("../../Training Problems/problem17.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	nn := solvers.NewNearestNeighborExp(vectors, solvers.GetTheBestByLengthAndCostMin(vectors), 5)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*29)
	routes, err := nn.Solve(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	fmt.Printf("%v\n", routes)

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

func TestSolverParallel(t *testing.T) {
	fn := "../../Training Problems/problem10.txt"
	vectors, err := utils.Parse(fn)
	if err != nil {
		t.Errorf(err.Error())
	}
	rnd, err := utils.NewRandFactory(fn)
	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*29)
	var N = 5
	nn := solvers.NewParallel(vectors, []solvers.Evaluator{
		solvers.GetTheBestByLengthAndCostMin(vectors),
		solvers.GetTheBestByLengthAndRandom(rnd.GetRandomGenerator()),
	}, N)
	routes, err := nn.Solve(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	fmt.Printf("%v\n", routes)

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

func getVectors(t *testing.T) [][]*solvers.Vector {
	vectors := [][]*solvers.Vector{}
	for _, f := range files {
		v, err := utils.Parse(f)
		if err != nil {
			t.Errorf(err.Error())
		}
		vectors = append(vectors, v)
	}
	return vectors
}

func TestAll(t *testing.T) {
	vectors := getVectors(t)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*29)
	var N = 5
	var values []float64
	for _, vectorss := range vectors {
		nn := solvers.NewParallel(vectorss, []solvers.Evaluator{solvers.GetTheBestByLengthAndCostMin(vectorss)}, N)
		//ctx, _ := context.WithTimeout(context.Background(), time.Second*29)
		result, err := nn.Solve(ctx)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		values = append(values, solvers.Cost(result, vectorss))
	}
	total := 0.0
	for _, value := range values {
		total += value
	}
	mean := total / float64(len(values))
	fmt.Println("Mean of values:", mean)

}

func TestGetRandom(t *testing.T) {
	fn := "../../Training Problems/problem17.txt"
	var checkRandom int
	for i := 0; i < 10; i++ {

		rnd, err := utils.NewRandFactory(fn)
		if err != nil {
			t.Errorf("Error: %s", err)
			return
		}
		if i == 0 {
			checkRandom = rnd.GetRandomGenerator().Int()
		} else {
			if v := rnd.GetRandomGenerator().Int(); checkRandom != v {
				t.Errorf("Random generator doesn't generate consistent output: %d but must %d", v, checkRandom)
				return
			} else {
				println(rnd.GetRandomGenerator().Int())
			}
		}
	}
}
