package check

import (
	"github.com/born2ngopi/example-dolpin/service"
	"github.com/born2ngopi/example-dolpin/types"
)

func CheckFunction(a, b int, aStr, bStr string) (int, int, error) {

	s := service.New()

	res1 := s.Sum(a, b)

	res2, err := s.SumFromStr(types.SumField{
		A: aStr,
		B: bStr,
	})
	if err != nil {
		return 0, 0, err
	}

	return res1, res2, nil
}
