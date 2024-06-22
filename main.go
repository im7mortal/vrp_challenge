package main

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"os"
	"vorto/vpr/pkg/solvers"
	"vorto/vpr/pkg/solvers/utils"
)

const debug = 9

var rootCmd = &cobra.Command{
	Use:   "solver [ARGUMENT]",
	Short: "VRP solver",
	Long:  `VRP solver finds optimal amount of drivers and routes to deliver all goods with minimal costs `,

	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		points, err := utils.Parse(os.Args[1])
		if err != nil {
			if glog.V(debug) {
				glog.Exit(err)
			}
			os.Exit(1)
		}
		sol := solvers.NewNearestNeighbor(points)
		printToFormat(sol.Solve())
	},
}

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
			if glog.V(debug) {
				glog.Exitf("Recovered from", r)
			}
			os.Exit(1)
		}
	}()
	if err := rootCmd.Execute(); err != nil {
		if glog.V(debug) {
			glog.Exit(err)
		}
		os.Exit(1)
	}
}
