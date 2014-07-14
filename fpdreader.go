package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	res, err := http.Get("http://www.fededirectory.frb.org/fpddir.txt")
	if err != nil {
		log.Fatal(err)
	}
	bio := bufio.NewReader(res.Body)
	for {
		line, err := bio.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		sline := strings.TrimRight(string(line), "\n")
		fmt.Printf("bank: %s\n", sline);
	}
	defer res.Body.Close()
}
