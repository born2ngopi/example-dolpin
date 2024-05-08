package service

import (
	"strconv"

	"github.com/born2ngopi/example-dolpin/types"
)

type service struct {
}

type Service interface {
	Sum(a, b int) int
	SumFromStr(data types.SumField) (int, error)
}

func New() Service {
	return &service{}
}

func (s *service) Sum(a, b int) int {
	return a + b
}

func (s *service) SumFromStr(data types.SumField) (int, error) {
	aInt, err := strconv.Atoi(data.A)
	if err != nil {
		return 0, err
	}

	bInt, err := strconv.Atoi(data.B)
	if err != nil {
		return 0, err
	}

	return aInt + bInt, nil
}
