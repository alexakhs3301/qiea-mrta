package ga

import "Code/common"

func (ind *Individual) EvaluateFitness(
	robots []common.Robot,
	tasks []common.Task,
	utility [][]int,
	costs [][]int,
	lambdaCap, lambdaCoop, lambdaCapCost, lambdaCoopCost float64,
	w1, w2 float64,
) {
	m := len(robots)
	n := len(tasks)
	var util, cost float64
	var vCap, vCoop float64

	// Compute utility and cost
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if ind.Assignments[i][j] == 1 {
				util += float64(utility[i][j])
				cost += float64(costs[i][j])
			}
		}
	}

	// Capacity Violation
	for i := 0; i < m; i++ {
		var sum int
		for j := 0; j < n; j++ {
			if ind.Assignments[i][j] == 1 {
				sum += costs[i][j]
			}
		}
		excess := sum - robots[i].Capacity
		if excess > 0 {
			vCap += float64(excess)
		}
	}

	// Cooperation Violation
	for j := 0; j < n; j++ {
		var assigned int
		for i := 0; i < m; i++ {
			assigned += ind.Assignments[i][j]
		}
		missing := tasks[j].RequiredUnits - assigned
		if missing > 0 {
			vCoop += float64(missing)
		}
	}

	penalty := lambdaCap*vCap + lambdaCoop*vCoop + lambdaCapCost*vCap + lambdaCoopCost*vCoop
	ind.Fitness = (w1 * util) - (w2 * cost) - penalty
}
