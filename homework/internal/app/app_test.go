package app

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"homework/internal/devices"
	deviceMock "homework/internal/mocks"
)

func TestCreateDevice(t *testing.T) {
	const (
		testSeqNum1 = "test 1"
		testIP1     = "test ip 1"
		testModel1  = "test model 1"
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := deviceMock.NewMockRepository(ctrl)

	device := &devices.Device{
		SerialNum: testSeqNum1,
		IP:        testIP1,
		Model:     testModel1,
	}
	repo.EXPECT().Create(device).Return(nil).Times(1)

	app := NewService(repo)
	err := app.CreateDevice(device)

	require.NoError(t, err)
}

func TestGetDevice(t *testing.T) {
	const (
		testSeqNum1 = "test 1"
		testIP1     = "test ip 1"
		testModel1  = "test model 1"
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := deviceMock.NewMockRepository(ctrl)

	expect := &devices.Device{
		SerialNum: testSeqNum1,
		IP:        testIP1,
		Model:     testModel1,
	}
	repo.EXPECT().Get(testSeqNum1).Return(expect, nil).Times(1)

	app := NewService(repo)
	actual, err := app.GetDevice(testSeqNum1)

	require.NoError(t, err)
	require.Equal(t, actual, expect)
}

func TestUpdateDevice(t *testing.T) {
	const (
		testSeqNum1 = "test 1"
		testIP1     = "test ip 1"
		testModel1  = "test model 1"
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := deviceMock.NewMockRepository(ctrl)

	expect := &devices.Device{
		SerialNum: testSeqNum1,
		IP:        testIP1,
		Model:     testModel1,
	}
	repo.EXPECT().Update(expect).Return(nil).Times(1)

	app := NewService(repo)
	err := app.UpdateDevice(expect)

	require.NoError(t, err)
}

func TestDeleteDevice(t *testing.T) {
	const (
		testSeqNum1 = "test 1"
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := deviceMock.NewMockRepository(ctrl)

	repo.EXPECT().Delete(testSeqNum1).Return(nil).Times(1)

	app := NewService(repo)
	err := app.DeleteDevice(testSeqNum1)

	require.NoError(t, err)
}
