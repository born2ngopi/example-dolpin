package check

import (
	"errors"
	"testing"

	"github.com/born2ngopi/example-dolpin/mocks"
	"github.com/born2ngopi/example-dolpin/types"
	"go.uber.org/mock/gomock"
)

func TestCheckFunction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := mocks.NewMockService(ctrl)

	tests := []struct {
		name       string
		a, b       int
		aStr, bStr string
		wantError  error
		wantRes1   int
		wantRes2   int
		prepare    func()
	}{
		{
			name:      "Success",
			a:         1,
			b:         2,
			aStr:      "3",
			bStr:      "4",
			wantError: nil,
			wantRes1:  3,
			wantRes2:  7,
			prepare: func() {
				s.EXPECT().Sum(1, 2).Return(3)
				s.EXPECT().SumFromStr(types.SumField{A: "3", B: "4"}).Return(7, nil)
			},
		},
		{
			name:      "Error in SumFromStr",
			a:         1,
			b:         2,
			aStr:      "3",
			bStr:      "4",
			wantError: errors.New("mock error"),
			wantRes1:  0,
			wantRes2:  0,
			prepare: func() {
				s.EXPECT().Sum(1, 2).Return(3)
				s.EXPECT().SumFromStr(types.SumField{A: "3", B: "4"}).Return(0, errors.New("mock error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()

			res1, res2, err := CheckFunction(tt.a, tt.b, tt.aStr, tt.bStr, s)

			if tt.wantError != nil {
				if err == nil {
					t.Errorf("Expected error: %v, got nil", tt.wantError)
				} else if err.Error() != tt.wantError.Error() {
					t.Errorf("Expected error: %v, got %v", tt.wantError, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected nil error, got %v", err)
				}

				if res1 != tt.wantRes1 {
					t.Errorf("Expected result 1: %v, got %v", tt.wantRes1, res1)
				}

				if res2 != tt.wantRes2 {
					t.Errorf("Expected result 2: %v, got %v", tt.wantRes2, res2)
				}
			}
		})
	}
}
