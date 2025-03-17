package models

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ValidUser struct {
	Valid        bool   `json:"valid"`
	Description  string `json:"description"`
	SessionToken string `json:"sessiontoken"`
	CSRFToken    string `json:"csrftoken"`
}
