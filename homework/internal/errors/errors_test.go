package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	errorMessageTestValue = "test error message"
)

func TestAlreadyExistDeviceError(t *testing.T) {
	err := NewAlreadyExistDeviceError(errorMessageTestValue)
	require.NotNil(t, err)
	require.EqualError(t, fmt.Errorf("device with 'SerialNum' = %s already exist", errorMessageTestValue), err.Error())
}

func TestNotFoundError(t *testing.T) {
	err := NewNotFoundError(errorMessageTestValue)
	require.NotNil(t, err)
	require.EqualError(t, fmt.Errorf("device with 'SerialNum' = %s not found", errorMessageTestValue), err.Error())
}
