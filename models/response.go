package models

type Response struct {
	Ok          bool
	Result      interface{}
	Description string
}
