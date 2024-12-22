package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/atadzan/goCalcAPI/models"
	"github.com/atadzan/goCalcAPI/pkg/service"
)

type Handler struct {
	service service.Service
}

func New(service service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Calculate(w http.ResponseWriter, r *http.Request) {
	var input models.InputParams

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		newErrorResp(w, http.StatusInternalServerError, internalServerErrorMsg)
		return
	}
	defer r.Body.Close()
	if err = json.Unmarshal(reqBody, &input); err != nil {
		newErrorResp(w, http.StatusInternalServerError, internalServerErrorMsg)
		return
	}
	resp, err := h.service.Calculate(input.Expression)
	if err != nil {
		if errors.Is(err, service.ErrExpressionIsNotValid) {
			newErrorResp(w, http.StatusUnprocessableEntity, expressionIsNotValidMsg)
		} else {
			newErrorResp(w, http.StatusInternalServerError, internalServerErrorMsg)
		}
		return
	}
	rawResp, err := json.Marshal(resp)
	if err != nil {
		newErrorResp(w, http.StatusInternalServerError, internalServerErrorMsg)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(rawResp)
}
