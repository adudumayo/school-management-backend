package view

// Teacher model
type Teacher struct {
	ID       int    `json:"id"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Title    string `json:"title"`
	Password string `json:"password"`
}
