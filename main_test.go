package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=50", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	cafeSlice := strings.Split(responseRecorder.Body.String(), ",")
	cafeCount := len(cafeSlice)

	errMsg := fmt.Sprintf("Ожидается число : %v\nПолучено число : %v", totalCount, cafeCount)

	assert.Equal(t, totalCount, cafeCount, errMsg)

	// здесь нужно добавить необходимые проверки
}

func TestMainHandlerWhenWrongCityParameter(t *testing.T) {

	waitingMsg := "wrong city value"
	req := httptest.NewRequest("GET", "/cafe?count=3&city=LA", nil)

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(resp, req)

	status := resp.Code

	errMsg := fmt.Sprintf("Не верный статус! \nОжидаемый статус: %v\nРеальный статус: %v", http.StatusBadRequest, status)
	assert.Equal(t, status, http.StatusBadRequest, errMsg)

	ans := resp.Body.String()

	assert.Equal(t, waitingMsg, ans)

}

func TestMainHandlerWithMaxNumReq(t *testing.T) {
	cafe, ok := cafeList["moscow"]
	if !ok {
		t.Errorf("Не создан слайс со всеми кафе")
	}

	needResp := strings.Join(cafe[:], ",")

	req := httptest.NewRequest("GET", "/cafe?count=50&city=moscow", nil)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(recorder, req)

	ans := recorder.Body.String()

	assert.Equal(t, needResp, ans)

}
