package main

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"os"
	"time"
	"vorto/vpr/pkg/solvers"
	"vorto/vpr/pkg/solvers/utils"
)

const (
	errorLvl = 1
	debugLvl = 9
)

const (
	mustExit      = 30 * time.Second
	mustFinishAll = 28 * time.Second
)

// there is high probability that during development we stack in infinite loop
// if we use all cores then it can parallelize UI; I have Ubuntu and it happened with me many times
func init() {
	time.AfterFunc(mustExit, func() {
		err := fmt.Errorf("Time exceeded: took more then %d seconds", mustExit/time.Second)
		if glog.V(errorLvl) {
			glog.Exit(err)
		}
		os.Exit(1)
	})
}

var rootCmd = &cobra.Command{
	Use:   "solver [ARGUMENT]",
	Short: "VRP solver",
	Long:  `VRP solver finds optimal amount of drivers and routes to deliver all goods with minimal costs `,

	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		vectors, err := utils.Parse(os.Args[1])
		if err != nil {
			if glog.V(errorLvl) {
				glog.Exit(err)
			}
			os.Exit(1)
		}

		ctx, _ := context.WithTimeout(context.Background(), mustFinishAll)
		sol := solvers.NewParallel(vectors, 5)

		result, err := sol.Solve(ctx)
		if r := recover(); r != nil {
			if glog.V(errorLvl) {
				glog.Exitf("Recovered from panic:\n", r)
			}
			os.Exit(1)
		}
		printToFormat(result)
	},
}

// python expects following format
// [0,1,2]
func printToFormat(results [][]int) {
	for _, entrance := range results {
		s := "["
		for i, e := range entrance {
			if i > 0 {
				s += ", "
			}
			// '+1' because the numeration of the routes doesn't include 0
			s += fmt.Sprintf("%d", e+1)
		}
		fmt.Println(s + "]")
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if glog.V(errorLvl) {
				glog.Exitf("Recovered from:", r)
			}
			os.Exit(1)
		}
	}()
	if err := rootCmd.Execute(); err != nil {
		if glog.V(errorLvl) {
			glog.Exit(err)
		}
		os.Exit(1)
	}
}
