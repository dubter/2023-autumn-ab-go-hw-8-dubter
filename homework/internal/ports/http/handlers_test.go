package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"homework/internal/devices"
	"homework/internal/errors"
	deviceMock "homework/internal/mocks"
)

const (
	testSeqNum1     = "1"
	testIP1         = "test ip 1"
	testModel1      = "test model 1"
	testInvalidBody = "test body"
)

func TestHandlerCreateDeviceSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	deviceService := deviceMock.NewMockService(ctrl)

	handler := &Handler{
		service: deviceService,
	}

	device := &devices.Device{
		SerialNum: testSeqNum1,
		IP:        testIP1,
		Model:     testModel1,
	}
	deviceService.EXPECT().CreateDevice(device).Return(nil).Times(1)

	deviceBytes, _ := json.Marshal(device)
	reqBody := bytes.NewReader(deviceBytes)

	router := chi.NewRouter()
	router.Post("/devices", handler.createDevice)

	r := httptest.NewRequest(http.MethodPost, "/devices", reqBody)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)
}

func TestHandlerCreateDeviceNilBody(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	deviceService := deviceMock.NewMockService(ctrl)

	handler := &Handler{
		service: deviceService,
	}

	router := chi.NewRouter()
	router.Post("/devices", handler.createDevice)

	r := httptest.NewRequest(http.MethodPost, "/devices", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestHandlerCreateDeviceInvalidRequestBody(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	deviceService := deviceMock.NewMockService(ctrl)

	handler := &Handler{
		service: deviceService,
	}

	invalidDeviceBytes := []byte(testInvalidBody)
	reqBody := bytes.NewReader(invalidDeviceBytes)

	router := chi.NewRouter()
	router.Post("/devices", handler.createDevice)

	r := httptest.NewRequest(http.MethodPost, "/devices", reqBody)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestHandlerCreateDeviceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	deviceService := deviceMock.NewMockService(ctrl)

	device := &devices.Device{
		SerialNum: testSeqNum1,
		IP:        testIP1,
		Model:     testModel1,
	}
	deviceService.EXPECT().CreateDevice(device).Return(errors.NewAlreadyExistDeviceError("")).Times(1)

	handler := &Handler{
		service: deviceService,
	}

	deviceBytes, _ := json.Marshal(device)
	reqBody := bytes.NewReader(deviceBytes)

	router := chi.NewRouter()
	router.Post("/devices", handler.createDevice)

	r := httptest.NewRequest(http.MethodPost, "/devices", reqBody)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestHandlerGetDeviceSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	deviceService := deviceMock.NewMockService(ctrl)

	device := &devices.Device{
		SerialNum: testSeqNum1,
		IP:        testIP1,
		Model:     testModel1,
	}
	deviceService.EXPECT().GetDevice(testSeqNum1).Return(device, nil).Times(1)

	handler := &Handler{
		service: deviceService,
	}
	router := chi.NewRouter()
	router.Get("/devices/{id}", handler.getDevice)

	r := httptest.NewRequest(http.MethodGet, "/devices/"+testSeqNum1, nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)

	var resBody []byte
	resBody, err := io.ReadAll(res.Body)
	require.NoError(t, err)

	var actual devices.Device
	err = json.Unmarshal(resBody, &actual)
	require.NoError(t, err)
	require.Equal(t, actual, *device)
}

func TestHandlerGetDeviceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	deviceService := deviceMock.NewMockService(ctrl)

	deviceService.EXPECT().GetDevice(testSeqNum1).Return(nil, errors.NewNotFoundError("")).Times(1)

	handler := &Handler{
		service: deviceService,
	}
	router := chi.NewRouter()
	router.Get("/devices/{id}", handler.getDevice)

	r := httptest.NewRequest(http.MethodGet, "/devices/"+testSeqNum1, nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestHandlerDeleteDeviceSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	deviceService := deviceMock.NewMockService(ctrl)
	deviceService.EXPECT().DeleteDevice(testSeqNum1).Return(nil).Times(1)

	handler := &Handler{
		service: deviceService,
	}
	router := chi.NewRouter()
	router.Delete("/devices/{id}", handler.deleteDevice)

	r := httptest.NewRequest(http.MethodDelete, "/devices/"+testSeqNum1, nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)
}

func TestHandlerDeleteDeviceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	deviceService := deviceMock.NewMockService(ctrl)
	deviceService.EXPECT().DeleteDevice(testSeqNum1).Return(errors.NewNotFoundError("")).Times(1)

	handler := &Handler{
		service: deviceService,
	}
	router := chi.NewRouter()
	router.Delete("/devices/{id}", handler.deleteDevice)

	r := httptest.NewRequest(http.MethodDelete, "/devices/"+testSeqNum1, nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestHandlerUpdateDeviceSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	deviceService := deviceMock.NewMockService(ctrl)

	device := &devices.Device{
		SerialNum: testSeqNum1,
		IP:        testIP1,
		Model:     testModel1,
	}
	deviceService.EXPECT().UpdateDevice(device).Return(nil).Times(1)

	handler := &Handler{
		service: deviceService,
	}
	router := chi.NewRouter()
	router.Put("/devices/{id}", handler.updateDevice)

	deviceBytes, _ := json.Marshal(device)
	reqBody := bytes.NewReader(deviceBytes)

	r := httptest.NewRequest(http.MethodPut, "/devices/"+testSeqNum1, reqBody)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)
}

func TestHandlerUpdateDeviceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	deviceService := deviceMock.NewMockService(ctrl)

	device := &devices.Device{
		SerialNum: testSeqNum1,
		IP:        testIP1,
		Model:     testModel1,
	}
	deviceService.EXPECT().UpdateDevice(device).Return(errors.NewNotFoundError("")).Times(1)

	handler := &Handler{
		service: deviceService,
	}
	router := chi.NewRouter()
	router.Put("/devices/{id}", handler.updateDevice)

	deviceBytes, _ := json.Marshal(device)
	reqBody := bytes.NewReader(deviceBytes)

	r := httptest.NewRequest(http.MethodPut, "/devices/"+testSeqNum1, reqBody)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestHandlerUpdateDeviceNilBody(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	deviceService := deviceMock.NewMockService(ctrl)

	handler := &Handler{
		service: deviceService,
	}
	router := chi.NewRouter()
	router.Put("/devices/{id}", handler.updateDevice)

	r := httptest.NewRequest(http.MethodPut, "/devices/"+testSeqNum1, nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusBadRequest, res.StatusCode)
}
