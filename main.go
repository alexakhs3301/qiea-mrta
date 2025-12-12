package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"Code/common"
	"Code/ga"
	"Code/quantum"
)

func writeFitnessCSV(path string, fitness []float64, seed int64) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	// header with seed
	if err := w.Write([]string{"seed", strconv.FormatInt(seed, 10)}); err != nil {
		return err
	}
	// column header
	if err := w.Write([]string{"iteration", "fitness"}); err != nil {
		return err
	}
	for i, v := range fitness {
		if err := w.Write([]string{strconv.Itoa(i), fmt.Sprintf("%g", v)}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	mode := "ga" // change this according to the algorithm you want to run ("ga" or "qiea")
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

	input, err := common.LoadInputFromJSON("common/input_hard.json")
	if err != nil {
		log.Fatal("Failed to load input:", err)
	}
	robots := input.Robots
	tasks := input.Tasks
	utility := input.Utility
	costs := input.Cost
	m := len(robots)
	n := len(tasks)

	// Ensure results directory exists
	if err := os.MkdirAll("results", 0755); err != nil {
		log.Fatal("Failed to create results directory:", err)
	}

	// Base seed: use current time by default. Change to a fixed integer for reproducibility.
	baseSeed := time.Now().UnixNano()
	fmt.Println("Base seed:", baseSeed)

	runs := 30
	for run := 1; run <= runs; run++ {
		// use deterministic per-run seed derived from baseSeed
		runSeed := baseSeed + int64(run)
		rand.Seed(runSeed)
		var fitnessHistory []float64

		switch mode {
		case "ga":
			best := ga.RunGA(
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
			fmt.Printf("Run %02d GA best fitness: %g (seed %d)\n", run, best.Fitness, runSeed)
			fn := filepath.Join("results", fmt.Sprintf("ga_seed_%d_run_%02d.csv", runSeed, run))
			if err := writeFitnessCSV(fn, fitnessHistory, runSeed); err != nil {
				log.Fatalf("Failed to write fitness CSV for run %d: %v", run, err)
			}
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
			fmt.Printf("Run %02d QIEA best fitness: %g (seed %d)\n", run, bestqiea.Fitness, runSeed)
			fn := filepath.Join("results", fmt.Sprintf("qiea_seed_%d_run_%02d.csv", runSeed, run))
			if err := writeFitnessCSV(fn, fitnessHistory, runSeed); err != nil {
				log.Fatalf("Failed to write fitness CSV for run %d: %v", run, err)
			}
		default:
			log.Fatalf("Unknown mode: %s", mode)
		}
	}
}
