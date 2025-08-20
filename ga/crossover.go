package ga

import "math/rand"

// UniformCrossover performs uniform crossover between two parents
func UniformCrossover(parent1, parent2 *Individual) *Individual {
	m := len(parent1.Assignments)
	n := len(parent1.Assignments[0])

	child := &Individual{
		Assignments: make([][]int, m),
	}

	for i := 0; i < m; i++ {
		child.Assignments[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if rand.Float64() < 0.5 {
				child.Assignments[i][j] = parent1.Assignments[i][j]
			} else {
				child.Assignments[i][j] = parent2.Assignments[i][j]
			}
		}
	}

	return child
}
