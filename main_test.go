package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenRequestIsCorrect(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "expected status code: 200")
	assert.NotEmpty(t, resp.Body.String(), "response body should not be empty")
}

func TestMainHandlerWhenCityIsWrong(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=piter", nil)
	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusBadRequest, resp.Code, "expected status code: 400")
	assert.Equal(t, "wrong city value", resp.Body.String(), "expected error message: wrong city value")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	count := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(resp, req)
	body := resp.Body.String()
	list := strings.Split(body, ",")
	assert.Equal(t, http.StatusOK, resp.Code, "expected status code: 200")
	assert.Equal(t, count, len(list), "expected cafe count to be all available cafes")
}
