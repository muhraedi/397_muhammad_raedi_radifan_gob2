package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestifyFailSum(t *testing.T) {
	result := Sum(10, 10)

	require.Equal(t, 40, result, "Result has to be 40")

	fmt.Println("TestifyFailSum Eksekusi Terhenti")
}

func TestifySum(t *testing.T) {
	result := Sum(10, 10)

	assert.Equal(t, 20, result, "Result has to be 20")

	fmt.Println("Testify Eksekusi Terhenti")
}
