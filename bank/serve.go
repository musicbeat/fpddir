package bank
import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// Serve implements the Server interface.
func (b Bank) Serve(port string) (searchurl string, err error) {
}
