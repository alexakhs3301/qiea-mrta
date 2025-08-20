package common

type Robot struct {
	ID       int `json:"id"`
	Capacity int `json:"capacity"`
}

type Task struct {
	ID            int `json:"id"`
	RequiredUnits int `json:"required_units"`
}

type Input struct {
	Robots  []Robot `json:"robots"`
	Tasks   []Task  `json:"tasks"`
	Utility [][]int `json:"utility"`
	Cost    [][]int `json:"cost"`
}
