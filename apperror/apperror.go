package apperror

import "fmt"

type AppError struct {
	ErrorCode    int
	ErrorMessage string
}

func (apperror AppError) Error() string {
	return fmt.Sprintf("%v - %v", apperror.ErrorCode, apperror.ErrorMessage)
}