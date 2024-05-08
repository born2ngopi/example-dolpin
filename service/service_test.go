package service

import (
	"errors"
	"testing"

	"github.com/born2ngopi/example-dolpin/types"
	"go.uber.org/mock/gomock"
)

func TestSumFromStr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name           string
		sumField       types.SumField
		expectedError  error
		expectedResult int
	}{
		{
			name:           "Valid input",
			sumField:       types.SumField{A: "1", B: "2"},
			expectedError:  nil,
			expectedResult: 3,
		},
		{
			name:           "Invalid input (A)",
			sumField:       types.SumField{A: "invalid", B: "2"},
			expectedError:  errors.New("strconv.Atoi: parsing \"invalid\": invalid syntax"),
			expectedResult: 0,
		},
		{
			name:           "Invalid input (B)",
			sumField:       types.SumField{A: "1", B: "invalid"},
			expectedError:  errors.New("strconv.Atoi: parsing \"invalid\": invalid syntax"),
			expectedResult: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Call function
			service := New()
			result, err := service.SumFromStr(tt.sumField)

			// Assert
			if tt.expectedError != nil {
				if err == nil {
					t.Errorf("Expected error: %v, got nil", tt.expectedError)
				} else if err.Error() != tt.expectedError.Error() {
					t.Errorf("Expected error: %v, got %v", tt.expectedError, err)
				}
			}
			if result != tt.expectedResult {
				t.Errorf("Unexpected result: %d", result)
			}
		})
	}
}
