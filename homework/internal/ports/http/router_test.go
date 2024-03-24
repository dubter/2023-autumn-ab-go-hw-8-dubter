package http

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	deviceMock "homework/internal/mocks"
)

func TestNewHandler(t *testing.T) {
	const (
		testPort         = "test port"
		testHost         = "test host"
		testReadTimeout  = time.Duration(10)
		testWriteTimeout = time.Duration(10)
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	app := deviceMock.NewMockService(ctrl)

	cases := []struct {
		name   string
		config *Config
	}{
		{
			name: "test1: all fields are correct",
			config: &Config{
				Service:      app,
				Port:         testPort,
				Host:         testHost,
				ReadTimeout:  testReadTimeout,
				WriteTimeout: testWriteTimeout,
			},
		},
		{
			name: "test2: empty full address",
			config: &Config{
				Service:      app,
				ReadTimeout:  testReadTimeout,
				WriteTimeout: testWriteTimeout,
			},
		},
		{
			name: "test2: empty ReadTimeout",
			config: &Config{
				Service:      app,
				Port:         testPort,
				Host:         testHost,
				WriteTimeout: testWriteTimeout,
			},
		},
		{
			name: "test3: empty WriteTimeout",
			config: &Config{
				Service:     app,
				Port:        testPort,
				Host:        testHost,
				ReadTimeout: testReadTimeout,
			},
		},
		{
			name: "test4: nil Service",
			config: &Config{
				Service:      nil,
				Port:         testPort,
				Host:         testHost,
				ReadTimeout:  testReadTimeout,
				WriteTimeout: testWriteTimeout,
			},
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			handler := NewHandler(tCase.config)

			server := handler.NewServer()

			require.NotNil(t, server)
		})
	}
}
