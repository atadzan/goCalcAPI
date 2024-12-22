package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/atadzan/goCalcAPI/models"
	"github.com/atadzan/goCalcAPI/pkg/service"
)

func TestCalculate(t *testing.T) {
	testCases := []struct {
		Name             string
		InputBody        models.InputParams
		ExpectedResponse models.Response
		IsErrExpected    bool
		ExpectedErr      errorStruct
	}{
		{
			Name:             "Success",
			InputBody:        models.InputParams{Expression: "2+2*2"},
			ExpectedResponse: models.Response{Result: 6},
			IsErrExpected:    false,
			ExpectedErr:      errorStruct{},
		},
		{
			Name:             "Invalid expression",
			InputBody:        models.InputParams{Expression: "a+2*2"},
			ExpectedResponse: models.Response{},
			IsErrExpected:    true,
			ExpectedErr: errorStruct{
				ErrorMsg: "Expression is not valid",
			},
		},
		{
			Name:             "Internal server error",
			InputBody:        models.InputParams{Expression: "10/0"},
			ExpectedResponse: models.Response{},
			IsErrExpected:    true,
			ExpectedErr: errorStruct{
				ErrorMsg: "Internal server error",
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			rawBody, err := json.Marshal(testCase.InputBody)
			if err != nil {
				t.Errorf("can't marshal input body. Err: %v", err)
				return
			}
			req := httptest.NewRequest("POST", "localhost:8080/api/v1/calculate", bytes.NewReader(rawBody))
			w := httptest.NewRecorder()
			serviceInstance := service.New()
			handler := New(serviceInstance)
			handler.Calculate(w, req)
			resp := w.Result()
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			if testCase.IsErrExpected {
				var response errorStruct
				if err = json.Unmarshal(body, &response); err != nil {
					return
				}
				if ok := reflect.DeepEqual(testCase.ExpectedErr, response); !ok {
					t.Errorf("expected %v but got %v", testCase.ExpectedErr, response)
				}
			} else {
				var response models.Response
				if err = json.Unmarshal(body, &response); err != nil {
					return
				}
				if ok := reflect.DeepEqual(testCase.ExpectedResponse, response); !ok {
					t.Errorf("expected %v but got %v", testCase.ExpectedResponse, response)
				}
			}
		})
	}

}
