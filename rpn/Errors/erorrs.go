package Errors

import (
	"errors"
)

var (
	//FOR CALC
	ErrInvalidExpression  = errors.New("invalid expression")
	ErrDivisionByZero     = errors.New("division by zero")
	ErrMismatchedBrackets = errors.New("mismatched brackets in expression")
	ErrEmptyExpression    = errors.New("expression is too short")
	ErrInvalidInput       = errors.New("invalid input")
	VeryImportantERRO     = errors.New("Фатальная ошибка лицеиста")

	//FOR HTTP
	ErrStartServer = errors.New("Failed to start server")

	//FOR LOAD JSON
	ErrLoafJson = errors.New("Failed to parse JSON")

	//FOR LOGGER
	ErrInvalidPrefix = errors.New("invalid prefixes configuration")
)
