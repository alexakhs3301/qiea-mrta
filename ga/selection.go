package ga

import "math/rand"

func TournamentSelection(population []*Individual) *Individual {
	tournamentSize := 3
	best := population[rand.Intn(len(population))]

	for i := 1; i < tournamentSize; i++ {
		challenger := population[rand.Intn(len(population))]
		if challenger.Fitness > best.Fitness {
			best = challenger
		}
	}

	return best.DeepCopy()
}
