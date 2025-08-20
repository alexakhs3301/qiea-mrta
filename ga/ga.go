package ga

import (
	"Code/common"
	"sort"
)

func RunGA(
	populationSize int,
	generations int,
	robots []common.Robot,
	tasks []common.Task,
	utility [][]int,
	costs [][]int,
	lambdaCap, lambdaCoop, lambdaCapCost, lambdaCoopCost float64,
	w1, w2 float64,
	mutationRate float64,
	fitnessHistory *[]float64,
) *Individual {
	m := len(robots)
	n := len(tasks)

	// --- Initialize population ---
	population := make([]*Individual, populationSize)
	for i := 0; i < populationSize; i++ {
		ind := NewIndividual(m, n)
		ind.EvaluateFitness(robots, tasks, utility, costs,
			lambdaCap, lambdaCoop, lambdaCapCost, lambdaCoopCost,
			w1, w2)
		population[i] = ind
	}

	// --- Evolution loop ---
	for gen := 0; gen < generations; gen++ {
		// Sort population by fitness (descending)
		sort.Slice(population, func(i, j int) bool {
			return population[i].Fitness > population[j].Fitness
		})

		if fitnessHistory != nil {
			*fitnessHistory = append(*fitnessHistory, population[0].Fitness)
		}

		// Elitism: keep top 1
		nextGen := []*Individual{population[0].DeepCopy()}

		// Generate offspring
		for len(nextGen) < populationSize {
			parent1 := TournamentSelection(population)
			parent2 := TournamentSelection(population)

			child := UniformCrossover(parent1, parent2)
			child.Mutate(mutationRate)
			child.EvaluateFitness(robots, tasks, utility, costs,
				lambdaCap, lambdaCoop, lambdaCapCost, lambdaCoopCost,
				w1, w2)
			nextGen = append(nextGen, child)
		}

		population = nextGen
	}

	// Return best individual after last generation
	sort.Slice(population, func(i, j int) bool {
		return population[i].Fitness > population[j].Fitness
	})

	return population[0]
}
