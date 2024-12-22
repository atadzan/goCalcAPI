package service

import (
	"errors"
	"strconv"
	"unicode"

	"github.com/atadzan/goCalcAPI/models"
)

type Service interface {
	Calculate(expression string) (resp models.Response, err error)
}

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) Calculate(expression string) (resp models.Response, err error) {
	var values []float64
	var ops []byte

	applyTopOperation := func() error {
		if len(ops) == 0 || len(values) < 2 {
			return ErrExpressionIsNotValid
		}

		b := values[len(values)-1]
		a := values[len(values)-2]
		values = values[:len(values)-2]

		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]

		result, err := applyOperation(op, a, b)
		if err != nil {
			return err
		}
		values = append(values, result)
		return nil
	}

	for i := 0; i < len(expression); i++ {
		ch := expression[i]

		if ch == ' ' {
			continue
		}
		if unicode.IsDigit(rune(ch)) {
			start := i
			for i < len(expression) && (unicode.IsDigit(rune(expression[i])) || expression[i] == '.') {
				i++
			}
			value, err := strconv.ParseFloat(expression[start:i], 64)
			if err != nil {
				return resp, ErrExpressionIsNotValid
			}
			values = append(values, value)
			i--
		} else if ch == '(' {
			ops = append(ops, ch)
		} else if ch == ')' {
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				if err = applyTopOperation(); err != nil {
					return resp, err
				}
			}
			if len(ops) == 0 {
				return resp, errors.New("mismatched parentheses")
			}
			ops = ops[:len(ops)-1]
		} else if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			for len(ops) > 0 && precedence(ops[len(ops)-1]) >= precedence(ch) {
				if err = applyTopOperation(); err != nil {
					return resp, err
				}
			}
			ops = append(ops, ch)
		} else {
			return resp, ErrExpressionIsNotValid
		}
	}
	for len(ops) > 0 {
		if err = applyTopOperation(); err != nil {
			return resp, err
		}
	}
	if len(values) != 1 {
		return resp, ErrExpressionIsNotValid
	}

	return models.Response{
		Result: values[0],
	}, nil
}
