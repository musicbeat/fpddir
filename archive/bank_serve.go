package bank
import (
	// "bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	// "strings"
)

func SearchServer(w http.ResponseWriter, req *http.Request) {
	if len(routingNumberMap) == 0 {
		io.WriteString(w, "searches will be fast, but empty\n")
	}
	io.WriteString(w, "search!\n")
}

func DumpServer(w http.ResponseWriter, req *http.Request) {
	j, err := json.MarshalIndent(routingNumberMap, "", "  ")
	if err == nil {
		io.WriteString(w, fmt.Sprintf("%s\n",j))
	} else {
		log.Fatal("bank.DumpServer: ", err)
	}
}

func KillServer(w http.ResponseWriter, req *http.Request) {
	os.Exit(0)
}

// Serve implements the Server interface.
func (b Bank) Serve(port string) (err error) {
	http.HandleFunc("/bank/search", SearchServer)
	http.HandleFunc("/bank/dump", DumpServer)
	http.HandleFunc("/bank/kill", KillServer)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	return err
}
