package helpers

type Entry struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Website  string `json:"website"`
	Category string `json:"category"`
	Note     string `json:"note"`
}

type Entries struct {
	Entries []Entry `json:"entries"`
}
