package products

import "errors"

var (
	ErrConnectDatabase  = errors.New("error to connect database")
	ErrResourceNotFound = errors.New("resource not found")
)
