package main
import (
	"github.com/musicbeat/stddata/bank"
	"log"
)
func main () {
	bank := new(bank.Bank)
	if n, err := bank.Load(); err != nil {
		log.Fatal("bank.Load: ", err)
		if n < 1 {
			log.Fatal("bank.Load: ", n)
		}
	}

	err := bank.Serve(":8080")
	if err != nil {
		log.Fatal("bank.Serve: ", err)
	}
}
