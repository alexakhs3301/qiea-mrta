package ga

import "math/rand"

type Individual struct {
	Assignments [][]int
	Fitness     float64
}

func NewIndividual(m, n int) *Individual {
	assignments := make([][]int, m)
	for i := range assignments {
		assignments[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if rand.Float64() < 0.5 {
				assignments[i][j] = 1
			} else {
				assignments[i][j] = 0
			}
		}
	}
	return &Individual{Assignments: assignments}
}

func (ind *Individual) DeepCopy() *Individual {
	copyAssignments := make([][]int, len(ind.Assignments))
	for i := range ind.Assignments {
		copyAssignments[i] = append([]int{}, ind.Assignments[i]...)
	}
	return &Individual{
		Assignments: copyAssignments,
		Fitness:     ind.Fitness,
	}
}
