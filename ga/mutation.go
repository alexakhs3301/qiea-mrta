package ga

import "math/rand"

// Mutate applies bit-flip mutation to the individual's assignments
func (ind *Individual) Mutate(mutationRate float64) {
	m := len(ind.Assignments)
	n := len(ind.Assignments[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rand.Float64() < mutationRate {
				ind.Assignments[i][j] = 1 - ind.Assignments[i][j] // flip bit
			}
		}
	}
}
