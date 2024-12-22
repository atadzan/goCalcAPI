package models

type InputParams struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result float64 `json:"result"`
}
