package models

type User struct {
	Id       int64  `json:"id" bun:",pk,autoincrement"`
	Name     string `json:"name" bun:"unique"`
	Password string `json:"password"`
}
