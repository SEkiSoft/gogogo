// Copyright (c) 2016 SEkiSoft
// See License.txt

package model

import "strconv"

type AppError struct {
	Where      string
	Message    string
	StatusCode int
}

func (er *AppError) Error() string {
	return er.Where + ": " + er.Message + ", " + strconv.Itoa(er.StatusCode)
}

func NewAppError(where, message string, statusCode int) *AppError {
	return &AppError{
		Where:      where,
		Message:    message,
		StatusCode: statusCode,
	}
}
