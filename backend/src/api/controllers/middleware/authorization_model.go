package middleware

type ValidUser struct {
	Valid       bool   `json:"valid"`
	Description string `json:"description"`
}
