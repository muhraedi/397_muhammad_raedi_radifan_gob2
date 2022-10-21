package helper

import "fmt"
import "testing"

func TestFatalMethodFailSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 40 {
		t.Error("Result has to be 40")
	}

	fmt.Println("TestFatalMethodFailSum Eksekusi Terhenti")
}

func TestFatalMethodSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 30 {
		t.Fatal("Result has to be 20")
	}

	fmt.Println("TestFatalMethodSum Eksekusi Terhenti")
}
