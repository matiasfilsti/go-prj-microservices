package models

type User struct {
	Id       int64  `json:"id" bun:",pk,unique"`
	Name     string `json:"name" bun:"unique"`
	Password string `json:"-"`
}
