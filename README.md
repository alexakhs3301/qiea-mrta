# QIEA‑MRTA

**Quantum-Inspired Evolutionary Algorithm for Multi-Robot Task Allocation**

A Go implementation of a **Quantum-Inspired Evolutionary Algorithm (QIEA)** tailored to solve **Multi-Robot Task Allocation (MRTA)** problems.

---

##  Background

- **QIEA** combines concepts from quantum computing—like qubits, superposition, and quantum gates—with classic evolutionary algorithms to enhance optimization performance. ([capalearning.com](https://capalearning.com/2023/04/02/what-is-quantum-inspired-evolutionary-algorithm/))
- **MRTA** refers to the problem of assigning tasks among multiple robots, aiming for efficient, conflict-free allocation in scenarios like dispatching, inspection, or cooperative missions. ([arxiv.org](https://arxiv.org/abs/2411.02062), [github.com](https://github.com/LT-UK/MRTA))
- This project bridges these areas: it applies quantum-inspired optimization techniques to MRTA challenges using Go.

---

##  Project Structure

```
qiea-mrta/
├── common/     → Shared utilities and data structures across QIEA and MRTA modules
├── ga/         → Genetic algorithm components for comparison or hybrid approaches
├── quantum/    → Core QIEA logic, including qubit encoding and quantum rotation gates
├── main.go     → Entry point: configuration, execution, and result output
├── go.mod      → Go module dependencies
├── LICENSE     → MIT License
└── README.md
```

---

##  Getting Started

### Prerequisites

- Go 1.18 or higher installed on your system.

### Installation

```bash
# Clone the repository
git clone https://github.com/alexakhs3301/qiea-mrta.git
cd qiea-mrta

# Tidy up dependencies
go mod tidy
```

### Running the Project

Edit parameters in `main.go`, such as:
- Problem setup (task and robot definitions)
- Population size and algorithm-specific settings

Then execute:
```bash
go run main.go
```

Expected output includes details on allocation performance, convergence progress, and any results summaries.

---

##  Customization & Configuration

Adjust key parameters in `main.go` to tailor algorithm behavior:
- Population size
- Number of generations
- Task fitness functions
- Quantum-inspired rotation angles or strategies

You might also compare with standard Genetic Algorithm (GA) approaches located in the `ga/` folder.

---

##  Results & Performance

The QIEA approach offers potential benefits in convergence speed and solution quality for MRTA tasks—particularly when compared to traditional evolutionary methods. Consider benchmarking QIEA against GA under different scenarios.

---

##  License

Licensed under the **MIT License** — see the [LICENSE](LICENSE) file for full details.

---

##  References

- What is Quantum-Inspired Evolutionary Algorithm? ([capalearning.com](https://capalearning.com/2023/04/02/what-is-quantum-inspired-evolutionary-algorithm/))  
- MRTA frameworks and methodologies ([arxiv.org](https://arxiv.org/abs/2411.02062?utm_source=chatgpt.com), [github.com](https://github.com/LT-UK/MRTA))  

---

## Disclaimer

This project is developed and maintained primarily for **academic research and educational purposes**.  
It is not intended for direct commercial use or deployment in safety-critical systems.  
Use at your own discretion and responsibility.

---

##  Contact

Maintained by **Alex M** (GitHub: @alexakhs3301). For support or questions, feel free to open an issue or reach out via GitHub.
