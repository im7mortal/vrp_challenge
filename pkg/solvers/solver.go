package solvers

func Solve1(vals [][4]float64) (sol [][]int) {

	for i := range vals {
		sol = append(sol, []int{i + 1})
	}
	return sol
}
