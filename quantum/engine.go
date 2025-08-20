package quantum

import (
	"Code/common"
	"math"
)

type Population struct {
	Individuals []*Individual
	GlobalBest  *Individual
	Size        int
	M           int
	N           int
}

func NewPopulation(size, m, n int) *Population {
	individuals := make([]*Individual, size)
	for i := 0; i < size; i++ {
		individuals[i] = NewIndividual(m, n)
	}
	return &Population{
		Individuals: individuals,
		GlobalBest:  nil,
		Size:        size,
		M:           m,
		N:           n,
	}
}

func (pop *Population) Run(
	robots []common.Robot,
	tasks []common.Task,
	utility [][]int,
	costs [][]int,
	generations int,
	lambdaCap, lambdaCoop, lambdaCapCost, lambdaCoopCost float64,
	w1, w2 float64,
	deltaTheta float64,
	fitnessHistory *[]float64,
) {
	for gen := 0; gen < generations; gen++ {

		for _, ind := range pop.Individuals {
			ind.Observe(pop.M, pop.N)
			ind.EvaluateFitness(robots, tasks, utility, costs,
				lambdaCap, lambdaCoop, lambdaCapCost, lambdaCoopCost,
				w1, w2)
		}

		for _, ind := range pop.Individuals {
			if pop.GlobalBest == nil || ind.Fitness > pop.GlobalBest.Fitness {
				pop.GlobalBest = ind.DeepCopy()
			}
		}

		if fitnessHistory != nil {
			*fitnessHistory = append(*fitnessHistory, pop.GlobalBest.Fitness)
		}

		for _, ind := range pop.Individuals {
			for i := 0; i < pop.M; i++ {
				for j := 0; j < pop.N; j++ {
					idx := i*pop.N + j

					bestBit := pop.GlobalBest.Assignments[i][j]
					currentBit := ind.Assignments[i][j]

					if bestBit != currentBit {
						if bestBit == 1 {
							ind.Theta[idx] += deltaTheta
						} else {
							ind.Theta[idx] -= deltaTheta
						}
						// theta within bounds
						ind.Theta[idx] = math.Max(0, math.Min(math.Pi/2, ind.Theta[idx]))
					}
				}
			}
		}
	}
}
