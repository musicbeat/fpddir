package bank

import (
	"fmt"
	"testing"
)

func TestServe(t *testing.T) {
	fmt.Println("Test: bank.Serve")
	bank := new(Bank)
	err := bank.Serve(":8080")
	if err != nil {
		t.Errorf("bank.Serve err %s\n", err)
	}
}
