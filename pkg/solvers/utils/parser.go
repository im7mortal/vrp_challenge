package utils

import (
	"encoding/csv"
	"fmt"
	"github.com/golang/glog"
	"os"
	"vorto/vpr/pkg/solvers"
)

const (
	errorLvl = 1
	debugLvl = 9
)

func Parse(filepath string) ([]*solvers.Vector, error) {
	file, err := os.Open(filepath)
	if err != nil {
		if glog.V(errorLvl) {
			glog.Error(err)
		}
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ' ' // we have whitespace separated document

	values, err := reader.ReadAll()

	if glog.V(debugLvl) {
		glog.Infof("Content of %s:", filepath)
		glog.Infof("%v", values)
	}

	if err != nil {
		if glog.V(errorLvl) {
			glog.Error(err)
		}
		return nil, err
	}

	// remove header
	values = values[1:]

	vectors := make([]*solvers.Vector, len(values))
	// init vectors as pointer can't be inited automatically
	for i := range vectors {
		vectors[i] = &solvers.Vector{}
	}

	for i := range values {
		if len(values[i]) != 3 {
			err = fmt.Errorf("the row must have 3 columns, got %d", len(values[i]))
			if glog.V(errorLvl) {
				glog.Error(err)
			}
			return nil, err
		}
		_, err = fmt.Sscanf(values[i][1], "(%f,%f)", &vectors[i].Start.X, &vectors[i].Start.Y)
		if err != nil {
			if glog.V(errorLvl) {
				glog.Error(err)
			}
			return nil, err
		}
		_, err = fmt.Sscanf(values[i][2], "(%f,%f)", &vectors[i].End.X, &vectors[i].End.Y)
		if err != nil {
			if glog.V(errorLvl) {
				glog.Error(err)
			}
			return nil, err
		}
	}
	return vectors, nil
}
