package main

import (
	"Code/common"
	"Code/ga"
	"Code/quantum"
	"fmt"
	"log"
)

func main() {

	mode := "qiea" // change this according to the algorithm you want to run ("ga" or "qiea")
	popSize := 500
	generations := 7000
	deltaTheta := 0.01 * 3.1415
	mutationRate := 0.01

	// Penalty weights
	lambdaCap := 10.0
	lambdaCoop := 15.0
	lambdaCapCost := 5.0
	lambdaCoopCost := 20.0

	// Fitness weights
	w1 := 1.0
	w2 := 1.0

	input, err := common.LoadInputFromJSON("common/input_hardest.json")
	if err != nil {
		log.Fatal("Failed to load input:", err)
	}
	robots := input.Robots
	tasks := input.Tasks
	utility := input.Utility
	costs := input.Cost
	m := len(robots)
	n := len(tasks)
	var fitnessHistory []float64
	var best *ga.Individual

	switch mode {
	case "ga":
		best = ga.RunGA(
			popSize,
			generations,
			robots,
			tasks,
			utility,
			costs,
			lambdaCap, lambdaCoop, lambdaCapCost, lambdaCoopCost,
			w1, w2,
			mutationRate,
			&fitnessHistory,
		)
		fmt.Println("Best GA fitness:", best.Fitness)
	case "qiea":
		population := quantum.NewPopulation(popSize, m, n)
		population.Run(
			robots,
			tasks,
			utility,
			costs,
			generations,
			lambdaCap, lambdaCoop, lambdaCapCost, lambdaCoopCost,
			w1, w2,
			deltaTheta,
			&fitnessHistory,
		)
		bestqiea := population.GlobalBest
		fmt.Println("Best QIEA fitness:", bestqiea.Fitness)
	}
	err = common.PlotFitnessHistory(fitnessHistory, "Fitness Over Time", "results/fitness.png")
	if err != nil {
		log.Fatal("Failed to plot fitness history:", err)
	}
}
