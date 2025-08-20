package quantum

import (
	"math"
	"math/rand"
)

func (ind *Individual) Observe(m, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			idx := i*n + j
			if rand.Float64() < math.Sin(ind.Theta[idx])*math.Sin(ind.Theta[idx]) {
				ind.Assignments[i][j] = 1
			} else {
				ind.Assignments[i][j] = 0
			}
		}
	}
}
