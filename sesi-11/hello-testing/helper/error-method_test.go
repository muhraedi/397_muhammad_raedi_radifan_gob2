package helper

import "fmt"
import "testing"

func TestErrorMethodFailSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 40 {
		t.Error("Result has to be 40")
	}

	fmt.Println("TestErrorMethodFailSum Eksekusi Terhenti")
}

func TestErrorMethodSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		t.FailNow()
	}

	fmt.Println("TestErrorMethodSum Eksekusi Terhenti")
}
