package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	expressionIsNotValidMsg = "Expression is not valid"
	internalServerErrorMsg  = "Internal server error"
)

type errorStruct struct {
	ErrorMsg string `json:"error"`
}

func newErrorResp(w http.ResponseWriter, httpStatus int, respMsg string) {
	rawBody, err := json.Marshal(errorStruct{ErrorMsg: respMsg})
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(httpStatus)
	_, err = w.Write(rawBody)
	if err != nil {
		log.Printf("can't write response body. Err: %v. HttpStatus: %d. Response msg: %s", err, httpStatus, respMsg)
	}
}
