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

type ValidUserJwt struct {
	Valid       bool   `json:"valid"`
	Description string `json:"description"`
	JwtToken    string `json:"jwttoken"`
}

type UserV struct {
	Name         string `json:"name"`
	Password     string `json:"password"`
	Valid        bool   `json:"valid"`
	Description  string `json:"description"`
	SessionToken string `json:"sessiontoken"`
	CSRFToken    string `json:"csrftoken"`
}

func (u *UserV) UpdateLoggedUserValues(User string, sessionToken string, csrfToken string) {
	u.Name = User
	u.SessionToken = sessionToken
	u.CSRFToken = csrfToken
}
