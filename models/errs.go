package models

import "errors"

var (
	ErrAlreadyExists = errors.New("username already exists.")
	ErrParamIsEmpty = errors.New("input param is empty.")
)
