package errors

import (
	"fmt"
)

type AlreadyExistDeviceError struct {
	err error
}

func (e *AlreadyExistDeviceError) Error() string {
	return e.err.Error()
}

func NewAlreadyExistDeviceError(serialNum string) *AlreadyExistDeviceError {
	return &AlreadyExistDeviceError{
		err: fmt.Errorf("device with 'SerialNum' = %s already exist", serialNum),
	}
}

type NotFoundError struct {
	err error
}

func (e *NotFoundError) Error() string {
	return e.err.Error()
}

func NewNotFoundError(serialNum string) *NotFoundError {
	return &NotFoundError{
		err: fmt.Errorf("device with 'SerialNum' = %s not found", serialNum),
	}
}
