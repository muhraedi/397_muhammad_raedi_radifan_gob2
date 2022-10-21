package helper

import "fmt"
import "testing"

func TestFailMethodFailSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 40 {
		t.Fail()
	}

	fmt.Println("TestFailMethodFailSum Eksekusi Terhenti")
}

func TestFailMethodSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		t.FailNow()
	}

	fmt.Println("TestFailMethodSum Eksekusi Terhenti")
}
