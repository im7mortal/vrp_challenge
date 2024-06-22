package main

import (
	"encoding/csv"
	"fmt"
	"github.com/golang/glog"
	"os"
	"vorto/vpr/pkg/solvers"
)

func main() {

	file, err := os.Open(os.Args[1])
	if err != nil {
		glog.Errorf("Error opening file: %i", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ' '

	values, err := reader.ReadAll()

	if err != nil {
		glog.Errorf("Error parsing csv: %i", err)
		os.Exit(1)
	}
	values = values[1:]

	f := make([][4]float64, len(values))

	for i := range values {
		if len(values[i]) != 3 {
			glog.Errorf("Error parsing csv: %i", values[i])
			os.Exit(1)
		}
		_, err = fmt.Sscanf(values[i][1], "(%f,%f)", &(f[i][0]), &(f[i][1]))
		if err != nil {
			glog.Errorf("Error parsing csv: %i", err)
			os.Exit(1)
		}
		_, err = fmt.Sscanf(values[i][1], "(%f,%f)", &(f[i][2]), &(f[i][3]))
		if err != nil {
			glog.Errorf("Error parsing csv: %i", err)
			os.Exit(1)
		}
		//fmt.Printf("[%d]\n", i+1)
	}
	sol := solvers.Solve1(f)

	for _, entrance := range sol {
		fmt.Print("[")
		for i, e := range entrance {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(e)
		}
		fmt.Println("]")
	}

}
