package errors

import (
	"fmt"
)

type ApiErr struct {
	Message error  `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func BadRequestHandler(message error, status int) *ApiErr {
	fmt.Println(message)
	return &ApiErr{
		Message: message,
		Status:  status,
		Error:   "Bad Request",
	}
}
