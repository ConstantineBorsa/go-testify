package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req, err := http.NewRequest("GET", "/?city=moscow&count="+strconv.Itoa(totalCount+1), nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(MainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	assert.Equal(t, totalCount, len(strings.Split(responseRecorder.Body.String(), ",")))
}

func TestMainHandlerWhenOk(t *testing.T) {
	req, err := http.NewRequest("GET", "/?city=moscow&count=4", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(MainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	assert.NotEmpty(t, responseRecorder.Body.String())

}

func TestMainHandlerWrongCity(t *testing.T) {
	req, err := http.NewRequest("GET", "/?city=wrongcity&count=1", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(MainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	expectedError := "wrong city value"
	assert.Contains(t, responseRecorder.Body.String(), expectedError)

}
