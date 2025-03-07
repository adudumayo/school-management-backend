package view

// Quiz model
type Quiz struct {
	ID       int    `json:"id"`
	Subject  string `json:"subject"`
	Topic    string `json:"topic"`
	Question string `json:"question"`
	Due_date string `json:"due_date"`
}
