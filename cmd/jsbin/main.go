package main

import (
	"fmt"
	"github.com/quii/sn-poker/jsonbin"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	bin, err := jsonbin.CreateNewJSONBin(client)

	if err != nil {
		log.Fatalf("failed to create bin %+v", err)
	}

	fmt.Println(bin)
}
