package quantum

import (
	"math"
	"math/rand"
)

type Individual struct {
	Theta       []float64
	Assignments [][]int
	Fitness     float64
}

func NewIndividual(m, n int) *Individual {
	size := m * n
	theta := make([]float64, size)
	for i := range theta {
		theta[i] = math.Pi/4 + (rand.Float64()-0.5)*0.1 // we add a small noise to the initial angle
	}
	assignments := make([][]int, m)
	for i := range assignments {
		assignments[i] = make([]int, n)
	}
	return &Individual{
		Theta:       theta,
		Assignments: assignments,
		Fitness:     0.0,
	}
}

func (ind *Individual) DeepCopy() *Individual {
	newTheta := make([]float64, len(ind.Theta))
	copy(newTheta, ind.Theta)

	m := len(ind.Assignments)
	n := len(ind.Assignments[0])
	newAssign := make([][]int, m)
	for i := 0; i < m; i++ {
		newAssign[i] = make([]int, n)
		copy(newAssign[i], ind.Assignments[i])
	}

	return &Individual{
		Theta:       newTheta,
		Assignments: newAssign,
		Fitness:     ind.Fitness,
	}
}
