package _struct

type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Error error `json:"error"`
}