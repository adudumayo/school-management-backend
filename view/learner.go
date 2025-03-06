package view

// Learner model
type Learner struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Class   int     `json:"class"`
	Average float64 `json:"average"`
}
