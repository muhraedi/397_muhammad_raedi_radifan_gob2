package helper

import (
	"fmt"
	"testing"

	_ "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableSum(t *testing.T) {
	tests := []struct {
		request  int
		expected int
		errMsg   string
	}{
		{
			request:  Sum(10, 10),
			expected: 20,
			errMsg:   "Result has to be 20",
		},
		{
			request:  Sum(20, 20),
			expected: 40,
			errMsg:   "Result has to be 40",
		},
		{
			request:  Sum(25, 25),
			expected: 50,
			errMsg:   "Result has to be 50",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require.Equal(t, test.expected, test.request, test.errMsg)
		})
	}
}
